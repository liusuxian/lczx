package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
)

var (
	User = cUser{}
)

type cUser struct{}

// Info 获取用户信息
func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	user := service.Context().Get(ctx).User
	// 用户信息
	var userInfo *entity.User
	userInfo, err = service.User().GetUserById(ctx, user.Id)
	userInfo.Password = ""
	userInfo.Salt = ""
	if err != nil {
		err = gerror.WrapCode(code.GetUserFailed, err)
		return
	}
	// 获取用户角色ID列表
	var roleIds []uint64
	roleIds, err = service.Role().GetUserRoleIds(ctx, user.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetUserFailed, err)
		return
	}
	// 获取用户菜单
	var menuList []string
	tagSuperUser := false
	service.UserManager().NotCheckAuthUserIds().Iterator(func(v any) bool {
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
		User:     userInfo,
		MenuList: menuList,
	}
	return
}

// Profile 获取个人中心信息
func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	user := service.Context().Get(ctx).User
	var profileInfo *v1.UserProfileInfo
	profileInfo, err = service.User().GetProfile(ctx, user.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetUserProfileFailed, err)
		return
	}

	res = &v1.UserProfileRes{ProfileInfo: profileInfo}
	return
}

// UploadAvatar 上传用户头像
func (c *cUser) UploadAvatar(ctx context.Context, req *v1.UserUploadAvatarReq) (res *v1.UserUploadAvatarRes, err error) {
	avatar := g.RequestFromCtx(ctx).GetUploadFile(req.UploadName)
	logger.Debug(ctx, "avatar: ", avatar)
	return
}

// ProfileEdit 编辑个人中心信息
func (c *cUser) ProfileEdit(ctx context.Context, req *v1.UserProfileEditReq) (res *v1.UserProfileEditRes, err error) {
	user := service.Context().Get(ctx).User
	err = service.User().EditProfile(ctx, user.Id, req)
	if err != nil {
		err = gerror.WrapCode(code.EditUserProfileFailed, err)
		return
	}

	return
}

// PwdEdit 修改用户密码
func (c *cUser) PwdEdit(ctx context.Context, req *v1.UserPwdEditReq) (res *v1.UserPwdEditRes, err error) {
	user := service.Context().Get(ctx).User
	err = service.User().EditPwd(ctx, user.Id, req.OldPassword, req.NewPassword)
	if err != nil {
		err = gerror.WrapCode(code.EditPwdFailed, err)
		return
	}

	return
}
