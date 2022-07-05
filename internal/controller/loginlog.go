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
	LoginLog = cLoginLog{}
)

type cLoginLog struct{}

// ClientOptions 客户端选项
func (c *cLoginLog) ClientOptions(ctx context.Context, req *v1.LoginLogClientOptionsReq) (res *v1.LoginLogClientOptionsRes, err error) {
	// 获取客户端选项Map
	clientOptionMap := service.LoginLog().GetClientOptionMap()

	res = &v1.LoginLogClientOptionsRes{
		StatusList: clientOptionMap["statusList"],
	}
	return
}

// List 登录日志列表
func (c *cLoginLog) List(ctx context.Context, req *v1.LoginLogListReq) (res *v1.LoginLogListRes, err error) {
	var total int
	var list []*entity.LoginLog
	total, list, err = service.LoginLog().GetLoginLogList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetLoginLogListFailed, err)
		return
	}

	res = &v1.LoginLogListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// Delete 删除登录日志
func (c *cLoginLog) Delete(ctx context.Context, req *v1.LoginLogDeleteReq) (res *v1.LoginLogDeleteRes, err error) {
	err = service.LoginLog().DeleteLoginLogByIds(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteLoginLogFailed, err)
		return
	}

	return
}

// Clear 清除登录日志
func (c *cLoginLog) Clear(ctx context.Context, req *v1.LoginLogClearReq) (res *v1.LoginLogClearRes, err error) {
	err = service.LoginLog().ClearLoginLog(ctx)
	if err != nil {
		err = gerror.WrapCode(code.ClearLoginLogFailed, err)
		return
	}

	return
}
