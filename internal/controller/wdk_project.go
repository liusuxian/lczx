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
	WdkProject = cWdkProject{}
)

type cWdkProject struct{}

// List 文档库项目列表
func (c *cWdkProject) List(ctx context.Context, req *v1.WdkProjectListReq) (res *v1.WdkProjectListRes, err error) {
	var total int
	var list []*entity.WdkProject
	total, list, err = service.WdkProject().GetWdkProjectList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkProjectListFailed, err)
		return
	}

	res = &v1.WdkProjectListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// Add 新增文档库项目
func (c *cWdkProject) Add(ctx context.Context, req *v1.WdkProjectAddReq) (res *v1.WdkProjectAddRes, err error) {
	err = service.WdkProject().AddWdkProject(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkProjectFailed, err)
		return
	}

	return
}

// Info 文档库项目信息
func (c *cWdkProject) Info(ctx context.Context, req *v1.WdkProjectInfoReq) (res *v1.WdkProjectInfoRes, err error) {
	var wdkProject *entity.WdkProject
	wdkProject, err = service.WdkProject().GetWdkProjectById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkProjectFailed, err)
		return
	}

	res = &v1.WdkProjectInfoRes{Info: wdkProject}
	return
}

// Edit 编辑文档库项目
func (c *cWdkProject) Edit(ctx context.Context, req *v1.WdkProjectEditReq) (res *v1.WdkProjectEditRes, err error) {
	err = service.WdkProject().EditWdkProject(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditWdkProjectFailed, err)
		return
	}

	return
}

// Delete 删除文档库项目
func (c *cWdkProject) Delete(ctx context.Context, req *v1.WdkProjectDeleteReq) (res *v1.WdkProjectDeleteRes, err error) {
	err = service.WdkProject().DeleteWdkProject(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteWdkProjectFailed, err)
		return
	}

	return
}
