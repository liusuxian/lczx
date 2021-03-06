package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

// Info 获取用户信息
func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 用户信息
	var user *model.ContextUser
	user, err = service.Context().GetUser(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetUserFailed, err)
		return
	}
	var userInfo *entity.User
	userInfo, err = service.User().GetUserById(ctx, user.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetUserFailed, err)
		return
	}
	// 获取头像访问url
	userInfo.Avatar, err = service.AliyunOSS().GetAccessUrl(ctx, userInfo.Avatar)
	if err != nil {
		err = gerror.WrapCode(code.GetUserFailed, err)
		return
	}
	userInfo.Password = ""
	userInfo.Salt = ""
	// 获取用户角色ID列表
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
		Roles:    roleNames,
	}
	return
}

// Profile 获取个人中心信息
func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	var profileInfo *v1.UserProfileInfo
	profileInfo, err = service.User().GetProfile(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetUserProfileFailed, err)
		return
	}

	res = &v1.UserProfileRes{ProfileInfo: profileInfo}
	return
}

// UploadAvatar 上传用户头像
func (c *cUser) UploadAvatar(ctx context.Context, req *v1.UserUploadAvatarReq) (res *v1.UserUploadAvatarRes, err error) {
	// 上传头像
	var fileInfo *model.UploadFileInfo
	fileInfo, err = service.AliyunOSS().UploadImg(req.AvatarFile, "user/avatar")
	if err != nil {
		err = gerror.WrapCode(code.SetUserAvatarFailed, err)
		return
	}
	// 设置用户头像
	err = service.User().SetAvatar(ctx, fileInfo.OriginFileUrl)
	if err != nil {
		err = gerror.WrapCode(code.SetUserAvatarFailed, err)
		return
	}
	// 获取头像访问url
	fileInfo.OriginFileUrl, err = service.AliyunOSS().GetAccessUrl(ctx, fileInfo.OriginFileUrl)
	if err != nil {
		err = gerror.WrapCode(code.SetUserAvatarFailed, err)
		return
	}

	res = &v1.UserUploadAvatarRes{FileInfo: fileInfo}
	return
}

// ProfileEdit 编辑个人中心信息
func (c *cUser) ProfileEdit(ctx context.Context, req *v1.UserProfileEditReq) (res *v1.UserProfileEditRes, err error) {
	err = service.User().EditProfile(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditUserProfileFailed, err)
		return
	}

	return
}

// PwdEdit 修改用户密码
func (c *cUser) PwdEdit(ctx context.Context, req *v1.UserPwdEditReq) (res *v1.UserPwdEditRes, err error) {
	err = service.User().EditPwd(ctx, req.OldPassword, req.NewPassword)
	if err != nil {
		err = gerror.WrapCode(code.EditPwdFailed, err)
		return
	}

	return
}
