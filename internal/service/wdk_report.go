package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
	err = dao.WdkReport.Ctx(ctx).Where(do.WdkReport{ProjectId: projectId}).
		WhereIn(dao.WdkReport.Columns().AuditStatus, []uint{2, 3}).ScanList(&list, "Report")
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
		// 通过报告类型ID列表获取报告类型配置
		typeIdSet := gset.NewFrom(req.TypeIds)
		typeIds := typeIdSet.Slice()
		var reportCfgInfos []*v1.WdkReportCfgInfo
		reportCfgInfos, terr = WdkReportCfg().GetWdkReportCfgByIds(ctx, gconv.Uint64s(typeIds))
		if terr != nil {
			return terr
		}
		if len(typeIds) != len(reportCfgInfos) {
			return gerror.Newf("报告类型ID列表[%s]不存在", gstr.JoinAny(req.TypeIds, ","))
		}
		// 获取文档库上传报告的审核人员们的姓名和数量
		auditNames, auditCount := s.GetWdkReportAuditNamesAndCount(reportCfgInfos)
		// 写入文档库上传报告记录数据
		var reportId int64
		user := Context().Get(ctx).User
		//user := &model.ContextUser{
		//	Id:       3,
		//	Realname: "普通用户",
		//	IsAdmin:  0,
		//}
		if user.IsAdmin == 1 {
			// 管理员不需要走审核流程
			reportId, terr = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
				ProjectId:    req.ProjectId,
				Name:         report.FileName,
				CreateBy:     user.Id,
				CreateName:   user.Realname,
				AuditStatus:  3,
				AuditNames:   auditNames,
				AuditCount:   auditCount,
				Excellence:   0,
				AuditEndTime: gtime.Now(),
				OriginUrl:    report.OriginFileUrl,
				PdfUrl:       report.PdfFileUrl,
			}).InsertAndGetId()
		} else {
			// 非管理员需要走审核流程
			reportId, terr = dao.WdkReport.Ctx(ctx).Data(do.WdkReport{
				ProjectId:   req.ProjectId,
				Name:        report.FileName,
				CreateBy:    user.Id,
				CreateName:  user.Realname,
				AuditStatus: 1,
				AuditNames:  auditNames,
				AuditCount:  auditCount,
				Excellence:  0,
				OriginUrl:   report.OriginFileUrl,
				PdfUrl:      report.PdfFileUrl,
			}).InsertAndGetId()
		}
		if terr != nil {
			return terr
		}
		// 写入文档库上传报告类型数据
		reportTypeData := g.List{}
		for _, v := range reportCfgInfos {
			reportTypeData = append(reportTypeData, g.Map{
				"id":        reportId,
				"type_id":   v.ReportCfg.Id,
				"type_name": v.ReportCfg.Name,
			})
		}
		_, terr = dao.WdkReportType.Ctx(ctx).Data(reportTypeData).Batch(len(reportTypeData)).Insert()
		if terr != nil {
			return terr
		}
		// 写入文档库上传报告审核数据
		if user.IsAdmin != 1 {
			terr = s.AddWdkReportAuditRecord(ctx, reportCfgInfos, gconv.Uint64(reportId))
			if terr != nil {
				return terr
			}
		}
		// 设置所属文档库项目阶段
		terr = WdkProject().SetWdkProjectStep(ctx, req.ProjectId, 8)
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目文件上传状态为是
		terr = WdkProject().SetWdkProjectFileUploadStatus(ctx, req.ProjectId)
		return terr
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

// GetWdkReportAuditNamesAndCount 获取文档库上传报告的审核人员们的姓名和数量
func (s *sWdkReport) GetWdkReportAuditNamesAndCount(reportCfgInfos []*v1.WdkReportCfgInfo) (auditNames string, auditCount int) {
	auditNameList := make([]string, 0, len(reportCfgInfos))
	for _, reportCfgInfo := range reportCfgInfos {
		for _, reportAuditCfgInfo := range reportCfgInfo.ReportAuditCfg {
			auditNameList = append(auditNameList, reportAuditCfgInfo.AuditName)
		}
	}
	auditNameSet := gset.NewFrom(auditNameList)
	auditNames = gstr.JoinAny(auditNameSet.Slice(), ",")
	auditCount = auditNameSet.Size()
	return
}

// AddWdkReportAuditRecord 新增文档库上传报告审核记录
func (s *sWdkReport) AddWdkReportAuditRecord(ctx context.Context, reportCfgInfos []*v1.WdkReportCfgInfo, reportId uint64) (err error) {
	reportAuditRecords := make([]entity.WdkReportAuditRecord, 0, len(reportCfgInfos))
	reportAuditTypes := make([]entity.WdkReportAuditType, 0, len(reportCfgInfos))
	for _, reportCfgInfo := range reportCfgInfos {
		for _, reportAuditCfgInfo := range reportCfgInfo.ReportAuditCfg {
			reportAuditRecords = append(reportAuditRecords, entity.WdkReportAuditRecord{
				AuditUid:  reportAuditCfgInfo.AuditUid,
				ReportId:  reportId,
				Status:    1,
				AuditName: reportAuditCfgInfo.AuditName,
			})
			reportAuditTypes = append(reportAuditTypes, entity.WdkReportAuditType{
				AuditUid: reportAuditCfgInfo.AuditUid,
				ReportId: reportId,
				TypeId:   reportAuditCfgInfo.Id,
				TypeName: reportAuditCfgInfo.TypeName,
			})
		}
	}
	reportAuditRecordData := gconv.Maps(gset.NewFrom(reportAuditRecords).Slice())
	reportAuditTypeData := gconv.Maps(gset.NewFrom(reportAuditTypes).Slice())
	_, err = dao.WdkReportAuditRecord.Ctx(ctx).Data(reportAuditRecordData).Batch(len(reportAuditRecordData)).Insert()
	if err != nil {
		return err
	}
	_, err = dao.WdkReportAuditType.Ctx(ctx).Data(reportAuditTypeData).Batch(len(reportAuditTypeData)).Insert()
	if err != nil {
		return err
	}
	return
}
