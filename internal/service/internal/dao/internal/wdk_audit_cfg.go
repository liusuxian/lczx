// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkAuditCfgDao is the data access object for table wdk_audit_cfg.
type WdkAuditCfgDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns WdkAuditCfgColumns // columns contains all the column names of Table for convenient usage.
}

// WdkAuditCfgColumns defines and stores column names for table wdk_audit_cfg.
type WdkAuditCfgColumns struct {
	TypeId    string // 上传文件类型ID
	AuditUid  string // 审核员用户ID
	AuditName string // 审核员姓名
}

//  wdkAuditCfgColumns holds the columns for table wdk_audit_cfg.
var wdkAuditCfgColumns = WdkAuditCfgColumns{
	TypeId:    "type_id",
	AuditUid:  "audit_uid",
	AuditName: "audit_name",
}

// NewWdkAuditCfgDao creates and returns a new DAO object for table data access.
func NewWdkAuditCfgDao() *WdkAuditCfgDao {
	return &WdkAuditCfgDao{
		group:   "default",
		table:   "wdk_audit_cfg",
		columns: wdkAuditCfgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkAuditCfgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkAuditCfgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkAuditCfgDao) Columns() WdkAuditCfgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkAuditCfgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkAuditCfgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkAuditCfgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}