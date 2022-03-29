package service

import (
	"context"
	"github.com/gogf/gf/v2/os/grpool"
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
