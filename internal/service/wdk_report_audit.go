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
	user := Context().Get(ctx).User
	model := dao.WdkReportAudit.Ctx(ctx).Where(do.WdkReportAudit{AuditUid: user.Id, Status: req.Status})
	columns := dao.WdkReportAudit.Columns()
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderAsc(columns.CreateAt).ScanList(&list, "ReportAudit")
	if err != nil {
		return
	}
	err = dao.WdkReport.Ctx(ctx).Where(dao.WdkReport.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAudit", "Id")).
		ScanList(&list, "Report", "ReportAudit", "Id:Id")
	if err != nil {
		return
	}
	err = dao.WdkProject.Ctx(ctx).Where(dao.WdkProject.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "ProjectId")).
		ScanList(&list, "Project", "Report", "Id:ProjectId")
	if err != nil {
		return
	}
	err = dao.WdkReportAuditType.Ctx(ctx).Where(dao.WdkReportAuditType.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAudit", "Id")).
		Where(do.WdkReportAuditType{AuditUid: user.Id}).ScanList(&list, "ReportAuditType", "ReportAudit", "Id:Id")
	return
}

// GetWdkReportBeAuditedList 获取文档库报告被审核列表
func (s *sWdkReportAudit) GetWdkReportBeAuditedList(ctx context.Context, req *v1.WdkReportBeAuditedListReq) (total int, list []*v1.WdkReportBeAuditedInfo, err error) {
	user := Context().Get(ctx).User
	model := dao.WdkReport.Ctx(ctx).Where(do.WdkReport{CreateBy: user.Id, AuditStatus: req.Status})
	columns := dao.WdkReport.Columns()
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderDesc(columns.CreateAt).ScanList(&list, "Report")
	if err != nil {
		return
	}
	err = dao.WdkProject.Ctx(ctx).Where(dao.WdkProject.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "ProjectId")).
		ScanList(&list, "Project", "Report", "Id:ProjectId")
	if err != nil {
		return
	}
	err = dao.WdkReportAudit.Ctx(ctx).Where(dao.WdkReportAudit.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportAudit", "Report", "Id:Id")
	if err != nil {
		return
	}
	err = dao.WdkReportAuditType.Ctx(ctx).Where(dao.WdkReportAuditType.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportAuditType", "Report", "Id:Id")
	return
}
