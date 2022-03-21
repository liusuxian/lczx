// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserExt is the golang structure of table user_ext for DAO operations like Where/Data.
type UserExt struct {
	g.Meta    `orm:"table:user_ext, do:true"`
	Id        interface{} // 用户ID
	BranchId  interface{} // 部门ID
	ZskRoleId interface{} // 知识库角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
	WdkRoleId interface{} // 文档库角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
}
