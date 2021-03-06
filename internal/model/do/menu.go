// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta     `orm:"table:menu, do:true"`
	Id         interface{} // 规则ID
	ParentId   interface{} // 父规则ID
	Rule       interface{} // 权限规则
	Name       interface{} // 菜单名称
	Condition  interface{} // 条件
	MenuType   interface{} // 类型 0:目录 1:菜单 2:按钮
	Status     interface{} // 菜单状态 0:停用 1:正常
	JumpPath   interface{} // 跳转路由
	IsFrame    interface{} // 是否外链 1是 0否
	ModuleType interface{} // 所属模块
	Remark     interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 软删除时间
}
