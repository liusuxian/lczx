// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WdkAttachmentFile is the golang structure of table wdk_attachment_file for DAO operations like Where/Data.
type WdkAttachmentFile struct {
	g.Meta    `orm:"table:wdk_attachment_file, do:true"`
	Id        interface{} // 附件上传记录ID
	Name      interface{} // 附件名
	OriginUrl interface{} // 原始附件url
	PdfUrl    interface{} // pdf附件url
}
