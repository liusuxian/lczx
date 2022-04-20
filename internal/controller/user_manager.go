package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	UserManager = cUserManager{}
)

type cUserManager struct{}

// List 用户列表
func (c *cUserManager) List(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	var total int
	var userList []*entity.User
	total, userList, err = service.UserManager().GetUserList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetUserListFailed, err)
		return
	}

	var profileInfos []*v1.UserProfileInfo
	profileInfos, err = service.UserManager().GetProfileList(ctx, userList)
	if err != nil {
		err = gerror.WrapCode(code.GetUserListFailed, err)
		return
	}

	res = &v1.UserListRes{
		CurPage:         req.CurPage,
		Total:           total,
		ProfileInfoList: profileInfos,
	}
	return
}

// Add 新增用户
func (c *cUserManager) Add(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error) {
	err = service.UserManager().AddUser(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddUserFailed, err)
		return
	}

	return
}
