// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type IWdkReportAudit interface {
	GetWdkReportAuditList(ctx context.Context, req *v1.WdkReportAuditListReq) (total int, list []*v1.WdkReportAuditInfo, err error)
	HandleWdkReportAudit(ctx context.Context, req *v1.WdkReportAuditReq) (err error)
	GetWdkReportUploadAuditList(ctx context.Context, req *v1.WdkReportUploadAuditListReq) (total int, list []*v1.WdkReportUploadAuditInfo, err error)
	HandleWdkReportRescindAudit(ctx context.Context, Id uint64) (err error)
	GetWdkReportAuditProcess(ctx context.Context, Id uint64) (list []*v1.WdkReportAuditProcessInfo, err error)
	GetWdkReportAuditByIdAndUid(ctx context.Context, id uint64, userId uint64) (wdkReportAudit *entity.WdkReportAudit, err error)
	GetWdkReportAuditListById(ctx context.Context, id uint64) (auditList []*entity.WdkReportAudit, err error)
	GetWdkReportAuditCompleteStatus(auditList []*entity.WdkReportAudit) uint
	GetWdkReportExcellence(auditList []*entity.WdkReportAudit) uint
	SetWdkReportAuditStatus(ctx context.Context, wdkReportAudit *entity.WdkReportAudit, req *v1.WdkReportAuditReq, auditTime *gtime.Time) (err error)
}

var localWdkReportAudit IWdkReportAudit

func WdkReportAudit() IWdkReportAudit {
	if localWdkReportAudit == nil {
		panic("implement not found for interface IWdkReportAudit, forgot register?")
	}
	return localWdkReportAudit
}

func RegisterWdkReportAudit(i IWdkReportAudit) {
	localWdkReportAudit = i
}
