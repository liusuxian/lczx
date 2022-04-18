package service

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sOperLog struct {
	pool *grpool.Pool
}

var (
	insOperLog = sOperLog{
		pool: grpool.New(100),
	}
)

// OperLog 操作日志服务
func OperLog() *sOperLog {
	return &insOperLog
}

// Invoke 异步保存日志
func (s *sOperLog) Invoke(req *ghttp.Request) {
	ctx := req.GetCtx()
	user := Context().Get(ctx).User
	if user == nil {
		return
	}
	// 请求地址
	url := req.Request.URL
	// 获取所有菜单
	var menuList []*entity.Menu
	var err error
	menuList, err = Menu().GetAllMenus(ctx)
	if err != nil {
		logger.Error(ctx, "Invoke GetAllMenus Error: ", err.Error())
		return
	}
	// 获取地址对应的菜单ID
	var menu *entity.Menu
	path := gstr.TrimLeft(url.Path, "/")
	for _, m := range menuList {
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
	data.OperName = user.Realname
	// 部门名称
	var deptName string
	deptName, err = Dept().GetDeptNameById(ctx, user.DeptId)
	if err != nil {
		logger.Error(ctx, "Invoke GetDeptNameById Error: ", err.Error())
	}
	data.DeptName = deptName
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
		res := make(map[string]interface{})
		if v, ok := params["apiReturnResp"]; ok {
			res = gconv.Map(v)
		} else {
			res = gconv.Map(req.GetHandlerResponse())
		}
		// 操作状态
		if gconv.Int(res["code"]) == 0 {
			data.Status = 1
		} else {
			data.Status = 0
		}
		if _, ok := res["data"]; ok {
			delete(res, "data")
		}
		// 返回参数
		jsonRes, _ := gjson.Encode(res)
		if len(jsonRes) > 0 {
			data.JsonResult = string(jsonRes)
		}
		delete(params, "apiReturnResp")
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
	model := dao.OperLog.Ctx(ctx)
	columns := dao.OperLog.Columns()
	order := "id DESC"
	if req.Title != "" {
		model = model.WhereLike(columns.Title, "%"+req.Title+"%")
	}
	if req.OperName != "" {
		model = model.WhereLike(columns.OperName, "%"+req.OperName+"%")
	}
	if req.ReqMethod != "" {
		model = model.Where(columns.ReqMethod, req.ReqMethod)
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if req.StartTime.String() != "" {
		model = model.WhereGTE(columns.Time, req.StartTime)
	}
	if req.EndTime.String() != "" {
		model = model.WhereLTE(columns.Time, req.EndTime)
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	return
}
