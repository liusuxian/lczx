package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id       uint64 // 用户ID
	Passport string // 账号
	Realname string // 姓名
	Nickname string // 昵称
	Avatar   string // 头像地址
	DeptId   uint64 // 部门ID
	Status   uint   // 状态 0:禁用 1:启用
	IsAdmin  uint   // 是否后台管理员 0:否 1:是
}
