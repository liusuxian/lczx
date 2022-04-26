// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkReport is the golang structure of table wdk_report for DAO operations like Where/Data.
type WdkReport struct {
	g.Meta       `orm:"table:wdk_report, do:true"`
	Id           interface{} // 报告ID
	ProjectId    interface{} // 所属项目ID
	Name         interface{} // 报告名称
	CreateBy     interface{} // 上传者用户ID
	CreateName   interface{} // 上传者姓名
	AuditStatus  interface{} // 审核状态 0:未通过 1:审核中 2:已通过
	AuditNames   interface{} // 审核人员们的姓名
	Excellence   interface{} // 是否是优秀报告 0:未被评选为优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告
	AuditEndTime *gtime.Time // 审核完成时间
	OriginUrl    interface{} // 原始文件url
	PdfUrl       interface{} // pdf文件url
	CreateAt     *gtime.Time // 上传时间
	UpdateAt     *gtime.Time // 更新时间
}
