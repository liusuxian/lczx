package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

type UserOnlineListReq struct {
	g.Meta   `path:"/list" tags:"UserOnlineList" method:"post" summary:"You first userOnline/list api"`
	Passport string `json:"passport" dc:"账号"`                                                        // 账号
	Ip       string `json:"ip" dc:"IP地址"`                                                            // IP地址
	CurPage  int    `json:"curPage" v:"required|regex:^[1-9][0-9]*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`  // 当前页码
	PageSize int    `json:"pageSize" v:"required|regex:^[1-9][0-9]*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"` // 每页数量
}
type UserOnlineListRes struct {
	CurPage int                  `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int                  `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.UserOnline `json:"list" dc:"用户在线列表"`  // 用户在线列表
}

type UserOnlineForceLogoutReq struct {
	g.Meta `path:"/forceLogout" tags:"UserOnlineForceLogout" method:"post" summary:"You first userOnline/forceLogout api"`
	Ids    []int `json:"ids" v:"required#ID列表不能为空" dc:"ID列表"` // ID列表
}
type UserOnlineForceLogoutRes struct {
}
