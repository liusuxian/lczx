package service

import (
	"golang.org/x/net/context"
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

// GetWdkReportCountByProjectId 通过项目ID获取文档库项目上传报告记录数量
func (s *sWdkReport) GetWdkReportCountByProjectId(ctx context.Context, projectId uint64) (count int, err error) {
	count, err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).Count()
	return
}
