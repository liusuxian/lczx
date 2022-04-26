// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkReport is the golang structure for table wdk_report.
type WdkReport struct {
	Id           uint64      `json:"id"           description:"报告ID"`                                       // 报告ID
	ProjectId    uint64      `json:"projectId"    description:"所属项目ID"`                                     // 所属项目ID
	Name         string      `json:"name"         description:"报告名称"`                                       // 报告名称
	CreateBy     uint64      `json:"createBy"     description:"上传者用户ID"`                                    // 上传者用户ID
	CreateName   string      `json:"createName"   description:"上传者姓名"`                                      // 上传者姓名
	AuditStatus  uint        `json:"auditStatus"  description:"审核状态 0:未通过 1:审核中 2:已通过"`                     // 审核状态 0:未通过 1:审核中 2:已通过
	AuditNames   string      `json:"auditNames"   description:"审核人员们的姓名"`                                   // 审核人员们的姓名
	Excellence   uint        `json:"excellence"   description:"是否是优秀报告 0:未被评选为优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告"` // 是否是优秀报告 0:未被评选为优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告
	AuditEndTime *gtime.Time `json:"auditEndTime" description:"审核完成时间"`                                     // 审核完成时间
	OriginUrl    string      `json:"originUrl"    description:"原始文件url"`                                    // 原始文件url
	PdfUrl       string      `json:"pdfUrl"       description:"pdf文件url"`                                   // pdf文件url
	CreateAt     *gtime.Time `json:"createAt"     description:"上传时间"`                                       // 上传时间
	UpdateAt     *gtime.Time `json:"updateAt"     description:"更新时间"`                                       // 更新时间
}
