package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// CrontabClientOptionsReq 获取定时任务客户端选项请求参数
type CrontabClientOptionsReq struct {
	g.Meta `path:"/clientOptions" tags:"CrontabClientOptions" method:"get" summary:"You first monitor/crontab/clientOptions api"`
}

// CrontabClientOptionsRes 获取定时任务客户端选项返回参数
type CrontabClientOptionsRes struct {
	GroupList        []*CrontabClientOption `json:"groupList" dc:"任务组名选项列表"`           // 任务组名选项列表
	StatusList       []*CrontabClientOption `json:"statusList" dc:"任务状态选项列表"`          // 任务状态选项列表
	InvokeTargetList []*CrontabClientOption `json:"invokeTargetList" dc:"调用目标字符串选项列表"` // 调用目标字符串选项列表
}

// CrontabClientOption 定时任务客户端选项
type CrontabClientOption struct {
	Name  string `json:"name" dc:"选项显示名"` // 选项显示名
	Value string `json:"value" dc:"选项值"`  // 选项值
}

// CrontabListReq 定时任务列表请求参数
type CrontabListReq struct {
	g.Meta    `path:"/list" tags:"CrontabList" method:"get" summary:"You first monitor/crontab/list api"`
	Name      string `json:"name" v:"max-length:20#任务名称长度不能超过20" dc:"任务名称"`                         // 任务名称
	Group     string `json:"group" v:"max-length:20#任务组名长度不能超过20" dc:"任务组名"`                        // 任务组名
	Status    string `json:"status" v:"in:0,1#状态只能是0,1" dc:"状态 0:暂停 1:正常"`                          // 状态 0:暂停 1:正常
	SortName  string `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`   // 排序字段
	SortOrder string `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`  // 排序方式
	CurPage   int    `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`  // 当前页码
	PageSize  int    `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"` // 每页数量
}

// CrontabListRes 定时任务列表返回参数
type CrontabListRes struct {
	CurPage int               `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int               `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.Crontab `json:"list" dc:"定时任务列表"`  // 定时任务列表
}

// CrontabAddReq 添加任务请求参数
type CrontabAddReq struct {
	g.Meta         `path:"/add" tags:"CrontabAdd" method:"post" summary:"You first monitor/crontab/add api"`
	Name           string `json:"name" v:"required|max-length:20#任务名称不能为空|任务名称长度不能超过20" dc:"任务名称"`                            // 任务名称
	Group          string `json:"group" v:"required|max-length:20#任务组名不能为空|任务组名长度不能超过20" dc:"任务组名"`                           // 任务组名
	Params         string `json:"params" v:"max-length:255#参数长度不能超过255" dc:"参数"`                                              // 参数
	InvokeTarget   string `json:"invokeTarget" v:"required|max-length:255#调用目标字符串不能为空|调用目标字符串长度不能超过255" dc:"调用目标字符串"`         // 调用目标字符串
	CronExpression string `json:"cronExpression" v:"required|max-length:255#cron执行表达式不能为空|cron执行表达式长度不能超过255" dc:"cron执行表达式"` // cron执行表达式
	MisfirePolicy  uint   `json:"misfirePolicy" v:"required|in:0,1#计划执行策略不能为空|计划执行策略只能是0,1" dc:"计划执行策略 0:执行一次 1:执行多次"`        // 计划执行策略 0:执行一次 1:执行多次
	Status         uint   `json:"status" v:"required|in:0,1#状态不能为空|状态只能是0,1" dc:"状态 0:暂停 1:正常"`                               // 状态 0:暂停 1:正常
	Remark         string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                              // 备注
}

// CrontabAddRes 添加任务返回参数
type CrontabAddRes struct {
}

// CrontabInfoReq 任务信息请求参数
type CrontabInfoReq struct {
	g.Meta `path:"/info" tags:"CrontabInfo" method:"get" summary:"You first monitor/crontab/info api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#任务ID不能为空|任务ID必须为正整数" dc:"任务ID"` // 任务ID
}

// CrontabInfoRes 任务信息返回参数
type CrontabInfoRes struct {
	Info *entity.Crontab `json:"info" dc:"任务信息"` // 任务信息
}

// CrontabEditReq 修改任务请求参数
type CrontabEditReq struct {
	g.Meta         `path:"/edit" tags:"CrontabEdit" method:"put" summary:"You monitor/crontab/edit api"`
	Id             uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#任务ID不能为空|任务ID必须为正整数" dc:"任务ID"`                            // 任务ID
	Name           string `json:"name" v:"required|max-length:20#任务名称不能为空|任务名称长度不能超过20" dc:"任务名称"`                            // 任务名称
	Group          string `json:"group" v:"required|max-length:20#任务组名不能为空|任务组名长度不能超过20" dc:"任务组名"`                           // 任务组名
	Params         string `json:"params" v:"max-length:255#参数长度不能超过255" dc:"参数"`                                              // 参数
	InvokeTarget   string `json:"invokeTarget" v:"required|max-length:255#调用目标字符串不能为空|调用目标字符串长度不能超过255" dc:"调用目标字符串"`         // 调用目标字符串
	CronExpression string `json:"cronExpression" v:"required|max-length:255#cron执行表达式不能为空|cron执行表达式长度不能超过255" dc:"cron执行表达式"` // cron执行表达式
	MisfirePolicy  uint   `json:"misfirePolicy" v:"required|in:0,1#计划执行策略不能为空|计划执行策略只能是0,1" dc:"计划执行策略 0:执行一次 1:执行多次"`        // 计划执行策略 0:执行一次 1:执行多次
	Status         uint   `json:"status" v:"required|in:0,1#状态不能为空|状态只能是0,1" dc:"状态 0:暂停 1:正常"`                               // 状态 0:暂停 1:正常
	Remark         string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                              // 备注
}

// CrontabEditRes 修改任务返回参数
type CrontabEditRes struct {
}

// CrontabStartReq 开启任务请求参数
type CrontabStartReq struct {
	g.Meta `path:"/start" tags:"CrontabStart" method:"put" summary:"You first monitor/crontab/start api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#任务ID不能为空|任务ID必须为正整数" dc:"任务ID"` // 任务ID
}

// CrontabStartRes 开启任务返回参数
type CrontabStartRes struct {
}

// CrontabStopReq 停止任务请求参数
type CrontabStopReq struct {
	g.Meta `path:"/stop" tags:"CrontabStop" method:"put" summary:"You first monitor/crontab/stop api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#任务ID不能为空|任务ID必须为正整数" dc:"任务ID"` // 任务ID
}

// CrontabStopRes 停止任务返回参数
type CrontabStopRes struct {
}

// CrontabRunReq 执行任务请求参数
type CrontabRunReq struct {
	g.Meta `path:"/run" tags:"CrontabRun" method:"put" summary:"You first monitor/crontab/run api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#任务ID不能为空|任务ID必须为正整数" dc:"任务ID"` // 任务ID
}

// CrontabRunRes 执行任务返回参数
type CrontabRunRes struct {
}

// CrontabDeleteReq 删除任务请求参数
type CrontabDeleteReq struct {
	g.Meta `path:"/delete" tags:"CrontabDelete" method:"delete" summary:"You first monitor/crontab/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#任务ID列表不能为空" dc:"任务ID列表"` // 任务ID列表
}

// CrontabDeleteRes 删除任务返回参数
type CrontabDeleteRes struct {
}
