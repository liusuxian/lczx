// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkFiletypeCfg is the golang structure for table wdk_filetype_cfg.
type WdkFiletypeCfg struct {
	Id        uint64      `json:"id"        description:"上传文件类型ID"`                                // 上传文件类型ID
	Name      string      `json:"name"      description:"类型名称"`                                    // 类型名称
	Multiple  uint        `json:"multiple"  description:"是否同时存在多个文件 0:否 1:是"`                      // 是否同时存在多个文件 0:否 1:是
	Audit     uint        `json:"audit"     description:"是否需要审核 0:不需要 1:需要"`                       // 是否需要审核 0:不需要 1:需要
	Step      uint        `json:"step"      description:"所属项目阶段 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘"` // 所属项目阶段 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	CreateAt  *gtime.Time `json:"createAt"  description:"创建时间"`                                    // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  description:"更新时间"`                                    // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" description:"软删除时间"`                                   // 软删除时间
}
