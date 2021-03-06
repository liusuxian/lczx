package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	Captcha = cCaptcha{}
)

type cCaptcha struct{}

// Img 获取验证码图片信息
func (c *cCaptcha) Img(ctx context.Context, req *v1.CaptchaImgReq) (res *v1.CaptchaImgRes, err error) {
	var verifyKey, verifyCode string
	verifyKey, verifyCode, err = service.Captcha().GetVerifyImgString()
	if err != nil {
		err = gerror.WrapCode(code.GetCaptchaImgFailed, err)
		return
	}

	res = &v1.CaptchaImgRes{
		VerifyKey:  verifyKey,
		VerifyCode: verifyCode,
	}
	return
}
