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
	g.Meta        `orm:"table:user, do:true"`
	Id            interface{} // 用户ID
	Passport      interface{} // 账号
	Password      interface{} // 密码
	Salt          interface{} // 加密盐
	Realname      interface{} // 姓名
	Nickname      interface{} // 昵称
	DeptId        interface{} // 部门ID
	Gender        interface{} // 性别 1: 男 2: 女
	Status        interface{} // 状态 0:禁用 1:启用
	Avatar        interface{} // 头像地址
	Mobile        interface{} // 手机号
	Email         interface{} // 用户邮箱
	Remark        interface{} // 备注
	LastLoginIp   interface{} // 最后登录ip
	LastLoginTime *gtime.Time // 最后登录时间
	CreateAt      *gtime.Time // 创建时间
	UpdateAt      *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
