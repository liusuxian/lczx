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
	OperLog = cOperLog{}
)

type cOperLog struct{}

// ClientOptions 客户端选项
func (c *cOperLog) ClientOptions(ctx context.Context, req *v1.OperLogClientOptionsReq) (res *v1.OperLogClientOptionsRes, err error) {
	// 获取客户端选项Map
	clientOptionMap := service.OperLog().GetClientOptionMap()

	res = &v1.OperLogClientOptionsRes{
		StatusList:   clientOptionMap["statusList"],
		OperTypeList: clientOptionMap["operTypeList"],
	}
	return
}

// List 操作日志列表
func (c *cOperLog) List(ctx context.Context, req *v1.OperLogListReq) (res *v1.OperLogListRes, err error) {
	var total int
	var list []*entity.OperLog
	total, list, err = service.OperLog().GetOperLogList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetOperLogListFailed, err)
		return
	}

	res = &v1.OperLogListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// Delete 删除操作日志
func (c *cOperLog) Delete(ctx context.Context, req *v1.OperLogDeleteReq) (res *v1.OperLogDeleteRes, err error) {
	err = service.OperLog().DeleteOperLogByIds(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteOperLogFailed, err)
		return
	}

	return
}

// Clear 清除操作日志
func (c *cOperLog) Clear(ctx context.Context, req *v1.OperLogClearReq) (res *v1.OperLogClearRes, err error) {
	err = service.OperLog().ClearOperLog(ctx)
	if err != nil {
		err = gerror.WrapCode(code.ClearOperLogFailed, err)
		return
	}

	return
}
