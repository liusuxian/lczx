// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkServiceRecord is the golang structure for table wdk_service_record.
type WdkServiceRecord struct {
	Id           uint64      `json:"id"           description:"服务记录ID"`    // 服务记录ID
	ProjectId    uint64      `json:"projectId"    description:"所属项目ID"`    // 所属项目ID
	ServiceTime  *gtime.Time `json:"serviceTime"  description:"服务时间"`      // 服务时间
	XchName      string      `json:"xchName"      description:"行程涵文件名"`    // 行程涵文件名
	XchOriginUrl string      `json:"xchOriginUrl" description:"原始行程涵url"`  // 原始行程涵url
	XchPdfUrl    string      `json:"xchPdfUrl"    description:"pdf行程涵url"` // pdf行程涵url
	Remark       string      `json:"remark"       description:"备注"`        // 备注
	CreateAt     *gtime.Time `json:"createAt"     description:"创建时间"`      // 创建时间
}
