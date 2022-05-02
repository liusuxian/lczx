// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WdkReportType is the golang structure for table wdk_report_type.
type WdkReportType struct {
	Id          uint64 `json:"id"          description:"报告ID"`   // 报告ID
	TypeId      uint64 `json:"typeId"      description:"报告类型ID"` // 报告类型ID
	TypeName    string `json:"typeName"    description:"报告类型名称"` // 报告类型名称
	ReportName  string `json:"reportName"  description:"报告名称"`   // 报告名称
	ProjectId   uint64 `json:"projectId"   description:"所属项目ID"` // 所属项目ID
	ProjectName string `json:"projectName" description:"所属项目名称"` // 所属项目名称
}
