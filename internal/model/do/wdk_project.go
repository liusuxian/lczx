// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkProject is the golang structure of table wdk_project for DAO operations like Where/Data.
type WdkProject struct {
	g.Meta           `orm:"table:wdk_project, do:true"`
	Id               interface{} // 文档库项目ID
	Name             interface{} // 项目名称
	Type             interface{} // 项目性质 0:蓝绿体系 1:非绿
	Origin           interface{} // 项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他
	Step             interface{} // 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘
	FileUploadStatus interface{} // 项目文件上传状态 0:异常 1:正常 2:已完成
	BusinessType     interface{} // 业务类型 0:物业 1:专项 2:全过程
	ContractStatus   interface{} // 签约状态 0:新签 1:续签 2:未签
	ContractSum      interface{} // 合同金额
	DeepCulture      interface{} // 是否为深耕 0:否 1:是
	Status           interface{} // 服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中
	EntrustCompany   interface{} // 委托方公司
	SignCompany      interface{} // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalUid     interface{} // 负责人用户ID
	PrincipalName    interface{} // 负责人姓名
	DeptId           interface{} // 项目所属部门ID
	DeptName         interface{} // 项目所属部门名称
	Region           interface{} // 地区(省/市/区县)
	StartTime        *gtime.Time // 项目开始时间
	EndTime          *gtime.Time // 项目结束时间
	CreateBy         interface{} // 项目创建者用户ID
	CreateName       interface{} // 项目创建者姓名
	UpdatedBy        interface{} // 项目修改者用户ID
	UpdatedName      interface{} // 项目修改者姓名
	Remark           interface{} // 备注
	CreateAt         *gtime.Time // 项目创建时间
	UpdateAt         *gtime.Time // 项目更新时间
	DeletedAt        *gtime.Time // 项目软删除时间
}