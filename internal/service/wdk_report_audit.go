package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
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
	model := dao.WdkReportAudit.Ctx(ctx).Where(do.WdkReportAudit{AuditUid: user.Id, Rescind: req.Rescind, PreauditStatus: 1, Status: req.Status})
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
	err = dao.WdkReportAuditType.Ctx(ctx).Where(dao.WdkReportAuditType.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAudit", "Id")).
		Where(do.WdkReportAuditType{AuditUid: user.Id}).ScanList(&list, "ReportAuditType", "ReportAudit", "Id:Id")
	return
}

// HandleWdkReportAudit 处理文档库报告审核
func (s *sWdkReportAudit) HandleWdkReportAudit(ctx context.Context, req *v1.WdkReportAuditReq) (err error) {
	user := Context().Get(ctx).User
	err = dao.WdkReportAudit.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 通过报告ID和用户ID判断文档库报告审核信息是否存在
		var terr error
		var wdkReportAudit *entity.WdkReportAudit
		wdkReportAudit, terr = s.GetWdkReportAuditByIdAndUid(ctx, req.Id, user.Id)
		if terr != nil {
			return terr
		}
		if wdkReportAudit == nil {
			return gerror.Newf(`文档库报告审核信息[%d]不存在`, req.Id)
		}
		// 设置文档库报告审核状态
		auditTime := gtime.Now()
		terr = s.SetWdkReportAuditStatus(ctx, wdkReportAudit.Id, wdkReportAudit.AuditUid, req.AuditStatus, req.Excellence, auditTime)
		if terr != nil {
			return terr
		}
		// 通过报告ID获取文档库报告审核列表
		var auditList []*entity.WdkReportAudit
		auditList, terr = s.GetWdkReportAuditListById(ctx, wdkReportAudit.Id)
		if terr != nil {
			return terr
		}
		// 获取文档库报告审核完成状态
		completeStatus := s.GetWdkReportAuditCompleteStatus(auditList)
		if completeStatus != 1 {
			// 获取文档库报告审核完成时的被推荐优秀报告状态
			excellence := s.GetWdkReportExcellence(auditList)
			// 设置文档库报告审核完成状态
			terr = WdkReport().SetWdkReportAuditCompleteStatus(ctx, wdkReportAudit.Id, completeStatus, excellence, auditTime)
			if terr != nil {
				return terr
			}
		}
		if completeStatus == 2 {
			// 设置所属文档库项目阶段
			terr = WdkProject().SetWdkProjectStep(ctx, wdkReportAudit.ProjectId, 8)
			if terr != nil {
				return terr
			}
			// 设置所属文档库项目文件上传状态为是
			terr = WdkProject().SetWdkProjectFileUploadStatus(ctx, wdkReportAudit.ProjectId)
			if terr != nil {
				return terr
			}
		}
		return nil
	})
	return
}

// GetWdkReportUploadAuditList 获取文档库报告上传审核列表
func (s *sWdkReportAudit) GetWdkReportUploadAuditList(ctx context.Context, req *v1.WdkReportUploadAuditListReq) (total int, list []*v1.WdkReportUploadAuditInfo, err error) {
	user := Context().Get(ctx).User
	model := dao.WdkReport.Ctx(ctx).Where(do.WdkReport{CreateBy: user.Id, Rescind: req.Rescind, AuditStatus: req.Status})
	columns := dao.WdkReport.Columns()
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderDesc(columns.CreateAt).ScanList(&list, "Report")
	if err != nil {
		return
	}
	err = dao.WdkReportType.Ctx(ctx).Where(dao.WdkReportType.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportType", "Report", "Id:Id")
	return
}

// HandleWdkReportRescindAudit 处理文档库报告撤销审核
func (s *sWdkReportAudit) HandleWdkReportRescindAudit(ctx context.Context, Id uint64) (err error) {
	err = dao.WdkReportAudit.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var terr error
		_, terr = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{Rescind: 1}).Where(do.WdkReport{Id: Id}).Update()
		if terr != nil {
			return terr
		}
		_, terr = dao.WdkReportAudit.Ctx(ctx).Data(do.WdkReportAudit{Rescind: 1}).Where(do.WdkReportAudit{Id: Id}).Update()
		return terr
	})
	return
}

// GetWdkReportAuditProcess 获取文档库报告审核流程
func (s *sWdkReportAudit) GetWdkReportAuditProcess(ctx context.Context, Id uint64) (list []*v1.WdkReportAuditProcessInfo, err error) {
	err = dao.WdkReportAudit.Ctx(ctx).Where(do.WdkReportAudit{Id: Id}).OrderAsc(dao.WdkReportAudit.Columns().AuditorType).
		OrderAsc(dao.WdkReportAudit.Columns().AuditTime).ScanList(&list, "ReportAudit")
	if err != nil {
		return
	}
	err = dao.WdkReportAuditType.Ctx(ctx).Where(dao.WdkReportAuditType.Columns().Id, gdb.ListItemValuesUnique(list, "ReportAudit", "AuditUid")).
		Where(do.WdkReportAuditType{Id: Id}).ScanList(&list, "ReportAuditType", "ReportAudit", "AuditUid:AuditUid")
	return
}

// GetWdkReportAuditByIdAndUid 通过报告ID和用户ID获取文档库报告审核信息
func (s *sWdkReportAudit) GetWdkReportAuditByIdAndUid(ctx context.Context, id uint64, userId uint64) (wdkReportAudit *entity.WdkReportAudit, err error) {
	err = dao.WdkReportAudit.Ctx(ctx).Where(do.WdkReportAudit{Id: id, AuditUid: userId}).Scan(&wdkReportAudit)
	return
}

// GetWdkReportAuditListById 通过报告ID获取文档库报告审核列表
func (s *sWdkReportAudit) GetWdkReportAuditListById(ctx context.Context, id uint64) (auditList []*entity.WdkReportAudit, err error) {
	err = dao.WdkReportAudit.Ctx(ctx).Where(do.WdkReportAudit{Id: id}).Scan(&auditList)
	return
}

// GetWdkReportAuditCompleteStatus 获取文档库报告审核完成状态
func (s *sWdkReportAudit) GetWdkReportAuditCompleteStatus(auditList []*entity.WdkReportAudit) (status uint) {
	for _, v := range auditList {
		if v.Status == 1 {
			status = 1
			return
		}
	}
	for _, v := range auditList {
		if v.Status == 0 {
			status = 0
			return
		}
	}
	status = 2
	return
}

// GetWdkReportExcellence 获取文档库报告审核完成时的被推荐优秀报告状态
func (s *sWdkReportAudit) GetWdkReportExcellence(auditList []*entity.WdkReportAudit) (excellence uint) {
	for _, v := range auditList {
		if v.Excellence == 0 {
			excellence = 0
			return
		}
	}
	excellence = 1
	return
}

// SetWdkReportAuditStatus 设置文档库报告审核状态
func (s *sWdkReportAudit) SetWdkReportAuditStatus(ctx context.Context, id uint64, userId uint64, status, excellence uint, auditTime *gtime.Time) (err error) {
	_, err = dao.WdkReportAudit.Ctx(ctx).Data(do.WdkReportAudit{
		Status:     status,
		Excellence: excellence,
		AuditTime:  auditTime,
	}).Where(do.WdkReportAudit{
		Id:       id,
		AuditUid: userId,
	}).Update()
	return
}
