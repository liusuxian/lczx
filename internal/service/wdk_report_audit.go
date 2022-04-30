package service

import (
	"github.com/gogf/gf/v2/database/gdb"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sWdkReportAudit struct{}

var (
	insWdkReportAudit = sWdkReportAudit{}
)

// WdkReportAudit 文档库报告审核记录管理服务
func WdkReportAudit() *sWdkReportAudit {
	return &insWdkReportAudit
}

// GetWdkReportAuditList 获取文档库报告审核列表
func (s *sWdkReportAudit) GetWdkReportAuditList(ctx context.Context, req *v1.WdkReportAuditListReq) (total int, list []*v1.WdkReportAuditInfo, err error) {
	model := dao.WdkReportAudit.Ctx(ctx)
	columns := dao.WdkReportAudit.Columns()
	user := Context().Get(ctx).User
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderAsc(columns.CreateAt).Where(do.WdkReportAudit{
		AuditUid: user.Id,
		Status:   req.Status,
	}).ScanList(&list, "ReportAudit")
	if err != nil {
		return
	}
	err = dao.WdkReport.Ctx(ctx).Where(dao.WdkReport.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAudit", "Id")).
		ScanList(&list, "Report", "ReportAudit", "Id:Id")
	if err != nil {
		return
	}
	err = dao.WdkReportAuditType.Ctx(ctx).Where(dao.WdkReportAuditType.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAudit", "Id")).
		Where(do.WdkReportAuditType{AuditUid: user.Id}).ScanList(&list, "ReportAuditType", "ReportAudit", "Id:Id")
	return
}
