// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRuleDao is the data access object for table auth_rule.
type AuthRuleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AuthRuleColumns // columns contains all the column names of Table for convenient usage.
}

// AuthRuleColumns defines and stores column names for table auth_rule.
type AuthRuleColumns struct {
	Id         string // 规则ID
	ParentId   string // 父规则ID
	ApiPath    string // 接口路径
	ApiName    string // 接口名称
	Condition  string // 条件
	MenuType   string // 类型 0:目录 1:菜单 2:按钮
	Status     string // 菜单状态 0:停用 1:正常
	Show       string // 显示状态 0:隐藏 1:显示
	Path       string // 路由地址
	JumpPath   string // 跳转路由
	Component  string // 组件路径
	IsFrame    string // 是否外链 1是 0否
	ModuleType string // 所属模块
	ModelId    string // 模型ID
	Remark     string // 备注
	CreateAt   string // 创建时间
	UpdateAt   string // 更新时间
	DeletedAt  string // 软删除时间
}

//  authRuleColumns holds the columns for table auth_rule.
var authRuleColumns = AuthRuleColumns{
	Id:         "id",
	ParentId:   "parent_id",
	ApiPath:    "api_path",
	ApiName:    "api_name",
	Condition:  "condition",
	MenuType:   "menu_type",
	Status:     "status",
	Show:       "show",
	Path:       "path",
	JumpPath:   "jump_path",
	Component:  "component",
	IsFrame:    "is_frame",
	ModuleType: "module_type",
	ModelId:    "model_id",
	Remark:     "remark",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeletedAt:  "deleted_at",
}

// NewAuthRuleDao creates and returns a new DAO object for table data access.
func NewAuthRuleDao() *AuthRuleDao {
	return &AuthRuleDao{
		group:   "default",
		table:   "auth_rule",
		columns: authRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthRuleDao) Columns() AuthRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
