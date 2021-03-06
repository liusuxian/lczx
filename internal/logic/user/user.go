package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sUser struct{}

func init() {
	service.RegisterUser(newUser())
}

// 用户服务
func newUser() *sUser {
	return &sUser{}
}

// GetUserByPassportAndPassword 通过账号和密码获取用户信息
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	user, err = s.GetUserByPassport(ctx, passport)
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
	if user.Status == 0 {
		return nil, gerror.New("账号已被禁用")
	}
	user.Password = ""
	user.Salt = ""
	return
}

// UpdateUserLogin 更新用户登录信息
func (s *sUser) UpdateUserLogin(ctx context.Context, id uint64, ip string) {
	_, err := dao.User.Ctx(ctx).Unscoped().Data(do.User{
		LastLoginIp:   ip,
		LastLoginTime: gtime.Now(),
	}).Where(do.User{Id: id}).Update()
	if err != nil {
		logger.Error(ctx, "UpdateUserLogin Error: ", err)
	}
}

// GetProfile 获取个人中心信息
func (s *sUser) GetProfile(ctx context.Context) (profileInfo *v1.UserProfileInfo, err error) {
	// 用户信息
	var curUser *model.ContextUser
	curUser, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
	var userInfo *entity.User
	userInfo, err = s.GetUserById(ctx, curUser.Id)
	if err != nil {
		return
	}
	if userInfo == nil {
		err = gerror.Newf(`用户ID[%d]不存在`, curUser.Id)
		return
	}
	// 获取头像访问url
	userInfo.Avatar, err = service.AliyunOSS().GetAccessUrl(ctx, userInfo.Avatar)
	if err != nil {
		return
	}
	userInfo.Password = ""
	userInfo.Salt = ""
	// 获取部门状态为正常的部门列表
	var depts []*entity.Dept
	depts, err = service.Dept().GetStatusEnableDepts(ctx)
	if err != nil {
		return
	}
	// 获取部门信息
	dept := service.Dept().GetDeptById(depts, userInfo.DeptId)
	// 获取用户角色
	var roles []*entity.Role
	roles, err = service.Role().GetUserRoles(ctx, userInfo.Id)
	if err != nil {
		return
	}

	profileInfo = &v1.UserProfileInfo{
		User:  userInfo,
		Dept:  dept,
		Roles: roles,
	}
	return
}

// SetAvatar 设置用户头像
func (s *sUser) SetAvatar(ctx context.Context, avatarUrl string) (err error) {
	var user *model.ContextUser
	user, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
	_, err = dao.User.Ctx(ctx).Data(do.User{Avatar: avatarUrl}).Where(do.User{Id: user.Id}).Update()
	return
}

// EditProfile 编辑个人中心信息
func (s *sUser) EditProfile(ctx context.Context, req *v1.UserProfileEditReq) (err error) {
	var user *model.ContextUser
	user, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Email:    req.Email,
	}).Where(do.User{Id: user.Id}).Update()
	return
}

// EditPwd 修改用户密码
func (s *sUser) EditPwd(ctx context.Context, oldPassword, newPassword string) (err error) {
	var curUser *model.ContextUser
	curUser, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
	var userInfo *entity.User
	userInfo, err = s.GetUserById(ctx, curUser.Id)
	if err != nil {
		return
	}
	if userInfo == nil {
		err = gerror.Newf(`用户ID[%d]不存在`, curUser.Id)
		return
	}
	oldEncryptPassword := utils.EncryptPassword(oldPassword, userInfo.Salt)
	if oldEncryptPassword != userInfo.Password {
		err = gerror.New("原始密码错误!")
		return
	}
	// 生成随机密码盐
	salt := grand.S(10)
	newEncryptPassword := utils.EncryptPassword(newPassword, salt)
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Salt:     salt,
		Password: newEncryptPassword,
	}).Where(do.User{Id: userInfo.Id}).Update()
	return
}

// GetUserByPassport 通过账号获取用户信息
func (s *sUser) GetUserByPassport(ctx context.Context, passport string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Scan(&user)
	return
}

// GetUserById 通过用户ID获取用户信息
func (s *sUser) GetUserById(ctx context.Context, id uint64) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Id: id}).Scan(&user)
	if err != nil {
		return
	}
	// 处理头像
	if user.Avatar == "" {
		user.Avatar = "user/avatar/logo.png"
	}
	return
}
