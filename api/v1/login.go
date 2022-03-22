package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"You first login api"`
	Passport string `json:"passport" v:"required#账号不能为空" dc:"账号"` // 账号
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"` // 密码
}
type LoginRes struct {
	Token   string `json:"token" dc:"token"`     // token
	UserKey string `json:"userKey" dc:"userKey"` // userKey
	Uuid    string `json:"uuid" dc:"uuid"`       // uuid
}
