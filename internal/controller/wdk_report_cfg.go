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
	WdkReportCfg = cWdkReportCfg{}
)

type cWdkReportCfg struct{}

// List 文档库报告类型配置列表
func (c *cWdkReportCfg) List(ctx context.Context, req *v1.WdkReportCfgListReq) (res *v1.WdkReportCfgListRes, err error) {
	var list []*v1.WdkReportCfgInfo
	list, err = service.WdkReportCfg().GetWdkReportCfgList(ctx, req.Name)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportCfgListFailed, err)
		return
	}

	res = &v1.WdkReportCfgListRes{List: list}
	return
}

// All 文档库全部报告类型配置列表
func (c *cWdkReportCfg) All(ctx context.Context, req *v1.WdkAllReportCfgReq) (res *v1.WdkAllReportCfgRes, err error) {
	var list []*entity.WdkReportCfg
	list, err = service.WdkReportCfg().GetWdkAllReportCfg(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkAllReportCfgFailed, err)
		return
	}

	res = &v1.WdkAllReportCfgRes{List: list}
	return
}

// Add 新增文档库报告类型配置
func (c *cWdkReportCfg) Add(ctx context.Context, req *v1.WdkReportCfgAddReq) (res *v1.WdkReportCfgAddRes, err error) {
	err = service.WdkReportCfg().AddWdkReportCfg(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkReportCfgFailed, err)
		return
	}

	return
}

// Info 文档库报告类型配置信息
func (c *cWdkReportCfg) Info(ctx context.Context, req *v1.WdkReportCfgInfoReq) (res *v1.WdkReportCfgInfoRes, err error) {
	var wdkReportCfg *v1.WdkReportCfgInfo
	wdkReportCfg, err = service.WdkReportCfg().GetWdkReportCfgById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportCfgFailed, err)
		return
	}

	res = &v1.WdkReportCfgInfoRes{Info: wdkReportCfg}
	return
}

// Edit 编辑文档库报告类型配置
func (c *cWdkReportCfg) Edit(ctx context.Context, req *v1.WdkReportCfgEditReq) (res *v1.WdkReportCfgEditRes, err error) {
	err = service.WdkReportCfg().EditWdkReportCfg(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditWdkReportCfgFailed, err)
		return
	}

	return
}

// Delete 删除文档库报告类型配置
func (c *cWdkReportCfg) Delete(ctx context.Context, req *v1.WdkReportCfgDeleteReq) (res *v1.WdkReportCfgDeleteRes, err error) {
	err = service.WdkReportCfg().DeleteWdkReportCfg(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteWdkReportCfgFailed, err)
		return
	}

	return
}
