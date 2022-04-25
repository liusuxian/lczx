package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	WdkFiletype = cWdkFiletype{}
)

type cWdkFiletype struct{}

// List 文档库上传文件类型列表
func (c *cWdkFiletype) List(ctx context.Context, req *v1.WdkFiletypeListReq) (res *v1.WdkFiletypeListRes, err error) {
	var total int
	var list []*v1.WdkFiletypeInfo
	total, list, err = service.WdkFiletype().GetWdkFiletypeList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkFiletypeListFailed, err)
		return
	}

	res = &v1.WdkFiletypeListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}
