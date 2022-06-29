// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportCfgDao is the data access object for table wdk_report_cfg.
type WdkReportCfgDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns WdkReportCfgColumns // columns contains all the column names of Table for convenient usage.
}

// WdkReportCfgColumns defines and stores column names for table wdk_report_cfg.
type WdkReportCfgColumns struct {
	Id        string // 报告类型ID
	Name      string // 报告类型名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

//  wdkReportCfgColumns holds the columns for table wdk_report_cfg.
var wdkReportCfgColumns = WdkReportCfgColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewWdkReportCfgDao creates and returns a new DAO object for table data access.
func NewWdkReportCfgDao() *WdkReportCfgDao {
	return &WdkReportCfgDao{
		group:   "default",
		table:   "wdk_report_cfg",
		columns: wdkReportCfgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkReportCfgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkReportCfgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkReportCfgDao) Columns() WdkReportCfgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkReportCfgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkReportCfgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkReportCfgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
