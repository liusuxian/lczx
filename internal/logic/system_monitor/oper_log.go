package system_monitor

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sOperLog struct {
	pool            *grpool.Pool
	clientOptionMap map[string][]*model.ClientOption // 客户端选项
}

func init() {
	service.RegisterOperLog(newOperLog())
}

// 操作日志服务
func newOperLog() *sOperLog {
	insOperLog := &sOperLog{
		pool: grpool.New(100),
	}
	insOperLog.clientOptionMap = map[string][]*model.ClientOption{}
	statusList := []*model.ClientOption{
		{
			Name:  "异常",
			Value: "0",
		},
		{
			Name:  "正常",
			Value: "1",
		},
	}
	operTypeList := []*model.ClientOption{
		{
			Name:  "读取",
			Value: "GET",
		},
		{
			Name:  "新增",
			Value: "POST",
		},
		{
			Name:  "修改",
			Value: "PUT",
		},
		{
			Name:  "删除",
			Value: "DELETE",
		},
	}
	insOperLog.clientOptionMap["statusList"] = statusList
	insOperLog.clientOptionMap["operTypeList"] = operTypeList

	return insOperLog
}

// Invoke 异步保存日志
func (s *sOperLog) Invoke(req *ghttp.Request) {
	ctx := req.GetCtx()
	curUser := service.Context().Get(ctx).User
	if curUser == nil {
		return
	}
	// 请求地址
	url := req.Request.URL
	// 获取所有菜单
	var err error
	var allMenus []*entity.Menu
	allMenus, err = service.Menu().GetAllMenus(ctx)
	if err != nil {
		logger.Error(ctx, "Invoke GetAllMenus Error: ", err.Error())
		return
	}
	// 获取地址对应的菜单ID
	var menu *entity.Menu
	path := gstr.TrimLeft(url.Path, "/")
	for _, m := range allMenus {
		if gstr.Equal(m.Rule, path) {
			menu = m
			break
		}
	}
	// 组装数据
	data := &entity.OperLog{}
	if menu != nil {
		// 模块标题
		data.Title = menu.Name
	}
	// 方法名称
	data.Method = url.Path
	// 请求方式
	data.ReqMethod = req.Method
	// 操作类别
	data.OperType = 1
	// 操作人员
	data.OperName = curUser.Realname
	// 部门名称
	// 获取部门状态为正常的部门列表
	var depts []*entity.Dept
	depts, err = service.Dept().GetStatusEnableDepts(ctx)
	if err != nil {
		logger.Error(ctx, "Invoke GetStatusEnableDepts Error: ", err.Error())
		return
	}
	data.DeptName = service.Dept().GetDeptAllNameById(depts, curUser.DeptId)
	// 请求URL
	rawQuery := url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	data.ReqUrl = url.Path + rawQuery
	// 操作IP地址
	operIp := utils.GetClientIp(req)
	data.OperIp = operIp
	// 操作地点
	data.OperLocation = utils.GetCityByIp(ctx, operIp)
	// 请求参数
	params := req.GetMap()
	if params != nil {
		if v, ok := params["apiReturnRes"]; ok {
			res := gconv.Map(v)
			// 操作状态
			if gconv.Int(res["code"]) == 0 {
				data.Status = 1
			} else {
				data.Status = 0
			}
			if _, ok = res["data"]; ok {
				delete(res, "data")
			}
			// 返回参数
			jsonRes, _ := gjson.Encode(res)
			if len(jsonRes) > 0 {
				data.JsonResult = string(jsonRes)
			}
			delete(params, "apiReturnRes")
		}
		// 请求参数
		args, _ := gjson.Encode(params)
		if len(args) > 0 {
			data.ReqParam = string(args)
		}
	}
	// 操作时间
	data.Time = gtime.Now()
	err = s.pool.Add(ctx, func(ctx context.Context) {
		// 保存日志
		s.SaveOperLog(ctx, data)
	})
	if err != nil {
		logger.Error(ctx, "OperLog Pool Add Error: ", err.Error())
	}
}

// SaveOperLog 保存日志
func (s *sOperLog) SaveOperLog(ctx context.Context, data *entity.OperLog) {
	_, err := dao.OperLog.Ctx(ctx).FieldsEx(dao.OperLog.Columns().Id).Insert(data)
	if err != nil {
		logger.Error(ctx, "SaveOperLog Error: ", err.Error())
	}
}

// GetOperLogList 获取操作日志列表
func (s *sOperLog) GetOperLogList(ctx context.Context, req *v1.OperLogListReq) (total int, list []*entity.OperLog, err error) {
	gmodel := dao.OperLog.Ctx(ctx)
	columns := dao.OperLog.Columns()
	order := "id DESC"
	if req.Title != "" {
		gmodel = gmodel.WhereLike(columns.Title, "%"+req.Title+"%")
	}
	if req.OperName != "" {
		gmodel = gmodel.WhereLike(columns.OperName, "%"+req.OperName+"%")
	}
	if req.ReqMethod != "" {
		gmodel = gmodel.Where(columns.ReqMethod, req.ReqMethod)
	}
	if req.Status != "" {
		gmodel = gmodel.Where(columns.Status, req.Status)
	}
	if req.StartTime.String() != "" {
		gmodel = gmodel.WhereGTE(columns.Time, req.StartTime)
	}
	if req.EndTime.String() != "" {
		gmodel = gmodel.WhereLTE(columns.Time, req.EndTime)
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	if total, err = gmodel.Count(); err != nil {
		return
	}
	err = gmodel.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	return
}

// DeleteOperLogByIds 通过ID列表删除操作日志
func (s *sOperLog) DeleteOperLogByIds(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.OperLog.Ctx(ctx).WhereIn(dao.OperLog.Columns().Id, ids).Delete()
	return
}

// ClearOperLog 清除操作日志
func (s *sOperLog) ClearOperLog(ctx context.Context) (err error) {
	_, err = dao.OperLog.DB().Exec(ctx, "truncate "+dao.OperLog.Table())
	return
}

// GetClientOptionMap 获取客户端选项Map
func (s *sOperLog) GetClientOptionMap() map[string][]*model.ClientOption {
	return s.clientOptionMap
}
