package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"You first login api"`
	Passport string `json:"passport" v:"required|passport#账号不能为空|账号以字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"账号"` // 账号
	Password string `json:"password" v:"required|password#密码不能为空|密码为任意可见字符，长度在6~18之间" dc:"密码"`             // 密码
}
type LoginRes struct {
	Token string `json:"token" dc:"token"` // token
}
