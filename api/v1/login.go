package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta     `path:"/login" tags:"Login" method:"get" summary:"You first login api"`
	Passport   string `json:"passport" v:"required|passport#账号不能为空|账号以字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"账号"`           // 账号
	Password   string `json:"password" v:"required|password#密码不能为空|密码为任意可见字符，长度在6~18之间" dc:"密码"`                       // 密码
	VerifyKey  string `json:"verifyKey" v:"required|size:20|regex:^[a-zA-Z0-9]*$#验证码key不能为空|验证码key长度必须为20|验证码key格式错误"` // 验证码key
	VerifyCode string `json:"verifyCode" v:"required|size:4|regex:^[a-zA-Z0-9]*$#验证码不能为空|验证码长度必须为4|验证码格式错误"`           // 验证码
}
type LoginRes struct {
	Token string `json:"token" dc:"token"` // token
}
