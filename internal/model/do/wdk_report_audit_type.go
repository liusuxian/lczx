// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportAuditType is the golang structure of table wdk_report_audit_type for DAO operations like Where/Data.
type WdkReportAuditType struct {
	g.Meta      `orm:"table:wdk_report_audit_type, do:true"`
	Id          interface{} // 报告ID
	AuditUid    interface{} // 审核员用户ID
	AuditorType interface{} // 审核员类型 0:负责人 1:专家组
	TypeId      interface{} // 报告类型ID
	TypeName    interface{} // 报告类型名称
}
