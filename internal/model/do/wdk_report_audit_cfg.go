// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WdkReportAuditCfg is the golang structure of table wdk_report_audit_cfg for DAO operations like Where/Data.
type WdkReportAuditCfg struct {
	g.Meta    `orm:"table:wdk_report_audit_cfg, do:true"`
	Id        interface{} // 报告类型ID
	AuditUid  interface{} // 审核员用户ID
	TypeName  interface{} // 报告类型名称
	AuditName interface{} // 审核员姓名
}
