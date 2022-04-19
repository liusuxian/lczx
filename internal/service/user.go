package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sUser struct{}

var (
	insUser = sUser{}
)

// User 用户服务
func User() *sUser {
	return &insUser
}

// GetUserByPassportAndPassword 通过账号和密码获取用户信息
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string, fieldNames ...string) (user *entity.User, err error) {
	user, err = s.GetUserByPassport(ctx, passport, fieldNames...)
	if err != nil {
		return
	}
	if user == nil {
		return nil, gerror.New("账号不存在")
	}
	// 验证密码
	if utils.EncryptPassword(password, user.Salt) != user.Password {
		return nil, gerror.New("密码错误")
	}
	// 账号状态
	if user.Status == consts.UserStatusDisabled {
		return nil, gerror.New("账号已被禁用")
	}
	return
}

// UpdateUserLogin 更新用户登录信息
func (s *sUser) UpdateUserLogin(ctx context.Context, id uint64, ip string) {
	_, err := dao.User.Ctx(ctx).Unscoped().Data(do.User{
		LastLoginIp:   ip,
		LastLoginTime: gtime.Now(),
	}).Where(do.User{Id: id}).Update()
	if err != nil {
		logger.Error(ctx, "UpdateUserLogin Error: ", err.Error())
	}
}

// GetUserByPassport 通过账号获取用户信息
func (s *sUser) GetUserByPassport(ctx context.Context, passport string, fieldNames ...string) (user *entity.User, err error) {
	model := dao.User.Ctx(ctx)
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
	}
	err = model.Where(do.User{Passport: passport}).Scan(&user)
	return
}

// SetAvatar 设置用户头像
func (s *sUser) SetAvatar(ctx context.Context, id uint64, avatarUrl string) (err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{Avatar: avatarUrl}).Where(do.User{Id: id}).Update()
	return
}

// ProfileEdit 编辑个人中心信息
func (s *sUser) ProfileEdit(ctx context.Context, req *v1.UserProfileEditReq) (err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Realname: req.Realname,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Gender:   req.Gender,
	}).Where(do.User{Id: req.Id}).Update()
	return
}
