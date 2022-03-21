// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Id       interface{} // 用户ID
	Passport interface{} // 账号
	Password interface{} // 密码
	Realname interface{} // 姓名
	Nickname interface{} // 昵称
	Gender   interface{} // 性别 0: 未设置 1: 男 2: 女
	Avatar   interface{} // 头像地址
	Telno    interface{} // 手机号
	Status   interface{} // 状态 0:启用 1:禁用
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}