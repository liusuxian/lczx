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
	// 用户信息
	userInfo := &entity.User{
		Id:            user.Id,
		Passport:      user.Passport,
		Realname:      user.Realname,
		Nickname:      user.Nickname,
		Gender:        user.Gender,
		Avatar:        user.Avatar,
		Mobile:        user.Mobile,
		DeptId:        user.DeptId,
		Status:        user.Status,
		IsAdmin:       user.IsAdmin,
		Email:         user.Email,
		Remark:        user.Remark,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		CreateAt:      user.CreateAt,
		UpdateAt:      user.UpdateAt,
	}
	// 获取用户角色
	var roles []*entity.Role
	roles, err = service.Role().GetUserRoles(ctx, user.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetUserFailed, err)
		return
	}
	roleNames := make([]string, len(roles))
	roleIds := make([]uint64, len(roles))
	for k, v := range roles {
		roleNames[k] = v.Name
		roleIds[k] = v.Id
	}
	// 获取用户菜单
	var menuList []string
	tagSuperUser := false
	service.UserManager().NotCheckAuthUserIds().Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == user.Id {
			tagSuperUser = true
			return false
		}
		return true
	})
	if tagSuperUser {
		menuList = []string{"*/*/*"}
	} else {
		menuList, err = service.Role().GetUserMenuList(ctx, roleIds)
		if err != nil {
			err = gerror.WrapCode(code.GetUserFailed, err)
			return
		}
	}

	res = &v1.UserInfoRes{
		User:         userInfo,
		RoleNameList: roleNames,
		MenuList:     menuList,
	}
	return
}
