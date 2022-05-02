package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/model"
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
	err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).
		WhereIn(dao.WdkReport.Columns().AuditStatus, []uint{2, 3}).ScanList(&list, "Report")
	if err != nil {
		return
	}
	err = dao.WdkReportType.Ctx(ctx).Where(dao.WdkReportType.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportType", "Report", "Id:Id")
	return
}

// AddWdkReport 新增文档库上传报告记录
func (s *sWdkReport) AddWdkReport(ctx context.Context, req *v1.WdkReportAddReq, report *upload.FileInfo) (err error) {
	err = dao.WdkReport.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查新增文档库上传报告记录权限
		var terr error
		var wdkProject *entity.WdkProject
		wdkProject, terr = s.AuthAdd(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}
		// 通过报告类型ID列表获取报告类型配置
		typeIdSet := gset.NewFrom(req.TypeIds)
		typeIds := typeIdSet.Slice()
		var reportCfgInfos []*v1.WdkReportCfgInfo
		reportCfgInfos, terr = WdkReportCfg().GetWdkReportCfgByIds(ctx, gconv.Uint64s(typeIds))
		if terr != nil {
			return terr
		}
		if len(reportCfgInfos) == 0 {
			return gerror.Newf("报告类型ID列表%v不存在", req.TypeIds)
		}
		// 保存文档库上传报告记录
		curUser := Context().Get(ctx).User
		var reportId int64
		reportId, terr = s.saveWdkReport(ctx, curUser, wdkProject, reportCfgInfos, report)
		if terr != nil {
			return terr
		}
		// 保存文档库上传报告审核信息
		terr = s.saveWdkReportAudit(ctx, curUser, reportCfgInfos, gconv.Uint64(reportId), wdkProject.Id)
		if terr != nil {
			return terr
		}
		if curUser.IsAdmin == 1 {
			// 设置所属文档库项目阶段
			terr = WdkProject().SetWdkProjectStep(ctx, wdkProject.Id, 8)
			if terr != nil {
				return terr
			}
			// 设置所属文档库项目文件上传状态为是
			terr = WdkProject().SetWdkProjectFileUploadStatus(ctx, wdkProject.Id)
			if terr != nil {
				return terr
			}
		}
		return nil
	})
	return
}

// GetWdkReportExcellenceList 获取文档库优秀报告列表
func (s *sWdkReport) GetWdkReportExcellenceList(ctx context.Context, req *v1.WdkReportExcellenceListReq) (total int, list []*v1.WdkReportInfo, err error) {
	var reportIds []uint64
	if req.TypeId != "" {
		var array []*gvar.Var
		array, err = dao.WdkReportType.Ctx(ctx).Fields(dao.WdkReportType.Columns().Id).Where(do.WdkReportType{TypeId: req.TypeId}).Array()
		if err != nil {
			return
		}
		reportIds = make([]uint64, 0, len(array))
		for _, v := range array {
			reportIds = append(reportIds, v.Uint64())
		}
	}
	reportModel := dao.WdkReport.Ctx(ctx).Where(do.WdkReport{Excellence: req.Excellence})
	columns := dao.WdkReport.Columns()
	if len(reportIds) != 0 {
		reportModel = reportModel.WhereIn(columns.Id, reportIds)
	}
	if req.ReportName != "" {
		reportModel = reportModel.WhereLike(columns.Name, "%"+req.ReportName+"%")
	}
	if req.ProjectName != "" {
		reportModel = reportModel.WhereLike(columns.ProjectName, "%"+req.ProjectName+"%")
	}
	total, err = reportModel.Count()
	if err != nil {
		return
	}
	err = reportModel.Page(req.CurPage, req.PageSize).OrderDesc(columns.AuditTime).Scan(&list)
	if err != nil {
		return
	}
	err = dao.WdkReportType.Ctx(ctx).Where(dao.WdkReportType.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportType", "Report", "Id:Id")
	return
}

// AuthAdd 检查新增文档库上传报告记录权限
func (s *sWdkReport) AuthAdd(ctx context.Context, projectId uint64) (wdkProject *entity.WdkProject, err error) {
	// 通过文档库项目ID判断文档库项目信息是否存在
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
	count, err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).
		WhereIn(dao.WdkReport.Columns().AuditStatus, []uint{2, 3}).Count()
	return
}

