// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkAuditRecord is the golang structure of table wdk_audit_record for DAO operations like Where/Data.
type WdkAuditRecord struct {
	g.Meta    `orm:"table:wdk_audit_record, do:true"`
	AuditUid  interface{} // 审核员用户ID
	FileId    interface{} // 需要审核的文件ID
	Status    interface{} // 审核状态 1:审核中 2:已通过 3:未通过
	AuditTime *gtime.Time // 审核时间
	AuditName interface{} // 审核员姓名
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 软删除时间
}
