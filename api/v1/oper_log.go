package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// OperLogListReq 操作日志列表请求参数
type OperLogListReq struct {
	g.Meta    `path:"/list" tags:"OperLogList" method:"get" summary:"You first monitor/operLog/list api"`
	Title     string      `json:"title" dc:"模块标题"`                                                           // 模块标题
	OperName  string      `json:"operName" v:"regex:^[\u4e00-\u9fa5]{0,20}$#操作人员必须为中文且长度不能超过20" dc:"操作人员"`   // 操作人员
	ReqMethod string      `json:"reqMethod" v:"in:GET,POST,PUT,DELETE#请求方式只能是GET,POST,PUT,DELETE" dc:"请求方式"` // 请求方式
	Status    string      `json:"status" v:"in:0,1#操作状态只能是0,1" dc:"操作状态 0:异常 1:正常"`                          // 操作状态 0:异常 1:正常
	StartTime *gtime.Time `json:"startTime" v:"datetime#开始时间不是有效的日期时间" dc:"开始时间"`                            // 开始时间
	EndTime   *gtime.Time `json:"endTime" v:"datetime#结束时间不是有效的日期时间" dc:"结束时间"`                              // 结束时间
	SortName  string      `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`       // 排序字段
	SortOrder string      `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`      // 排序方式
	CurPage   int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`      // 当前页码
	PageSize  int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`     // 每页数量
}

// OperLogListRes 操作日志列表返回参数
type OperLogListRes struct {
	CurPage int               `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int               `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.OperLog `json:"list" dc:"操作日志列表"`  // 操作日志列表
}

// OperLogDeleteReq 删除操作日志请求参数
type OperLogDeleteReq struct {
	g.Meta `path:"/delete" tags:"OperLogDelete" method:"delete" summary:"You first monitor/operLog/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#操作日志ID列表不能为空" dc:"操作日志ID列表"` // 操作日志ID列表
}

// OperLogDeleteRes 删除操作日志返回参数
type OperLogDeleteRes struct {
}

// OperLogClearReq 清除操作日志请求参数
type OperLogClearReq struct {
	g.Meta `path:"/clear" tags:"OperLogClear" method:"delete" summary:"You first monitor/operLog/clear api"`
}

// OperLogClearRes 清除操作日志返回参数
type OperLogClearRes struct {
}
