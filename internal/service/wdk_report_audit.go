package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
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

// HandleWdkReportAudit 处理文档库报告审核
func (s *sWdkReportAudit) HandleWdkReportAudit(ctx context.Context, req *v1.WdkReportAuditReq) (err error) {
	user := Context().Get(ctx).User
	err = dao.WdkReportAudit.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 通过报告ID和用户ID判断文档库报告审核信息是否存在
		var terr error
		var wdkReportAuditExists bool
		wdkReportAuditExists, terr = s.WdkReportAuditExistsByIdAndUserId(ctx, req.Id, user.Id)
		if terr != nil {
			return terr
		}
		if !wdkReportAuditExists {
			return gerror.Newf(`文档库报告审核信息[%d]不存在`, req.Id)
		}
		// 设置文档库报告审核状态
		auditTime := gtime.Now()
		terr = s.SetWdkReportAuditStatus(ctx, req.Id, user.Id, req.AuditStatus, auditTime)
		if terr != nil {
			return terr
		}

		// 设置所属文档库项目阶段
		terr = WdkProject().SetWdkProjectStep(ctx, req.ProjectId, req.Type)
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目文件上传状态为是
		terr = WdkProject().SetWdkProjectFileUploadStatus(ctx, req.ProjectId)
		return nil
	})

	return
}

// WdkReportAuditExistsByIdAndUserId 通过报告ID和用户ID判断文档库报告审核信息是否存在
func (s *sWdkReportAudit) WdkReportAuditExistsByIdAndUserId(ctx context.Context, id uint64, userId uint64) (bool, error) {
	count, err := dao.WdkReportAudit.Ctx(ctx).Where(do.WdkReportAudit{Id: id, AuditUid: userId}).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// SetWdkReportAuditStatus 设置文档库报告审核状态
func (s *sWdkReportAudit) SetWdkReportAuditStatus(ctx context.Context, id uint64, userId uint64, status uint, auditTime *gtime.Time) (err error) {
	_, err = dao.WdkReportAudit.Ctx(ctx).Data(do.WdkReportAudit{
		Status:    status,
		AuditTime: auditTime,
	}).Where(do.WdkReportAudit{
		Id:       id,
		AuditUid: userId,
	}).Update()
	return
}
