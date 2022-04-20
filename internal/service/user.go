package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/logger"
	"lczx/utility/utils"
	"time"
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
	if user.Status == consts.UserStatusDisabled {
		return nil, gerror.New("账号已被禁用")
	}
	user.Password = ""
	user.Salt = ""
	return
}

// UpdateUserLogin 更新用户登录信息
func (s *sUser) UpdateUserLogin(ctx context.Context, id uint64, ip string) {
	_, err := dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     s.UserCacheKey(id),
		Force:    false,
	}).Unscoped().Data(do.User{
		LastLoginIp:   ip,
		LastLoginTime: gtime.Now(),
	}).Where(do.User{Id: id}).Update()
	if err != nil {
		logger.Error(ctx, "UpdateUserLogin Error: ", err.Error())
	}
}

// GetProfile 获取个人中心信息
func (s *sUser) GetProfile(ctx context.Context, id uint64) (profileInfo *v1.UserProfileInfo, err error) {
	// 用户信息
	var user *entity.User
	user, err = s.GetUserById(ctx, id)
	if err != nil {
		return
	}
	user.Password = ""
	user.Salt = ""
	// 获取部门状态为正常的部门列表
	var depts []*entity.Dept
	depts, err = Dept().GetStatusEnableDepts(ctx)
	if err != nil {
		return
	}
	// 获取部门信息
	deptNames := Dept().GetDeptAllNameById(depts, user.DeptId)
	utils.Reverse(deptNames)
	dept := Dept().GetDeptById(depts, user.DeptId)
	dept.Name = gstr.Join(deptNames, "/")
	// 获取用户角色
	var roles []*entity.Role
	roles, err = Role().GetUserRoles(ctx, user.Id)
	if err != nil {
		return
	}

	profileInfo = &v1.UserProfileInfo{
		User:  user,
		Dept:  dept,
		Roles: roles,
	}
	return
}

// SetAvatar 设置用户头像
func (s *sUser) SetAvatar(ctx context.Context, id uint64, avatarUrl string) (err error) {
	_, err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     s.UserCacheKey(id),
		Force:    false,
	}).Data(do.User{Avatar: avatarUrl}).Where(do.User{Id: id}).Update()
	return
}

// EditProfile 编辑个人中心信息
func (s *sUser) EditProfile(ctx context.Context, id uint64, req *v1.UserProfileEditReq) (err error) {
	_, err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     s.UserCacheKey(id),
		Force:    false,
	}).Data(do.User{
		Realname: req.Realname,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Gender:   req.Gender,
	}).Where(do.User{Id: id}).Update()
	return
}

// EditPwd 修改用户密码
func (s *sUser) EditPwd(ctx context.Context, id uint64, oldPassword, newPassword string) (err error) {
	var user *entity.User
	user, err = s.GetUserById(ctx, id)
	if err != nil {
		return
	}
	oldEncryptPassword := utils.EncryptPassword(oldPassword, user.Salt)
	if oldEncryptPassword != user.Password {
		err = gerror.New("原始密码错误!")
		return
	}
	// 生成随机密码盐
	salt := grand.S(10)
	newEncryptPassword := utils.EncryptPassword(newPassword, salt)
	_, err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     s.UserCacheKey(id),
		Force:    false,
	}).Data(do.User{
		Salt:     salt,
		Password: newEncryptPassword,
	}).Where(do.User{Id: id}).Update()
	return
}

// GetUserByPassport 通过账号获取用户信息
func (s *sUser) GetUserByPassport(ctx context.Context, passport string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Scan(&user)
	return
}

// GetUserById 通过用户ID获取用户信息
func (s *sUser) GetUserById(ctx context.Context, id uint64) (user *entity.User, err error) {
	// 从缓存获取
	userCacheKey := s.UserCacheKey(id)
	userCacheVal := Cache().GetCache(ctx, userCacheKey)
	if userCacheVal != nil {
		var userList []*entity.User
		err = gconv.Structs(userCacheVal, &userList)
		if err != nil {
			return
		}
		if userList != nil && len(userList) > 0 {
			user = userList[0]
			return
		}
	}
	// 从数据库获取
	err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 2,
		Name:     userCacheKey,
		Force:    false,
	}).Where(do.User{Id: id}).Scan(&user)
	return
}

// UserCacheKey 用户缓存key
func (s *sUser) UserCacheKey(id uint64) string {
	return consts.CachePrefix + "user:" + gconv.String(id)
}
