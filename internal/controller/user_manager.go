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

// GetEdit 获取被编辑用户的信息
func (c *cUserManager) GetEdit(ctx context.Context, req *v1.UserGetEditInfoReq) (res *v1.UserGetEditInfoRes, err error) {
	// 获取用户信息
	var user *entity.User
	user, err = service.User().GetUserById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetEditUserFailed, err)
		return
	}
	if user == nil {
		err = gerror.WrapCode(code.GetEditUserFailed, gerror.Newf(`用户ID[%d]不存在`, req.Id))
		return
	}
	user.Password = ""
	user.Salt = ""
	// 获取全部可用的角色
	var roles []*entity.Role
	roles, err = service.Role().GetEnableRoles(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetEditUserFailed, err)
		return
	}
	// 获取用户角色ID列表
	var roleIds []uint64
	roleIds, err = service.Role().GetUserRoleIds(ctx, user.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetEditUserFailed, err)
		return
	}

	res = &v1.UserGetEditInfoRes{
		User:     user,
		RoleList: roles,
		RoleIds:  roleIds,
	}
	return
}

// Edit 新增用户
func (c *cUserManager) Edit(ctx context.Context, req *v1.UserEditReq) (res *v1.UserEditRes, err error) {
	err = service.UserManager().EditUser(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditUserFailed, err)
		return
	}

	return
}

// ResetPwd 重置用户密码
func (c *cUserManager) ResetPwd(ctx context.Context, req *v1.UserResetPwdReq) (res *v1.UserResetPwdRes, err error) {
	err = service.UserManager().ResetUserPwd(ctx, req.Id, req.Password)
	if err != nil {
		err = gerror.WrapCode(code.ResetPwdFailed, err)
		return
	}

	return
}

// SetStatus 设置用户状态
func (c *cUserManager) SetStatus(ctx context.Context, req *v1.UserSetStatusReq) (res *v1.UserSetStatusRes, err error) {
	err = service.UserManager().SetUserStatus(ctx, req.Id, req.Status)
	if err != nil {
		err = gerror.WrapCode(code.SetUserStatusFailed, err)
		return
	}

	return
}

// Delete 删除用户
func (c *cUserManager) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	err = service.UserManager().DeleteUser(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteUserFailed, err)
		return
	}

	return
}

// SearchByRealname 通过姓名搜索用户
func (c *cUserManager) SearchByRealname(ctx context.Context, req *v1.UserSearchByRealnameReq) (res *v1.UserSearchByRealnameRes, err error) {
	var userList []*entity.User
	userList, err = service.UserManager().SearchByRealname(ctx, req.Realname)
	if err != nil {
		err = gerror.WrapCode(code.SearchByRealnameFailed, err)
		return
	}

	res = &v1.UserSearchByRealnameRes{List: userList}
	return
}
