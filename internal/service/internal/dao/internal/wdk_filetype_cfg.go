// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkFiletypeCfgDao is the data access object for table wdk_filetype_cfg.
type WdkFiletypeCfgDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns WdkFiletypeCfgColumns // columns contains all the column names of Table for convenient usage.
}

// WdkFiletypeCfgColumns defines and stores column names for table wdk_filetype_cfg.
type WdkFiletypeCfgColumns struct {
	Id       string // 上传文件类型ID
	Name     string // 类型名称
	Multiple string // 是否同时存在多个文件 0:否 1:是
	Audit    string // 是否需要审核 0:不需要 1:需要
	Step     string // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

//  wdkFiletypeCfgColumns holds the columns for table wdk_filetype_cfg.
var wdkFiletypeCfgColumns = WdkFiletypeCfgColumns{
	Id:       "id",
	Name:     "name",
	Multiple: "multiple",
	Audit:    "audit",
	Step:     "step",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewWdkFiletypeCfgDao creates and returns a new DAO object for table data access.
func NewWdkFiletypeCfgDao() *WdkFiletypeCfgDao {
	return &WdkFiletypeCfgDao{
		group:   "default",
		table:   "wdk_filetype_cfg",
		columns: wdkFiletypeCfgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkFiletypeCfgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkFiletypeCfgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkFiletypeCfgDao) Columns() WdkFiletypeCfgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkFiletypeCfgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkFiletypeCfgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkFiletypeCfgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
