// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/internal/model"

	"github.com/gogf/gf/v2/os/gtime"
)

type IWdkReport interface {
	GetWdkReportRecord(ctx context.Context, projectId uint64) (list []*v1.WdkReportInfo, err error)
	AddWdkReport(ctx context.Context, req *v1.WdkReportAddReq, report *model.UploadFileInfo) (err error)
	GetWdkReportExcellenceList(ctx context.Context, req *v1.WdkReportExcellenceListReq) (total int, list []*v1.WdkReportInfo, err error)
	GetWdkReportCountByProjectId(ctx context.Context, projectId uint64) (count int, err error)
	SetWdkReportExcellence(ctx context.Context, id uint64, excellence uint) (err error)
	SetWdkReportAuditCompleteStatus(ctx context.Context, id uint64, status, excellence uint, auditTime *gtime.Time) (err error)
}

var localWdkReport IWdkReport

func WdkReport() IWdkReport {
	if localWdkReport == nil {
		panic("implement not found for interface IWdkReport, forgot register?")
	}
	return localWdkReport
}

func RegisterWdkReport(i IWdkReport) {
	localWdkReport = i
}
