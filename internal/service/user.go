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

func User() *sUser {
	return &insUser
}

// GetUserByPassportAndPassword 通过账号和密码获取用户信息
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
		Password: password,
		Status:   0,
	}).Scan(&user)
	return user, err
}

// CreateUser 创建用户
func (s *sUser) CreateUser(ctx context.Context, req *v1.UserCreateReq) (err error) {
	var available bool
	available, err = s.IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: req.Passport,
			Password: "12345678",
			Realname: req.Realname,
			Nickname: "",
			Gender:   req.Gender,
			Avatar:   "",
			Mobile:   req.Mobile,
			DeptId:   req.DeptId,
			RoleId:   req.RoleId,
			Status:   0,
		}).Insert()
		return err
	})
}

func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
