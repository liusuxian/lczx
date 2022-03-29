package service

import (
	"context"
	"github.com/gogf/gf/v2/os/grpool"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/logger"
)

type sUserOnline struct {
	Pool *grpool.Pool
}

var insUserOnline = sUserOnline{
	Pool: grpool.New(100),
}

// UserOnline 用户在线表管理服务
func UserOnline() *sUserOnline {
	return &insUserOnline
}

// Invoke 异步保存用户在线状态
func (s *sUserOnline) Invoke(ctx context.Context, data *entity.UserOnline) {
	err := s.Pool.Add(ctx, func(ctx context.Context) {
		// 保存用户在线状态
		s.SaveUserOnline(ctx, data)
	})
	if err != nil {
		logger.Error(ctx, "UserOnline Pool Add Error: ", err.Error())
	}
}

// SaveUserOnline 保存用户在线状态
func (s *sUserOnline) SaveUserOnline(ctx context.Context, data *entity.UserOnline) {
	// 查询是否已存在当前用户
	var info *entity.UserOnline
	var err error
	err = dao.UserOnline.Ctx(ctx).Fields(dao.UserOnline.Columns().Id).Where(do.UserOnline{Token: data.Token}).Scan(&info)
	if err != nil {
		logger.Error(ctx, "SaveUserOnline Error: ", err.Error())
	}
	if info != nil {
		// 已存在则更新
		_, err = dao.UserOnline.Ctx(ctx).Where(do.UserOnline{Id: info.Id}).FieldsEx(dao.UserOnline.Columns().Id).Update(data)
		if err != nil {
			logger.Error(ctx, "SaveUserOnline Update Error: ", err.Error())
		}
	} else {
		// 不存在则新增
		_, err = dao.UserOnline.Ctx(ctx).FieldsEx(dao.UserOnline.Columns().Id).Insert(data)
		if err != nil {
			logger.Error(ctx, "SaveUserOnline Insert Error: ", err.Error())
		}
	}
}
