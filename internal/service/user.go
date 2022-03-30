package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sUser struct{}

var insUser = sUser{}

// User 用户管理服务
func User() *sUser {
	return &insUser
}

// IsPassportAvailable 用户账号是否可用
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// UserExistsById 通过用户ID判断用户信息是否存在
func (s *sUser) UserExistsById(ctx context.Context, id uint64) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Id: id}).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// GetUserById 通过用户ID获取用户信息
func (s *sUser) GetUserById(ctx context.Context, id uint64) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Id: id}).Scan(&user)
	return
}

// GetUserByPassport 通过账号获取用户信息
func (s *sUser) GetUserByPassport(ctx context.Context, passport string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Scan(&user)
	return
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
	return
}

// AddUser 新增用户
func (s *sUser) AddUser(ctx context.Context, req *v1.UserAddReq) (id int64, err error) {
	var available bool
	available, err = s.IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`账号[%s]已存在`, req.Passport)
		return
	}
	var deptExists bool
	deptExists, err = Dept().DeptExistsById(ctx, req.DeptId)
	if err != nil {
		return
	}
	if !deptExists {
		err = gerror.Newf(`部门ID[%d]不存在`, id)
		return
	}
	// 生成随机密码盐
	salt := grand.S(10)
	// 密码加密
	password := utils.EncryptPassword(req.Passport, salt)
	model := dao.User.Ctx(ctx)
	err = model.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		id, err = model.Data(do.User{
			Passport: req.Passport,
			Password: password,
			Salt:     salt,
			Realname: req.Realname,
			Gender:   req.Gender,
			Mobile:   req.Mobile,
			DeptId:   req.DeptId,
			RoleId:   req.RoleId,
			Status:   consts.UserStatusEnable,
		}).FieldsEx(dao.User.Columns().Id).InsertAndGetId()
		return err
	})
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
