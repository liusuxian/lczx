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
	Gender   uint   // 性别 1: 男 2: 女
	Avatar   string // 头像地址
	Mobile   string // 手机号
	DeptId   uint   // 部门ID
	RoleId   uint   // 角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员
}
