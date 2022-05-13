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

// BeAuditedList 文档库报告被审核记录列表
func (c *cWdkReportAudit) BeAuditedList(ctx context.Context, req *v1.WdkReportBeAuditedListReq) (res *v1.WdkReportBeAuditedListRes, err error) {
	var total int
	var list []*v1.WdkReportBeAuditedInfo
	total, list, err = service.WdkReportAudit().GetWdkReportBeAuditedList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportBeAuditedListFailed, err)
		return
	}

	res = &v1.WdkReportBeAuditedListRes{
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
