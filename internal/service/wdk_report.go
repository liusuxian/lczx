package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
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
		// 检查部门负责人是否存在
		var terr error
		user := Context().Get(ctx).User
		var dept *entity.Dept
		terr = dao.Dept.Ctx(ctx).Where(do.Dept{Id: user.DeptId}).Scan(&dept)
		if terr != nil {
			return terr
		}
		if dept == nil {
			return gerror.Newf(`部门ID[%d]不存在`, user.DeptId)
		}
		if dept.PrincipalUid == 0 {
			return gerror.Newf(`部门[%s]负责人不存在，请添加部门负责人`, dept.Name)
		}
		// 处理客户端传来的参数
		auditTypeInfoMap := gmap.New()
		for _, auditTypeInfo := range req.AuditTypeInfos {
			auditUidSet := gset.NewFrom(auditTypeInfo.AuditUids)
			if auditUidSet.Size() != len(auditTypeInfo.AuditUids) {
				return gerror.Newf("报告审核类型信息列表%v不存在，请刷新页面后重试", req.AuditTypeInfos)
			}
			auditTypeInfoMap.Set(auditTypeInfo.TypeId, auditUidSet)
		}
		if len(req.AuditTypeInfos) != auditTypeInfoMap.Size() {
			return gerror.Newf("报告审核类型信息列表%v不存在，请刷新页面后重试", req.AuditTypeInfos)
		}
		// 通过报告类型ID列表获取报告类型配置
		var reportCfg []*v1.WdkReportCfgInfo
		reportCfg, terr = WdkReportCfg().GetWdkReportCfgByIds(ctx, gconv.Uint64s(auditTypeInfoMap.Keys()))
		if terr != nil {
			return terr
		}
		if len(reportCfg) == 0 || len(reportCfg) != auditTypeInfoMap.Size() {
			return gerror.Newf("报告审核类型信息列表%v不存在，请刷新页面后重试", req.AuditTypeInfos)
		}
		// 处理报告类型
		reportTypeInfo := make([]*v1.WdkReportCfgInfo, 0, len(reportCfg))
		for _, r := range reportCfg {
			auditUidSet := auditTypeInfoMap.Get(r.ReportCfg.Id).(*gset.Set)
			reportAuditCfg := make([]*entity.WdkReportAuditCfg, 0, len(r.ReportAuditCfg))
			for _, ra := range r.ReportAuditCfg {
				if auditUidSet.Contains(ra.AuditUid) {
					reportAuditCfg = append(reportAuditCfg, ra)
				}
			}
			if len(reportAuditCfg) > 0 && len(reportAuditCfg) == auditUidSet.Size() {
				reportTypeInfo = append(reportTypeInfo, &v1.WdkReportCfgInfo{
					ReportCfg:      r.ReportCfg,
					ReportAuditCfg: reportAuditCfg,
				})
			}
		}
		if len(reportTypeInfo) == 0 || len(reportTypeInfo) != auditTypeInfoMap.Size() {
			return gerror.Newf("报告审核类型信息列表%v不存在", req.AuditTypeInfos)
		}
		// 保存文档库上传报告记录
		var reportId int64
		reportId, terr = s.saveWdkReport(ctx, user, req, reportTypeInfo, report)
		if terr != nil {
			return terr
		}
		// 保存文档库上传报告审核信息
		terr = s.saveWdkReportAudit(ctx, user, reportTypeInfo, dept, gconv.Uint64(reportId), req.ProjectId)
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目阶段
		terr = WdkProject().SetWdkProjectStepByReportStep(ctx, req.ProjectId, req.Step)
		if terr != nil {
			return terr
		}
		if user.IsAdmin == 1 {
			// 设置文档库项目文件上传状态为已完成
			terr = WdkProject().SetWdkProjectFileUploadStatusFinish(ctx, req.ProjectId)
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
	reportModel := dao.WdkReport.Ctx(ctx).Where(do.WdkReport{Excellence: req.Excellence, AuditStatus: 2})
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
	err = reportModel.Page(req.CurPage, req.PageSize).OrderDesc(columns.AuditTime).ScanList(&list, "Report")
	if err != nil {
		return
	}
	err = dao.WdkReportType.Ctx(ctx).Where(dao.WdkReportType.Columns().Id, gdb.ListItemValuesUnique(list, "Report", "Id")).
		ScanList(&list, "ReportType", "Report", "Id:Id")
	return
}