// SetWdkReportExcellence 设置文档库报告是否被评选为优秀报告
func (s *sWdkReport) SetWdkReportExcellence(ctx context.Context, id uint64, excellence uint) (err error) {
	_, err = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{Excellence: excellence}).Where(do.WdkReport{Id: id}).Update()
	return
}

// SetWdkReportAuditCompleteStatus 设置文档库报告审核完成状态
func (s *sWdkReport) SetWdkReportAuditCompleteStatus(ctx context.Context, id uint64, status, excellence uint, auditTime *gtime.Time) (err error) {
	_, err = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
		AuditStatus: status,
		Excellence:  excellence,
		AuditTime:   auditTime,
	}).Where(do.WdkReport{Id: id}).Update()
	return
}

// saveWdkReport 保存文档库上传报告记录
func (s *sWdkReport) saveWdkReport(ctx context.Context, user *model.ContextUser, wdkProject *entity.WdkProject, reportCfgInfos []*v1.WdkReportCfgInfo, report *upload.FileInfo) (reportId int64, err error) {
	if user.IsAdmin == 1 {
		// 管理员不需要走审核流程
		reportId, err = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
			ProjectId:   wdkProject.Id,
			ProjectName: wdkProject.Name,
			Name:        report.FileName,
			CreateBy:    user.Id,
			CreateName:  user.Realname,
			AuditStatus: 3,
			Excellence:  0,
			AuditTime:   gtime.Now(),
			OriginUrl:   report.OriginFileUrl,
			PdfUrl:      report.PdfFileUrl,
		}).FieldsEx(dao.WdkReport.Columns().Id).InsertAndGetId()
	} else {
		// 非管理员需要走审核流程
		reportId, err = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
			ProjectId:   wdkProject.Id,
			ProjectName: wdkProject.Name,
			Name:        report.FileName,
			CreateBy:    user.Id,
			CreateName:  user.Realname,
			AuditStatus: 1,
			Excellence:  0,
			OriginUrl:   report.OriginFileUrl,
			PdfUrl:      report.PdfFileUrl,
		}).FieldsEx(dao.WdkReport.Columns().Id).InsertAndGetId()
	}
	if err != nil {
		return
	}
	// 保存文档库上传报告类型数据
	reportTypeData := g.List{}
	for _, v := range reportCfgInfos {
		reportTypeData = append(reportTypeData, g.Map{
			"id":        reportId,
			"type_id":   v.ReportCfg.Id,
			"type_name": v.ReportCfg.Name,
		})
	}
	_, err = dao.WdkReportType.Ctx(ctx).Data(reportTypeData).Batch(len(reportTypeData)).Insert()
	return
}

// saveWdkReportAuditRecord 保存文档库上传报告审核信息
func (s *sWdkReport) saveWdkReportAudit(ctx context.Context, user *model.ContextUser, reportCfgInfos []*v1.WdkReportCfgInfo, reportId, projectId uint64) (err error) {
	if user.IsAdmin != 1 {
		reportAudits := make([]entity.WdkReportAudit, 0, len(reportCfgInfos))
		reportAuditTypes := make([]entity.WdkReportAuditType, 0, len(reportCfgInfos))
		for _, reportCfgInfo := range reportCfgInfos {
			for _, reportAuditCfgInfo := range reportCfgInfo.ReportAuditCfg {
				reportAudits = append(reportAudits, entity.WdkReportAudit{
					Id:         reportId,
					AuditUid:   reportAuditCfgInfo.AuditUid,
					ProjectId:  projectId,
					AuditName:  reportAuditCfgInfo.AuditName,
					Status:     1,
					Excellence: 0,
				})
				reportAuditTypes = append(reportAuditTypes, entity.WdkReportAuditType{
					Id:       reportId,
					AuditUid: reportAuditCfgInfo.AuditUid,
					TypeId:   reportAuditCfgInfo.Id,
					TypeName: reportAuditCfgInfo.TypeName,
				})
			}
		}
		reportAuditData := gconv.Maps(gset.NewFrom(reportAudits).Slice())
		reportAuditTypeData := gconv.Maps(gset.NewFrom(reportAuditTypes).Slice())
		_, err = dao.WdkReportAudit.Ctx(ctx).Data(reportAuditData).Batch(len(reportAuditData)).Insert()
		if err != nil {
			return
		}
		_, err = dao.WdkReportAuditType.Ctx(ctx).Data(reportAuditTypeData).Batch(len(reportAuditTypeData)).Insert()
		if err != nil {
			return
		}
	}
	return
}
