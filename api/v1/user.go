package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfo 用户信息
type UserInfo struct {
	Id       uint64 `json:"id"        dc:"用户ID"`                                  // 用户ID
	Passport string `json:"passport"  dc:"账号"`                                    // 账号
	Realname string `json:"realname"  dc:"姓名"`                                    // 姓名
	Nickname string `json:"nickname"  dc:"昵称"`                                    // 昵称
	Gender   uint   `json:"gender"    dc:"性别 1: 男 2: 女"`                          // 性别 1: 男 2: 女
	Avatar   string `json:"avatar"    dc:"头像地址"`                                  // 头像地址
	Mobile   string `json:"mobile"    dc:"手机号"`                                   // 手机号
	DeptId   uint   `json:"deptId"    dc:"部门ID"`                                  // 部门ID
	RoleId   uint   `json:"roleId"    dc:"角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员"` // 角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
}

type UserInfoReq struct {
	g.Meta `path:"/info" tags:"UserInfo" method:"post" summary:"You first /user/info api"`
}
type UserInfoRes struct {
	User *UserInfo `json:"user" dc:"用户信息"` // 用户信息
}

type UserAddReq struct {
	g.Meta   `path:"/add" tags:"UserAdd" method:"post" summary:"You first /user/add api"`
	Passport string `json:"passport" v:"required|passport#账号不能为空|账号以字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"账号"`                           // 账号
	Realname string `json:"realname" v:"required|max-length:10|regex:^[\u4e00-\u9fa5]+$#姓名不能为空|姓名不能超过10个字|姓名必须为纯中文" dc:"姓名"`         // 姓名
	Gender   uint   `json:"gender" v:"required|in:1,2#性别不能为空|性别只能是1,2" dc:"性别 1: 男 2: 女"`                                            // 性别 1: 男 2: 女
	Mobile   string `json:"mobile" v:"phone#不是有效的手机号码" dc:"手机号"`                                                                     // 手机号
	DeptId   uint   `json:"deptId" v:"required|regex:^[1-9][0-9]*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"`                                   // 部门ID
	RoleId   uint   `json:"roleId" v:"required|in:0,900,1000#角色ID不能为空|角色ID只能是0,900,1000" dc:"角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员"` // 角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
}
type UserAddRes struct {
	User *UserInfo `json:"user" dc:"用户信息"` // 用户信息
}
