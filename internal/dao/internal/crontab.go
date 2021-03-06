// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CrontabDao is the data access object for table crontab.
type CrontabDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CrontabColumns // columns contains all the column names of Table for convenient usage.
}

// CrontabColumns defines and stores column names for table crontab.
type CrontabColumns struct {
	Id             string // 任务ID
	Name           string // 任务名称
	Group          string // 任务组名
	Params         string // 参数
	InvokeTarget   string // 调用目标字符串
	CronExpression string // cron执行表达式
	MisfirePolicy  string // 计划执行策略 0:执行一次 1:执行多次
	Status         string // 状态 0:暂停 1:正常
	CreateBy       string // 创建者用户ID
	UpdateBy       string // 更新者用户ID
	Remark         string // 备注
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

//  crontabColumns holds the columns for table crontab.
var crontabColumns = CrontabColumns{
	Id:             "id",
	Name:           "name",
	Group:          "group",
	Params:         "params",
	InvokeTarget:   "invoke_target",
	CronExpression: "cron_expression",
	MisfirePolicy:  "misfire_policy",
	Status:         "status",
	CreateBy:       "create_by",
	UpdateBy:       "update_by",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewCrontabDao creates and returns a new DAO object for table data access.
func NewCrontabDao() *CrontabDao {
	return &CrontabDao{
		group:   "default",
		table:   "crontab",
		columns: crontabColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CrontabDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CrontabDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CrontabDao) Columns() CrontabColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CrontabDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CrontabDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CrontabDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
