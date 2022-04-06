package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id            uint64      // 用户ID
	Passport      string      // 账号
	Realname      string      // 姓名
	Nickname      string      // 昵称
	Gender        uint        // 性别 1: 男 2: 女
	Avatar        string      // 头像地址
	Mobile        string      // 手机号
	DeptId        uint        // 部门ID
	Status        uint        // 状态 0:禁用 1:启用
	Email         string      // 用户邮箱
	LastLoginIp   string      // 最后登录ip
	LastLoginTime *gtime.Time // 最后登录时间
	CreateAt      *gtime.Time // 创建时间
}
