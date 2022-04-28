package controller

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	WdkReport = cWdkReport{}
)

type cWdkReport struct{}

// Record 文档库上传报告记录
func (c *cWdkReport) Record(ctx context.Context, req *v1.WdkReportRecordReq) (res *v1.WdkReportRecordRes, err error) {
	var list []*v1.WdkReportInfo
	list, err = service.WdkReport().GetWdkReportRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportRecordFailed, err)
		return
	}

	res = &v1.WdkReportRecordRes{List: list}
	return
}
