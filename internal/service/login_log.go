package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/grpool"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/utility/logger"
)

type sLoginLog struct {
	pool *grpool.Pool
}

var insLoginLog = sLoginLog{
	pool: grpool.New(100),
}

// LoginLog 系统登录日志管理服务
func LoginLog() *sLoginLog {
	return &insLoginLog
}

// Invoke 异步保存日志
func (s *sLoginLog) Invoke(ctx context.Context, data *entity.LoginLog) {
	err := s.pool.Add(ctx, func(ctx context.Context) {
		// 保存日志
		s.SaveLoginLog(ctx, data)
	})
	if err != nil {
		logger.Error(ctx, "SysLoginLog Pool Add Error: ", err.Error())
	}
}

// SaveLoginLog 保存日志
func (s *sLoginLog) SaveLoginLog(ctx context.Context, data *entity.LoginLog) {
	_, err := dao.LoginLog.Ctx(ctx).FieldsEx(dao.LoginLog.Columns().Id).Insert(data)
	if err != nil {
		logger.Error(ctx, "SaveLoginLog Error: ", err.Error())
	}
}

// LoginLogList 获取登录日志列表
func (s *sLoginLog) LoginLogList(ctx context.Context, req *v1.LoginLogListReq) (total int, list []*v1.LoginLogListRes, err error) {
	model := dao.LoginLog.Ctx(ctx)
	columns := dao.LoginLog.Columns()
	order := "id DESC"

	if req.Passport != "" {
		model = model.WhereLike(columns.Passport, "%"+req.Passport+"%")
	}
	if req.Ip != "" {
		model = model.WhereLike(columns.Ip, "%"+req.Ip+"%")
	}
	if req.Location != "" {
		model = model.WhereLike(columns.Location, "%"+req.Location+"%")
	}
	if req.Status == consts.LoginFailed || req.Status == consts.LoginSucc {
		model = model.Where(columns.Status, req.Status)
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
		err = gerror.Wrap(err, "获取总行数失败")
		return
	}
	err = model.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败")
	}
	return
}
