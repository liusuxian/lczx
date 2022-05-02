// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportTypeDao is the data access object for table wdk_report_type.
type WdkReportTypeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns WdkReportTypeColumns // columns contains all the column names of Table for convenient usage.
}

// WdkReportTypeColumns defines and stores column names for table wdk_report_type.
type WdkReportTypeColumns struct {
	Id          string // 报告ID
	TypeId      string // 报告类型ID
	TypeName    string // 报告类型名称
	ReportName  string // 报告名称
	ProjectId   string // 所属项目ID
	ProjectName string // 所属项目名称
}

//  wdkReportTypeColumns holds the columns for table wdk_report_type.
var wdkReportTypeColumns = WdkReportTypeColumns{
	Id:          "id",
	TypeId:      "type_id",
	TypeName:    "type_name",
	ReportName:  "report_name",
	ProjectId:   "project_id",
	ProjectName: "project_name",
}

// NewWdkReportTypeDao creates and returns a new DAO object for table data access.
func NewWdkReportTypeDao() *WdkReportTypeDao {
	return &WdkReportTypeDao{
		group:   "default",
		table:   "wdk_report_type",
		columns: wdkReportTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkReportTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkReportTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkReportTypeDao) Columns() WdkReportTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkReportTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkReportTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkReportTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
