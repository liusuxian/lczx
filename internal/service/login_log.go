package service

import (
	"context"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
)

type sLoginLog struct {
	pool *grpool.Pool
}

var (
	insLoginLog = sLoginLog{
		pool: grpool.New(100),
	}
)

// LoginLog 系统登录日志服务
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
		logger.Error(ctx, "LoginLog Pool Add Error: ", err.Error())
	}
}

// SaveLoginLog 保存日志
func (s *sLoginLog) SaveLoginLog(ctx context.Context, data *entity.LoginLog) {
	_, err := dao.LoginLog.Ctx(ctx).FieldsEx(dao.LoginLog.Columns().Id).Insert(data)
	if err != nil {
		logger.Error(ctx, "SaveLoginLog Error: ", err.Error())
	}
}

// GetLoginLogList 获取登录日志列表
func (s *sLoginLog) GetLoginLogList(ctx context.Context, req *v1.LoginLogListReq) (total int, list []*entity.LoginLog, err error) {
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
