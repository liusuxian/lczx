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

// Add 新增文档库上传文件类型
func (c *cWdkFiletype) Add(ctx context.Context, req *v1.WdkFiletypeAddReq) (res *v1.WdkFiletypeAddRes, err error) {
	err = service.WdkFiletype().AddWdkFiletype(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkFiletypeFailed, err)
		return
	}

	return
}

// Info 文档库上传文件类型信息
func (c *cWdkFiletype) Info(ctx context.Context, req *v1.WdkFiletypeInfoReq) (res *v1.WdkFiletypeInfoRes, err error) {
	var wdkFiletype *v1.WdkFiletypeInfo
	wdkFiletype, err = service.WdkFiletype().GetWdkFiletypeById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkFiletypeFailed, err)
		return
	}

	res = &v1.WdkFiletypeInfoRes{Info: wdkFiletype}
	return
}

// Edit 编辑文档库项目
func (c *cWdkFiletype) Edit(ctx context.Context, req *v1.WdkFiletypeEditReq) (res *v1.WdkFiletypeEditRes, err error) {
	err = service.WdkFiletype().EditWdkFiletype(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditWdkFiletypeFailed, err)
		return
	}

	return
}

// Delete 删除文档库项目
func (c *cWdkFiletype) Delete(ctx context.Context, req *v1.WdkFiletypeDeleteReq) (res *v1.WdkFiletypeDeleteRes, err error) {
	err = service.WdkFiletype().DeleteWdkFiletype(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteWdkFiletypeFailed, err)
		return
	}

	return
}