// GetWdkReportCountByProjectId 通过项目ID获取文档库项目上传报告类型数量
func (s *sWdkReport) GetWdkReportCountByProjectId(ctx context.Context, projectId uint64) (count int, err error) {
	// 查询审核已通过的报告信息
	var list []*entity.WdkReport
	err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).
		WhereIn(dao.WdkReport.Columns().AuditStatus, []uint{2, 3}).Scan(&list)
	if err != nil {
		return
	}
	if len(list) > 0 {
		// 获取审核已通过的报告ID
		ids := make([]uint64, 0, len(list))
		for _, v := range list {
			ids = append(ids, v.Id)
		}
		count, err = dao.WdkReportType.Ctx(ctx).Fields(dao.WdkReportType.Columns().TypeId).
			WhereIn(dao.WdkReportType.Columns().Id, ids).Distinct().Count()
	}
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
func (s *sWdkReport) saveWdkReport(ctx context.Context, user *model.ContextUser, req *v1.WdkReportAddReq, reportTypeInfo []*v1.WdkReportCfgInfo, report *upload.FileInfo) (reportId int64, err error) {
	if user.IsAdmin == 1 {
		// 管理员不需要走审核流程
		reportId, err = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
			ProjectId:   req.ProjectId,
			ProjectName: req.ProjectName,
			Name:        report.FileName,
			CreateBy:    user.Id,
			CreateName:  user.Realname,
			Rescind:     0,
			AuditStatus: 3,
			Excellence:  0,
			AuditTime:   gtime.Now(),
			OriginUrl:   report.OriginFileUrl,
			PdfUrl:      report.PdfFileUrl,
		}).FieldsEx(dao.WdkReport.Columns().Id).InsertAndGetId()
	} else {
		// 非管理员需要走审核流程
		reportId, err = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
			ProjectId:   req.ProjectId,
			ProjectName: req.ProjectName,
			Name:        report.FileName,
			CreateBy:    user.Id,
			CreateName:  user.Realname,
			Rescind:     0,
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
	for _, v := range reportTypeInfo {
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
func (s *sWdkReport) saveWdkReportAudit(ctx context.Context, user *model.ContextUser, reportTypeInfo []*v1.WdkReportCfgInfo, dept *entity.Dept, reportId, projectId uint64) (err error) {
	if user.IsAdmin != 1 {
		reportAudits := make([]entity.WdkReportAudit, 0, len(reportTypeInfo))
		reportAuditTypes := make([]entity.WdkReportAuditType, 0, len(reportTypeInfo))
		// 添加负责人审核
		reportAudits = append(reportAudits, entity.WdkReportAudit{
			Id:             reportId,
			AuditUid:       dept.PrincipalUid,
			AuditorType:    0,
			ProjectId:      projectId,
			AuditName:      dept.PrincipalName,
			Rescind:        0,
			PreauditStatus: 1,
			Status:         1,
			Excellence:     0,
		})
		for _, rt := range reportTypeInfo {
			// 添加负责人审核
			reportAuditTypes = append(reportAuditTypes, entity.WdkReportAuditType{
				Id:          reportId,
				AuditUid:    dept.PrincipalUid,
				AuditorType: 0,
				TypeId:      rt.ReportCfg.Id,
				TypeName:    rt.ReportCfg.Name,
			})
			for _, ra := range rt.ReportAuditCfg {
				reportAudits = append(reportAudits, entity.WdkReportAudit{
					Id:             reportId,
					AuditUid:       ra.AuditUid,
					AuditorType:    1,
					ProjectId:      projectId,
					AuditName:      ra.AuditName,
					Rescind:        0,
					PreauditStatus: 0,
					Status:         1,
					Excellence:     0,
				})
				reportAuditTypes = append(reportAuditTypes, entity.WdkReportAuditType{
					Id:          reportId,
					AuditUid:    ra.AuditUid,
					AuditorType: 1,
					TypeId:      ra.Id,
					TypeName:    ra.TypeName,
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
