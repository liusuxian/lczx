package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

type LoginLogListReq struct {
	g.Meta    `path:"/list" tags:"UserOnlineList" method:"post" summary:"You first userOnline/list api"`
	Passport  string      `json:"passport" dc:"账号"`                                                        // 账号
	Ip        string      `json:"ip" dc:"IP地址"`                                                            // IP地址
	Location  string      `json:"location" dc:"登录地点"`                                                      // 登录地点
	Status    uint        `json:"status" dc:"登录状态 0:失败 1:成功 2:全部"`                                         // 登录状态 0:失败 1:成功 2:全部
	StartTime *gtime.Time `json:"startTime" v:"datetime#开始时间不是有效的日期时间" dc:"开始时间"`                          // 开始时间
	EndTime   *gtime.Time `json:"endTime" v:"datetime#结束时间不是有效的日期时间" dc:"结束时间"`                            // 结束时间
	SortName  string      `json:"sortName" dc:"排序字段"`                                                      // 排序字段
	SortOrder string      `json:"sortOrder" dc:"排序方式"`                                                     // 排序方式
	CurPage   int         `json:"curPage" v:"required|regex:^[1-9][0-9]*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`  // 当前页码
	PageSize  int         `json:"pageSize" v:"required|regex:^[1-9][0-9]*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"` // 每页数量
}
type LoginLogListRes struct {
	CurPage int                `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int                `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.LoginLog `json:"list" dc:"登录日志列表"`  // 登录日志列表
}
