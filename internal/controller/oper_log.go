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
