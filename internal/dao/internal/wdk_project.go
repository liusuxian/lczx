// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WdkProjectDao is the data access object for table wdk_project.
type WdkProjectDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns WdkProjectColumns // columns contains all the column names of Table for convenient usage.
}

// WdkProjectColumns defines and stores column names for table wdk_project.
type WdkProjectColumns struct {
	Id                 string // 文档库项目ID
	Name               string // 项目名称
	Type               string // 项目性质 0:蓝绿体系 1:非绿
	Origin             string // 项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他
	Step               string // 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘
	FileUploadStatus   string // 项目文件上传状态 0:异常 1:正常 2:已完成
	BusinessType       string // 业务类型 0:物业 1:专项 2:全过程
	ContractStatus     string // 签约状态 0:新签 1:续签 2:未签
	ContractSum        string // 合同金额
	DeepCulture        string // 是否为深耕 0:否 1:是
	Status             string // 服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束
	EntrustCompany     string // 委托方公司
	SignCompany        string // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalUid       string // 负责人用户ID
	PrincipalName      string // 负责人姓名
	DeptId             string // 项目所属部门ID
	DeptName           string // 项目所属部门名称
	Region             string // 地区(省/市/区县)
	StartTime          string // 项目开始时间
	EndTime            string // 项目结束时间
	CreateBy           string // 项目创建者用户ID
	CreateName         string // 项目创建者姓名
	UpdatedBy          string // 项目修改者用户ID
	UpdatedName        string // 项目修改者姓名
	Remark             string // 备注
	FileUploadLastTime string // 项目文件最近一次上传时间
	CreatedAt          string // 项目创建时间
	UpdatedAt          string // 项目更新时间
	DeletedAt          string // 项目软删除时间
}

//  wdkProjectColumns holds the columns for table wdk_project.
var wdkProjectColumns = WdkProjectColumns{
	Id:                 "id",
	Name:               "name",
	Type:               "type",
	Origin:             "origin",
	Step:               "step",
	FileUploadStatus:   "file_upload_status",
	BusinessType:       "business_type",
	ContractStatus:     "contract_status",
	ContractSum:        "contract_sum",
	DeepCulture:        "deep_culture",
	Status:             "status",
	EntrustCompany:     "entrust_company",
	SignCompany:        "sign_company",
	PrincipalUid:       "principal_uid",
	PrincipalName:      "principal_name",
	DeptId:             "dept_id",
	DeptName:           "dept_name",
	Region:             "region",
	StartTime:          "start_time",
	EndTime:            "end_time",
	CreateBy:           "create_by",
	CreateName:         "create_name",
	UpdatedBy:          "updated_by",
	UpdatedName:        "updated_name",
	Remark:             "remark",
	FileUploadLastTime: "file_upload_last_time",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// NewWdkProjectDao creates and returns a new DAO object for table data access.
func NewWdkProjectDao() *WdkProjectDao {
	return &WdkProjectDao{
		group:   "default",
		table:   "wdk_project",
		columns: wdkProjectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WdkProjectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WdkProjectDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WdkProjectDao) Columns() WdkProjectColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WdkProjectDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WdkProjectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WdkProjectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
