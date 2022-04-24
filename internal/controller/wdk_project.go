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
