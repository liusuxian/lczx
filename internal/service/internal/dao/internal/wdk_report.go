// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportDao is the data access object for table wdk_report.
type WdkReportDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns WdkReportColumns // columns contains all the column names of Table for convenient usage.
}

// WdkReportColumns defines and stores column names for table wdk_report.
type WdkReportColumns struct {
	Id          string // 报告ID
	ProjectId   string // 所属项目ID
	Name        string // 报告名称
	CreateBy    string // 上传者用户ID
	CreateName  string // 上传者姓名
	AuditStatus string // 审核状态 0:未通过 1:审核中 2:已通过 3:后台管理员自动通过
	Excellence  string // 是否是优秀报告 0:未被评选为优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告
	AuditTime   string // 审核时间
	OriginUrl   string // 原始文件url
	PdfUrl      string // pdf文件url
	CreateAt    string // 上传时间
	UpdateAt    string // 更新时间
}

//  wdkReportColumns holds the columns for table wdk_report.
var wdkReportColumns = WdkReportColumns{
	Id:          "id",
	ProjectId:   "project_id",
	Name:        "name",
	CreateBy:    "create_by",
	CreateName:  "create_name",
	AuditStatus: "audit_status",
	Excellence:  "excellence",
	AuditTime:   "audit_time",
	OriginUrl:   "origin_url",
	PdfUrl:      "pdf_url",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewWdkReportDao creates and returns a new DAO object for table data access.
func NewWdkReportDao() *WdkReportDao {
	return &WdkReportDao{
		group:   "default",
		table:   "wdk_report",
		columns: wdkReportColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkReportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkReportDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkReportDao) Columns() WdkReportColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkReportDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkReportDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkReportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
