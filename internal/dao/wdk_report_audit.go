// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/dao/internal"
)

// internalWdkReportAuditDao is internal type for wrapping internal DAO implements.
type internalWdkReportAuditDao = *internal.WdkReportAuditDao

// wdkReportAuditDao is the data access object for table wdk_report_audit.
// You can define custom methods on it to extend its functionality as you wish.
type wdkReportAuditDao struct {
	internalWdkReportAuditDao
}

var (
	// WdkReportAudit is globally public accessible object for table wdk_report_audit operations.
	WdkReportAudit = wdkReportAuditDao{
		internal.NewWdkReportAuditDao(),
	}
)

// Fill with you ideas below.
