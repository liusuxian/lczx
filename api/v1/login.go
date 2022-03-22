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
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
