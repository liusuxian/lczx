// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/service/internal/dao/internal"
)

// internalWdkReportAuditRecordDao is internal type for wrapping internal DAO implements.
type internalWdkReportAuditRecordDao = *internal.WdkReportAuditRecordDao

// wdkReportAuditRecordDao is the data access object for table wdk_report_audit_record.
// You can define custom methods on it to extend its functionality as you wish.
type wdkReportAuditRecordDao struct {
	internalWdkReportAuditRecordDao
}

var (
	// WdkReportAuditRecord is globally public accessible object for table wdk_report_audit_record operations.
	WdkReportAuditRecord = wdkReportAuditRecordDao{
		internal.NewWdkReportAuditRecordDao(),
	}
)

// Fill with you ideas below.
