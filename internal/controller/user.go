package controller

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/utility/logger"
)

var (
	User = cUser{}
)

type cUser struct{}

// GetUserInfo 获取用户信息
func (c *cUser) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	logger.Debug(ctx, "GetUserInfo req: ", req)
	return res, nil
}

// GetUserDept 获取用户部门信息
func (c *cUser) GetUserDept(ctx context.Context, req *v1.UserDeptReq) (res *v1.UserDeptRes, err error) {
	logger.Debug(ctx, "GetUserDept req: ", req)
	return res, nil
}
