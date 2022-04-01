// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure for table dept.
type Dept struct {
	Id        uint64      `json:"id"        description:"部门ID"`           // 部门ID
	ParentId  uint64      `json:"parentId"  description:"父部门id"`          // 父部门id
	Ancestors string      `json:"ancestors" description:"祖级列表"`           // 祖级列表
	Name      string      `json:"name"      description:"部门名称"`           // 部门名称
	Status    uint        `json:"status"    description:"部门状态 0:停用 1:正常"` // 部门状态 0:停用 1:正常
	CreatedBy uint64      `json:"createdBy" description:"创建人"`            // 创建人
	UpdatedBy uint64      `json:"updatedBy" description:"修改人"`            // 修改人
	CreateAt  *gtime.Time `json:"createAt"  description:"创建时间"`           // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  description:"更新时间"`           // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" description:"软删除时间"`          // 软删除时间
}
