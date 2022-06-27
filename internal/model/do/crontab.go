// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Crontab is the golang structure of table crontab for DAO operations like Where/Data.
type Crontab struct {
	g.Meta         `orm:"table:crontab, do:true"`
	Id             interface{} // 任务ID
	Name           interface{} // 任务名称
	Group          interface{} // 任务组名
	Params         interface{} // 参数
	InvokeTarget   interface{} // 调用目标字符串
	CronExpression interface{} // cron执行表达式
	MisfirePolicy  interface{} // 计划执行策略 0:执行一次 1:执行多次
	Status         interface{} // 状态 0:暂停 1:正常
	CreateBy       interface{} // 创建者用户ID
	UpdateBy       interface{} // 更新者用户ID
	Remark         interface{} // 备注
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
