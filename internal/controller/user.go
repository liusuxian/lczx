package controller

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

// Info 获取用户信息
func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	user := service.Context().Get(ctx).User
	res = &v1.UserInfoRes{
		User: &entity.User{
			Id:            user.Id,
			Passport:      user.Passport,
			Realname:      user.Realname,
			Nickname:      user.Nickname,
			Gender:        user.Gender,
			Avatar:        user.Avatar,
			Mobile:        user.Mobile,
			DeptId:        user.DeptId,
			Status:        user.Status,
			Email:         user.Email,
			LastLoginIp:   user.LastLoginIp,
			LastLoginTime: user.LastLoginTime,
			CreateAt:      user.CreateAt,
		},
	}
	return
}
