// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginLog is the golang structure for table login_log.
type LoginLog struct {
	Id       uint64      `json:"id"       description:"访问ID"`           // 访问ID
	Passport string      `json:"passport" description:"登录账号"`           // 登录账号
	Ip       string      `json:"ip"       description:"登录IP地址"`         // 登录IP地址
	Location string      `json:"location" description:"登录地点"`           // 登录地点
	Browser  string      `json:"browser"  description:"浏览器类型"`          // 浏览器类型
	Os       string      `json:"os"       description:"操作系统"`           // 操作系统
	Status   uint        `json:"status"   description:"登录状态 0:失败 1:成功"` // 登录状态 0:失败 1:成功
	Msg      string      `json:"msg"      description:"提示消息"`           // 提示消息
	Time     *gtime.Time `json:"time"     description:"登录时间"`           // 登录时间
	Module   string      `json:"module"   description:"登录模块"`           // 登录模块
}