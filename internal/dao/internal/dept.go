// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeptDao is the data access object for table dept.
type DeptDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns DeptColumns // columns contains all the column names of Table for convenient usage.
}

// DeptColumns defines and stores column names for table dept.
type DeptColumns struct {
	Id            string // 部门ID
	ParentId      string // 父部门ID
	Name          string // 部门名称
	Status        string // 部门状态 0:停用 1:正常
	CreatedBy     string // 创建人
	UpdatedBy     string // 修改人
	PrincipalUid  string // 负责人用户ID
	PrincipalName string // 负责人姓名
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 软删除时间
}

//  deptColumns holds the columns for table dept.
var deptColumns = DeptColumns{
	Id:            "id",
	ParentId:      "parent_id",
	Name:          "name",
	Status:        "status",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	PrincipalUid:  "principal_uid",
	PrincipalName: "principal_name",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewDeptDao creates and returns a new DAO object for table data access.
func NewDeptDao() *DeptDao {
	return &DeptDao{
		group:   "default",
		table:   "dept",
		columns: deptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeptDao) Columns() DeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
