package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// UserOnlineListReq 在线用户列表请求参数
type UserOnlineListReq struct {
	g.Meta   `path:"/list" tags:"UserOnlineList" method:"get" summary:"You first monitor/userOnline/list api"`
	Passport string `json:"passport" v:"regex:^[a-zA-Z]\\w{0,18}$#账号以字母开头，只能包含字母、数字和下划线且长度不能超过18" dc:"账号"` // 账号
	Ip       string `json:"ip" v:"regex:^[\\d\\.]{0,15}$#IP地址只能包含数字和.号且长度不能超过15" dc:"IP地址"`                // IP地址
	CurPage  int    `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`          // 当前页码
	PageSize int    `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`         // 每页数量
}

// UserOnlineListRes 在线用户列表返回参数
type UserOnlineListRes struct {
	CurPage int                  `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int                  `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.UserOnline `json:"list" dc:"用户在线列表"`  // 用户在线列表
}

// UserOnlineForceLogoutReq 强退在线用户请求参数
type UserOnlineForceLogoutReq struct {
	g.Meta `path:"/forceLogout" tags:"UserOnlineForceLogout" method:"put" summary:"You first monitor/userOnline/forceLogout api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#在线用户ID列表不能为空" dc:"在线用户ID列表"` // 在线用户ID列表
}

// UserOnlineForceLogoutRes 强退在线用户返回参数
type UserOnlineForceLogoutRes struct {
}
