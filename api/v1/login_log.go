package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// LoginLogListReq 登录日志列表请求参数
type LoginLogListReq struct {
	g.Meta    `path:"/list" tags:"LoginLogList" method:"get" summary:"You first monitor/loginLog/list api"`
	Passport  string      `json:"passport" v:"regex:^[a-zA-Z]\\w{0,18}$#账号以字母开头，只能包含字母、数字和下划线且长度不能超过18" dc:"账号"` // 账号
	Ip        string      `json:"ip" dc:"IP地址"`                                                                  // IP地址
	Location  string      `json:"location" v:"regex:^[\u4e00-\u9fa5]{0,255}$#登录地点必须为中文且长度不能超过255" dc:"登录地点"`     // 登录地点
	Status    string      `json:"status" v:"in:0,1#登录状态只能是0,1" dc:"登录状态 0:失败 1:成功"`                              // 登录状态 0:失败 1:成功
	StartTime *gtime.Time `json:"startTime" v:"datetime#开始时间不是有效的日期时间" dc:"开始时间"`                                // 开始时间
	EndTime   *gtime.Time `json:"endTime" v:"datetime#结束时间不是有效的日期时间" dc:"结束时间"`                                  // 结束时间
	SortName  string      `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`           // 排序字段
	SortOrder string      `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`          // 排序方式
	CurPage   int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`          // 当前页码
	PageSize  int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`         // 每页数量
}

// LoginLogListRes 登录日志列表返回参数
type LoginLogListRes struct {
	CurPage int                `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int                `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.LoginLog `json:"list" dc:"登录日志列表"`  // 登录日志列表
}

// LoginLogDeleteReq 删除登录日志请求参数
type LoginLogDeleteReq struct {
	g.Meta `path:"/delete" tags:"LoginLogDelete" method:"delete" summary:"You first monitor/loginLog/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#登录日志ID列表不能为空" dc:"登录日志ID列表"` // 登录日志ID列表
}

// LoginLogDeleteRes 删除登录日志返回参数
type LoginLogDeleteRes struct {
}

// LoginLogClearReq 清除登录日志请求参数
type LoginLogClearReq struct {
	g.Meta `path:"/clear" tags:"LoginLogClear" method:"delete" summary:"You first monitor/loginLog/clear api"`
}

// LoginLogClearRes 清除登录日志返回参数
type LoginLogClearRes struct {
}
