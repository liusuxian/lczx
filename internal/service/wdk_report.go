package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/internal/upload"
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

// AddWdkReport 检查新增文档库上传报告记录权限
func (s *sWdkReport) AddWdkReport(ctx context.Context, req *v1.WdkReportAddReq, report *upload.FileInfo) (err error) {
	err = dao.WdkReport.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查新增文档库上传报告记录权限
		var terr error
		terr = s.AuthAdd(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}

		return nil
	})
	return
}

// AuthAdd 检查新增文档库上传报告记录权限
func (s *sWdkReport) AuthAdd(ctx context.Context, projectId uint64) (err error) {
	// 通过文档库项目ID判断文档库项目信息是否存在
	var wdkProject *entity.WdkProject
	wdkProject, err = WdkProject().GetWdkProjectById(ctx, projectId)
	if err != nil {
		return
	}
	if wdkProject == nil {
		err = gerror.Newf(`文档库项目ID[%d]不存在`, projectId)
		return
	}
	// 判断写入权限
	user := Context().Get(ctx).User
	if user.Id != wdkProject.PrincipalUid && user.IsAdmin != 1 {
		err = gerror.New("抱歉！！！该项目您没有上传报告的权限")
		return
	}
	return
}

// GetWdkReportCountByProjectId 通过项目ID获取文档库项目上传报告记录数量
func (s *sWdkReport) GetWdkReportCountByProjectId(ctx context.Context, projectId uint64) (count int, err error) {
	count, err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).Count()
	return
}
