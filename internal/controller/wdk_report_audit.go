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
