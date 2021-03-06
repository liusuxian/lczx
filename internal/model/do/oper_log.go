// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OperLog is the golang structure of table oper_log for DAO operations like Where/Data.
type OperLog struct {
	g.Meta       `orm:"table:oper_log, do:true"`
	Id           interface{} // 日志ID
	Title        interface{} // 模块标题
	Method       interface{} // 方法名称
	ReqMethod    interface{} // 请求方式
	OperType     interface{} // 操作类别 0:其它 1:后台用户 2:手机端用户
	OperName     interface{} // 操作人员
	DeptName     interface{} // 部门名称
	ReqUrl       interface{} // 请求URL
	OperIp       interface{} // 操作IP地址
	OperLocation interface{} // 操作地点
	ReqParam     interface{} // 请求参数
	JsonResult   interface{} // 返回参数
	Status       interface{} // 操作状态 0:异常 1:正常
	Time         *gtime.Time // 操作时间
}
