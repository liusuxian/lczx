// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/dao/internal"
)

// internalWdkReportAuditCfgDao is internal type for wrapping internal DAO implements.
type internalWdkReportAuditCfgDao = *internal.WdkReportAuditCfgDao

// wdkReportAuditCfgDao is the data access object for table wdk_report_audit_cfg.
// You can define custom methods on it to extend its functionality as you wish.
type wdkReportAuditCfgDao struct {
	internalWdkReportAuditCfgDao
}

var (
	// WdkReportAuditCfg is globally public accessible object for table wdk_report_audit_cfg operations.
	WdkReportAuditCfg = wdkReportAuditCfgDao{
		internal.NewWdkReportAuditCfgDao(),
	}
)

// Fill with you ideas below.
