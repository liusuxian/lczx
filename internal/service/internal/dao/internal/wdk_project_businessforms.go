// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkProjectBusinessformsDao is the data access object for table wdk_project_businessforms.
type WdkProjectBusinessformsDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns WdkProjectBusinessformsColumns // columns contains all the column names of Table for convenient usage.
}

// WdkProjectBusinessformsColumns defines and stores column names for table wdk_project_businessforms.
type WdkProjectBusinessformsColumns struct {
	ProjectId     string // 文档库项目ID
	BusinessForms string // 业态 0:住宅-综合性住宅 1:住宅-高层住宅 2:住宅-多层住宅 3:住宅-联排别墅 4:住宅-独立式住宅 5:非住宅-办公用房 6:非住宅-场馆（体育馆、博物馆、机场） 7:非住宅-综合类 8:非住宅-学校 9:非住宅-工业园 10:非住宅-医院 11:非住宅-商场 12:非住宅-商铺 13:非住宅-车位 14:非住宅-酒店式公寓 15:非住宅-仓储 16:非住宅-经营用房 17:非住宅-物业用房 18:非住宅-配套用房 19:非住宅-车库 20:非住宅-会所 21:非住宅-临停车位 22:非住宅-非机动车位 23:非住宅-社会福利院 24:非住宅-旅游景区
}

//  wdkProjectBusinessformsColumns holds the columns for table wdk_project_businessforms.
var wdkProjectBusinessformsColumns = WdkProjectBusinessformsColumns{
	ProjectId:     "project_id",
	BusinessForms: "business_forms",
}

// NewWdkProjectBusinessformsDao creates and returns a new DAO object for table data access.
func NewWdkProjectBusinessformsDao() *WdkProjectBusinessformsDao {
	return &WdkProjectBusinessformsDao{
		group:   "default",
		table:   "wdk_project_businessforms",
		columns: wdkProjectBusinessformsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkProjectBusinessformsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkProjectBusinessformsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkProjectBusinessformsDao) Columns() WdkProjectBusinessformsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkProjectBusinessformsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkProjectBusinessformsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkProjectBusinessformsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}