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

// GetWdkReportAuditList 获取文档库报告审核
func (s *sWdkReportAudit) GetWdkReportAuditList(ctx context.Context, req *v1.WdkReportAuditListReq) (total int, list []*v1.WdkReportAuditInfo, err error) {
	model := dao.WdkReportAuditRecord.Ctx(ctx)
	columns := dao.WdkReportAuditRecord.Columns()
	user := Context().Get(ctx).User
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderAsc(columns.CreateAt).Where(do.WdkReportAuditRecord{
		AuditUid: user.Id,
		Status:   req.Status,
	}).ScanList(&list, "ReportAuditRecord")
	if err != nil {
		return
	}
	err = dao.WdkReport.Ctx(ctx).Where(dao.WdkReport.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAuditRecord", "ReportId")).
		ScanList(&list, "Report", "ReportAuditRecord", "Id:ReportId")
	if err != nil {
		return
	}
	for _, v := range list {
		err = dao.WdkReportAuditType.Ctx(ctx).Where(do.WdkReportAuditType{
			AuditUid: v.ReportAuditRecord.AuditUid,
			ReportId: v.ReportAuditRecord.ReportId,
		}).Scan(&v.ReportAuditType)
		if err != nil {
			return
		}
	}
	return
}
