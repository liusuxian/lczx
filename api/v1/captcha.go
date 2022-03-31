package v1

import "github.com/gogf/gf/v2/frame/g"

// CaptchaImgReq 获取验证码图片信息请求参数
type CaptchaImgReq struct {
	g.Meta `path:"/img" tags:"CaptchaImg" method:"get" summary:"You first /captcha/img api"`
}

// CaptchaImgRes 获取验证码图片信息返回参数
type CaptchaImgRes struct {
	VerifyKey  string `json:"verifyKey" dc:"验证码key"` // 验证码key
	VerifyCode string `json:"verifyCode" dc:"验证码"`   // 验证码
}
