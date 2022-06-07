// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportAuditCfgDao is the data access object for table wdk_report_audit_cfg.
type WdkReportAuditCfgDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns WdkReportAuditCfgColumns // columns contains all the column names of Table for convenient usage.
}

// WdkReportAuditCfgColumns defines and stores column names for table wdk_report_audit_cfg.
type WdkReportAuditCfgColumns struct {
	Id        string // 报告类型ID
	AuditUid  string // 审核员用户ID
	TypeName  string // 报告类型名称
	AuditName string // 审核员姓名
}

//  wdkReportAuditCfgColumns holds the columns for table wdk_report_audit_cfg.
var wdkReportAuditCfgColumns = WdkReportAuditCfgColumns{
	Id:        "id",
	AuditUid:  "audit_uid",
	TypeName:  "type_name",
	AuditName: "audit_name",
}

// NewWdkReportAuditCfgDao creates and returns a new DAO object for table data access.
func NewWdkReportAuditCfgDao() *WdkReportAuditCfgDao {
	return &WdkReportAuditCfgDao{
		group:   "default",
		table:   "wdk_report_audit_cfg",
		columns: wdkReportAuditCfgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkReportAuditCfgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkReportAuditCfgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkReportAuditCfgDao) Columns() WdkReportAuditCfgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkReportAuditCfgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkReportAuditCfgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkReportAuditCfgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
