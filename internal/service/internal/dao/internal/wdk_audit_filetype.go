// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkAuditFiletypeDao is the data access object for table wdk_audit_filetype.
type WdkAuditFiletypeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns WdkAuditFiletypeColumns // columns contains all the column names of Table for convenient usage.
}

// WdkAuditFiletypeColumns defines and stores column names for table wdk_audit_filetype.
type WdkAuditFiletypeColumns struct {
	AuditUid string // 审核员用户ID
	FileId   string // 审核的文件ID
	TypeId   string // 审核文件类型ID 详见wdk_filetype_cfg配置
}

//  wdkAuditFiletypeColumns holds the columns for table wdk_audit_filetype.
var wdkAuditFiletypeColumns = WdkAuditFiletypeColumns{
	AuditUid: "audit_uid",
	FileId:   "file_id",
	TypeId:   "type_id",
}

// NewWdkAuditFiletypeDao creates and returns a new DAO object for table data access.
func NewWdkAuditFiletypeDao() *WdkAuditFiletypeDao {
	return &WdkAuditFiletypeDao{
		group:   "default",
		table:   "wdk_audit_filetype",
		columns: wdkAuditFiletypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkAuditFiletypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkAuditFiletypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkAuditFiletypeDao) Columns() WdkAuditFiletypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkAuditFiletypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkAuditFiletypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkAuditFiletypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}