// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportAuditDao is the data access object for table wdk_report_audit.
type WdkReportAuditDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns WdkReportAuditColumns // columns contains all the column names of Table for convenient usage.
}

// WdkReportAuditColumns defines and stores column names for table wdk_report_audit.
type WdkReportAuditColumns struct {
	Id             string // 报告ID
	AuditUid       string // 审核员用户ID
	AuditorType    string // 审核员类型 0:负责人 1:专家组
	ProjectId      string // 所属项目ID
	AuditName      string // 审核员姓名
	Rescind        string // 是否已撤销 0:否 1:是
	PreauditStatus string // 前置审核是否已通过 0:否 1:是
	Status         string // 审核状态 0:未通过 1:审核中 2:已通过
	Excellence     string // 是否被推荐为优秀报告 0:未被推荐为优秀报告 1:已被推荐为优秀报告
	AuditTime      string // 审核时间
	AuditOpinion   string // 审核意见
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

//  wdkReportAuditColumns holds the columns for table wdk_report_audit.
var wdkReportAuditColumns = WdkReportAuditColumns{
	Id:             "id",
	AuditUid:       "audit_uid",
	AuditorType:    "auditor_type",
	ProjectId:      "project_id",
	AuditName:      "audit_name",
	Rescind:        "rescind",
	PreauditStatus: "preaudit_status",
	Status:         "status",
	Excellence:     "excellence",
	AuditTime:      "audit_time",
	AuditOpinion:   "audit_opinion",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewWdkReportAuditDao creates and returns a new DAO object for table data access.
func NewWdkReportAuditDao() *WdkReportAuditDao {
	return &WdkReportAuditDao{
		group:   "default",
		table:   "wdk_report_audit",
		columns: wdkReportAuditColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkReportAuditDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkReportAuditDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkReportAuditDao) Columns() WdkReportAuditColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkReportAuditDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkReportAuditDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkReportAuditDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
