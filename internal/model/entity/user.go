// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        description:"用户ID"`                                  // 用户ID
	Passport  string      `json:"passport"  description:"账号"`                                    // 账号
	Password  string      `json:"password"  description:"密码"`                                    // 密码
	Realname  string      `json:"realname"  description:"姓名"`                                    // 姓名
	Nickname  string      `json:"nickname"  description:"昵称"`                                    // 昵称
	Gender    uint        `json:"gender"    description:"性别 1: 男 2: 女"`                          // 性别 1: 男 2: 女
	Avatar    string      `json:"avatar"    description:"头像地址"`                                  // 头像地址
	Mobile    string      `json:"mobile"    description:"手机号"`                                   // 手机号
	DeptId    uint        `json:"deptId"    description:"部门ID"`                                  // 部门ID
	RoleId    uint        `json:"roleId"    description:"角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员"` // 角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
	Status    uint        `json:"status"    description:"状态 0:启用 1:禁用"`                          // 状态 0:启用 1:禁用
	CreateAt  *gtime.Time `json:"createAt"  description:"创建时间"`                                  // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  description:"更新时间"`                                  // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" description:"软删除时间"`                                 // 软删除时间
}
