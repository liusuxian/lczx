package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	UserOnline = cUserOnline{}
)

type cUserOnline struct{}

// List 在线用户列表
func (c *cUserOnline) List(ctx context.Context, req *v1.UserOnlineListReq) (res *v1.UserOnlineListRes, err error) {
	var total int
	var list []*entity.UserOnline
	total, list, err = service.UserOnline().GetOnlineList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetUserOnlineListFailed, err)
		return
	}

	res = &v1.UserOnlineListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// ForceLogout 强退用户
func (c *cUserOnline) ForceLogout(ctx context.Context, req *v1.UserOnlineForceLogoutReq) (res *v1.UserOnlineForceLogoutRes, err error) {
	var tokens []string
	tokens, err = service.UserOnline().GetOnlineTokensByIds(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.ForceLogoutUserOnlineFailed, err)
	}

	for _, token := range tokens {
		service.Auth().Token().RemoveToken(ctx, token)
	}
	return
}
