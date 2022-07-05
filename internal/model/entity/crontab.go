// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Crontab is the golang structure for table crontab.
type Crontab struct {
	Id             uint64      `json:"id"             description:"任务ID"`                 // 任务ID
	Name           string      `json:"name"           description:"任务名称"`                 // 任务名称
	Group          string      `json:"group"          description:"任务组名"`                 // 任务组名
	Params         string      `json:"params"         description:"参数"`                   // 参数
	InvokeTarget   string      `json:"invokeTarget"   description:"调用目标字符串"`              // 调用目标字符串
	CronExpression string      `json:"cronExpression" description:"cron执行表达式"`            // cron执行表达式
	MisfirePolicy  uint        `json:"misfirePolicy"  description:"计划执行策略 0:执行一次 1:执行多次"` // 计划执行策略 0:执行一次 1:执行多次
	Status         uint        `json:"status"         description:"状态 0:暂停 1:正常"`         // 状态 0:暂停 1:正常
	CreateBy       uint64      `json:"createBy"       description:"创建者用户ID"`              // 创建者用户ID
	UpdateBy       uint64      `json:"updateBy"       description:"更新者用户ID"`              // 更新者用户ID
	Remark         string      `json:"remark"         description:"备注"`                   // 备注
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`                 // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`                 // 更新时间
}
