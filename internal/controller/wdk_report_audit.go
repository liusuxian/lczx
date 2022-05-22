package controller

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	WdkReportAudit = cWdkReportAudit{}
)

type cWdkReportAudit struct{}

// List 文档库报告审核记录列表
func (c *cWdkReportAudit) List(ctx context.Context, req *v1.WdkReportAuditListReq) (res *v1.WdkReportAuditListRes, err error) {
	var total int
	var list []*v1.WdkReportAuditInfo
	total, list, err = service.WdkReportAudit().GetWdkReportAuditList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportAuditListFailed, err)
		return
	}

	res = &v1.WdkReportAuditListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// Audit 文档库报告审核
func (c *cWdkReportAudit) Audit(ctx context.Context, req *v1.WdkReportAuditReq) (res *v1.WdkReportAuditRes, err error) {
	err = service.WdkReportAudit().HandleWdkReportAudit(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.HandleWdkReportAuditFailed, err)
		return
	}

	return
}

// UploadAuditList 文档库报告上传审核列表
func (c *cWdkReportAudit) UploadAuditList(ctx context.Context, req *v1.WdkReportUploadAuditListReq) (res *v1.WdkReportUploadAuditListRes, err error) {
	var total int
	var list []*v1.WdkReportUploadAuditInfo
	total, list, err = service.WdkReportAudit().GetWdkReportUploadAuditList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportUploadAuditListFailed, err)
		return
	}

	res = &v1.WdkReportUploadAuditListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// RescindAudit 文档库报告撤销审核
func (c *cWdkReportAudit) RescindAudit(ctx context.Context, req *v1.WdkReportRescindAuditReq) (res *v1.WdkReportRescindAuditRes, err error) {
	err = service.WdkReportAudit().HandleWdkReportRescindAudit(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.HandleWdkReportRescindAuditFailed, err)
		return
	}

	return
}

// AuditProcess 文档库报告审核流程
func (c *cWdkReportAudit) AuditProcess(ctx context.Context, req *v1.WdkReportAuditProcessReq) (res *v1.WdkReportAuditProcessRes, err error) {
	var list []*v1.WdkReportAuditProcessInfo
	list, err = service.WdkReportAudit().GetWdkReportAuditProcess(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportAuditProcessFailed, err)
		return
	}

	res = &v1.WdkReportAuditProcessRes{List: list}
	return
}
