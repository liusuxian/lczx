package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/code"
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
		User: &v1.UserInfo{
			Id:       user.Id,
			Passport: user.Passport,
			Realname: user.Realname,
			Nickname: user.Nickname,
			Gender:   user.Gender,
			Avatar:   user.Avatar,
			Mobile:   user.Mobile,
			DeptId:   user.DeptId,
			RoleId:   user.RoleId,
		},
	}
	return
}

// Add 新增用户
func (c *cUser) Add(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error) {
	var id int64
	id, err = service.User().AddUser(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddUserFailed, err)
		return
	}

	var user *entity.User
	newId := gconv.Uint(id)
	user, err = service.User().GetUserById(ctx, newId)
	if err != nil {
		err = gerror.WrapCode(code.AddUserFailed, err)
		return
	}

	res = &v1.UserAddRes{
		User: &v1.UserInfo{
			Id:       user.Id,
			Passport: user.Passport,
			Realname: user.Realname,
			Nickname: user.Nickname,
			Gender:   user.Gender,
			Avatar:   user.Avatar,
			Mobile:   user.Mobile,
			DeptId:   user.DeptId,
			RoleId:   user.RoleId,
		},
	}
	return
}
