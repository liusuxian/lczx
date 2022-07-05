package systemmonitor

import (
	"context"
	"github.com/gogf/gf/v2/os/grpool"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
)

type sLoginLog struct {
	pool            *grpool.Pool
	clientOptionMap map[string][]*model.ClientOption // 客户端选项
}

func init() {
	service.RegisterLoginLog(newLoginLog())
}

// 系统登录日志服务
func newLoginLog() *sLoginLog {
	insLoginLog := &sLoginLog{
		pool: grpool.New(100),
	}
	insLoginLog.clientOptionMap = map[string][]*model.ClientOption{}
	statusList := []*model.ClientOption{
		{
			Name:  "失败",
			Value: "0",
		},
		{
			Name:  "成功",
			Value: "1",
		},
	}
	insLoginLog.clientOptionMap["statusList"] = statusList

	return insLoginLog
}

// Invoke 异步保存日志
func (s *sLoginLog) Invoke(ctx context.Context, data *entity.LoginLog) {
	err := s.pool.Add(ctx, func(ctx context.Context) {
		// 保存日志
		s.SaveLoginLog(ctx, data)
	})
	if err != nil {
		logger.Error(ctx, "LoginLog Pool Add Error: ", err)
	}
}

// SaveLoginLog 保存日志
func (s *sLoginLog) SaveLoginLog(ctx context.Context, data *entity.LoginLog) {
	_, err := dao.LoginLog.Ctx(ctx).FieldsEx(dao.LoginLog.Columns().Id).Insert(data)
	if err != nil {
		logger.Error(ctx, "SaveLoginLog Error: ", err)
	}
}

// GetLoginLogList 获取登录日志列表
func (s *sLoginLog) GetLoginLogList(ctx context.Context, req *v1.LoginLogListReq) (total int, list []*entity.LoginLog, err error) {
	gmodel := dao.LoginLog.Ctx(ctx)
	columns := dao.LoginLog.Columns()
	order := "id DESC"
	if req.Passport != "" {
		gmodel = gmodel.WhereLike(columns.Passport, "%"+req.Passport+"%")
	}
	if req.Ip != "" {
		gmodel = gmodel.WhereLike(columns.Ip, "%"+req.Ip+"%")
	}
	if req.Location != "" {
		gmodel = gmodel.WhereLike(columns.Location, "%"+req.Location+"%")
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

// DeleteLoginLogByIds 通过ID列表删除登录日志
func (s *sLoginLog) DeleteLoginLogByIds(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.LoginLog.Ctx(ctx).WhereIn(dao.LoginLog.Columns().Id, ids).Delete()
	return
}

// ClearLoginLog 清除登录日志
func (s *sLoginLog) ClearLoginLog(ctx context.Context) (err error) {
	_, err = dao.LoginLog.DB().Exec(ctx, "truncate "+dao.LoginLog.Table())
	return
}

// GetClientOptionMap 获取客户端选项Map
func (s *sLoginLog) GetClientOptionMap() map[string][]*model.ClientOption {
	return s.clientOptionMap
}
