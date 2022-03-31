package v1

import "github.com/gogf/gf/v2/frame/g"

type CaptchaImgReq struct {
	g.Meta `path:"/img" tags:"CaptchaImg" method:"get" summary:"You first /captcha/img api"`
}
type CaptchaImgRes struct {
	VerifyKey  string `json:"verifyKey" dc:"验证码key"` // 验证码key
	VerifyCode string `json:"verifyCode" dc:"验证码"`   // 验证码
}
