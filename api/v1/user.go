package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfo 用户信息
type UserInfo struct {
	Id       uint   `json:"id"        dc:"用户ID"`                                  // 用户ID
	Passport string `json:"passport"  dc:"账号"`                                    // 账号
	Realname string `json:"realname"  dc:"姓名"`                                    // 姓名
	Nickname string `json:"nickname"  dc:"昵称"`                                    // 昵称
	Gender   uint   `json:"gender"    dc:"性别 0: 未设置 1: 男 2: 女"`                   // 性别 0: 未设置 1: 男 2: 女
	Avatar   string `json:"avatar"    dc:"头像地址"`                                  // 头像地址
	Mobile   string `json:"mobile"    dc:"手机号"`                                   // 手机号
	DeptId   uint   `json:"deptId"    dc:"部门ID"`                                  // 部门ID
	RoleId   uint   `json:"roleId"    dc:"角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员"` // 角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
}

type UserInfoReq struct {
	g.Meta `path:"/getUserInfo" tags:"GetUserInfo" method:"post" summary:"You first getUserInfo api"`
}
type UserInfoRes struct {
	*UserInfo
}

type UserCreateReq struct {
	g.Meta   `path:"/userCreate" tags:"UserCreate" method:"post" summary:"You first userCreate api"`
	Passport string `json:"passport" v:"required#账号不能为空" dc:"账号"`                                    // 账号
	Realname string `json:"realname" v:"required#姓名不能为空" dc:"姓名"`                                    // 姓名
	Gender   uint   `json:"gender" v:"required#性别不能为空" dc:"性别 0: 未设置 1: 男 2: 女"`                     // 性别 0: 未设置 1: 男 2: 女
	Mobile   string `json:"mobile" dc:"手机号"`                                                         // 手机号
	DeptId   uint   `json:"deptId" v:"required#部门ID不能为空" dc:"部门ID"`                                  // 部门ID
	RoleId   uint   `json:"roleId" v:"required#角色ID不能为空" dc:"角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员"` // 角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
}
type UserCreateRes struct {
	*UserInfo
}
