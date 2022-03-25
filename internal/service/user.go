package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
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
func (s *sUser) UserExistsById(ctx context.Context, id uint) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Id: id}).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// GetUserByPassportAndPassword 通过账号和密码获取用户信息
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Passport: passport, Password: password}).Scan(&user)
	return
}

// GetUserById 通过用户ID获取用户信息
func (s *sUser) GetUserById(ctx context.Context, id uint) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{Id: id}).Scan(&user)
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
	err = dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		id, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: req.Passport,
			Password: req.Passport,
			Realname: req.Realname,
			Gender:   req.Gender,
			Mobile:   req.Mobile,
			DeptId:   req.DeptId,
			RoleId:   req.RoleId,
			Status:   0,
		}).InsertAndGetId()
		return err
	})
	return
}
