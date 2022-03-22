package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfoReq struct {
	g.Meta `path:"/getUserInfo" tags:"GetUserInfo" method:"post" summary:"You first getUserInfo api"`
}
type UserInfoRes struct {
	g.Meta   `mime:"text/html" type:"string" example:"<html/>"`
	Id       uint   `json:"id" dc:"用户ID"`                    // 用户ID
	Passport string `json:"passport" dc:"账号"`                // 账号
	Realname string `json:"realname" dc:"姓名"`                // 姓名
	Nickname string `json:"nickname" dc:"昵称"`                // 昵称
	Gender   uint   `json:"gender" dc:"性别 0: 未设置 1: 男 2: 女"` // 性别 0: 未设置 1: 男 2: 女
	Avatar   string `json:"avatar" dc:"头像地址"`                // 头像地址
	Mobile   string `json:"mobile" dc:"手机号"`                 // 手机号
}

type UserDeptReq struct {
	g.Meta `path:"/getUserDept" tags:"GetUserDept" method:"post" summary:"You first getUserDept api"`
}
type UserDeptRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	Id     uint   `json:"id" dc:"部门ID"`   // 部门ID
	Name   string `json:"name" dc:"部门名称"` // 部门名称
}
