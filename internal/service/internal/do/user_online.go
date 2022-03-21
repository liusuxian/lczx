// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserOnline is the golang structure of table user_online for DAO operations like Where/Data.
type UserOnline struct {
	g.Meta       `orm:"table:user_online, do:true"`
	Id           interface{} // 用户ID
	LastLoginAt  *gtime.Time // 最近登录时间
	LastLogoutAt *gtime.Time // 最近登出时间
}
