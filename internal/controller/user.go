package controller

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/response"
)

var (
	User = cUser{}
)

type cUser struct{}

// GetUserInfo 获取用户信息
func (c *cUser) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	userInfo := service.Context().Get(ctx).User
	if userInfo != nil && userInfo.Id > 0 {
		response.JsonOK(ctx, &v1.UserInfoRes{
			UserInfo: &v1.UserInfo{
				Id:       userInfo.Id,
				Passport: userInfo.Passport,
				Realname: userInfo.Realname,
				Nickname: userInfo.Nickname,
				Gender:   userInfo.Gender,
				Avatar:   userInfo.Avatar,
				Mobile:   userInfo.Mobile,
				DeptId:   userInfo.DeptId,
				RoleId:   userInfo.RoleId,
			},
		})
		return
	}

	response.Json(ctx, consts.CodeGetUserFailed, nil)
	return
}

// UserCreate 创建用户
func (c *cUser) UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	logger.Debug(ctx, "UserCreate Req: ", req)
	err = service.User().CreateUser(ctx, req)
	if err != nil {
		logger.Error(ctx, "UserCreate Error: ", err.Error())
		response.Json(ctx, consts.CodeAddUserFailed, nil)
		return
	}

	response.JsonOK(ctx, &v1.UserCreateRes{})
	return
}
