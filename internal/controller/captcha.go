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

// Info 获取验证码信息
func (c *cCaptcha) Info(ctx context.Context, req *v1.CaptchaInfoReq) (res *v1.CaptchaInfoRes, err error) {
	var verifyKey, verifyCode string
	verifyKey, verifyCode, err = service.Captcha().GetVerifyImgString()
	if err != nil {
		err = gerror.WrapCode(code.GetCaptchaFailed, err)
		return
	}

	res = &v1.CaptchaInfoRes{
		VerifyKey:  verifyKey,
		VerifyCode: verifyCode,
	}
	return
}
