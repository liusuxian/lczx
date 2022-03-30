package v1

import "github.com/gogf/gf/v2/frame/g"

type CaptchaInfoReq struct {
	g.Meta `path:"/info" tags:"CaptchaInfo" method:"get" summary:"You first /captcha/info api"`
}
type CaptchaInfoRes struct {
	VerifyKey  string `json:"verifyKey" dc:"验证码key"` // 验证码key
	VerifyCode string `json:"verifyCode" dc:"验证码"`   // 验证码
}
