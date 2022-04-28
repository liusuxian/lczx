package service

import (
	"github.com/gogf/gf/v2/database/gdb"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sWdkReport struct{}

var (
	insWdkReport = sWdkReport{}
)

// WdkReport 文档库上传报告记录管理服务
func WdkReport() *sWdkReport {
	return &insWdkReport
}

// GetWdkReportRecord 获取文档库上传报告记录
func (s *sWdkReport) GetWdkReportRecord(ctx context.Context, projectId uint64) (list []*v1.WdkReportInfo, err error) {
	err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{
		ProjectId:   projectId,
		AuditStatus: 2,
	}).ScanList(&list, "Report")
	if err != nil {
		return
	}
	err = dao.WdkReportType.Ctx(ctx).Where(dao.WdkReportType.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportType", "Report", "Id:Id")
	return
}

// GetWdkReportCountByProjectId 通过项目ID获取文档库项目上传报告记录数量
func (s *sWdkReport) GetWdkReportCountByProjectId(ctx context.Context, projectId uint64) (count int, err error) {
	count, err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).Count()
	return
}
