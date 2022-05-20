// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure of table dept for DAO operations like Where/Data.
type Dept struct {
	g.Meta        `orm:"table:dept, do:true"`
	Id            interface{} // 部门ID
	ParentId      interface{} // 父部门ID
	Name          interface{} // 部门名称
	PrincipalUid  interface{} // 负责人用户ID
	PrincipalName interface{} // 负责人姓名
	Status        interface{} // 部门状态 0:停用 1:正常
	CreatedBy     interface{} // 创建人
	UpdatedBy     interface{} // 修改人
	CreateAt      *gtime.Time // 创建时间
	UpdateAt      *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
