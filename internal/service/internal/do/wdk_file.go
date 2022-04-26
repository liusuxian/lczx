// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkFile is the golang structure of table wdk_file for DAO operations like Where/Data.
type WdkFile struct {
	g.Meta       `orm:"table:wdk_file, do:true"`
	FileId       interface{} // 上传文件ID
	ProjectId    interface{} // 所属项目ID
	Name         interface{} // 文件名
	CreateBy     interface{} // 上传者用户ID
	CreateName   interface{} // 上传者姓名
	AuditStatus  interface{} // 审核状态 0:不需要审核 1:审核中 2:已通过 3:未通过
	AuditNames   interface{} // 审核人员们的姓名
	AuditEndTime *gtime.Time // 审核完成时间
	Step         interface{} // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	Excellence   interface{} // 是否是优秀报告 0:无该属性 1:被推荐为优秀报告 2:未被评选为优秀报告 3:已被评选为优秀报告
	OriginUrl    interface{} // 原始文件url
	PdfUrl       interface{} // pdf文件url
	CreateAt     *gtime.Time // 上传时间
	UpdateAt     *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 软删除时间
}
