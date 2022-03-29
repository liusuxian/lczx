package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
	"lczx/utility/response"
)

var (
	Captcha = cCaptcha{}
)

type cCaptcha struct{}

func (c *cCaptcha) Img(ctx context.Context, req *v1.CaptchaImgReq) (res *v1.CaptchaImgRes, err error) {
	var idKeyC, base64stringC string
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetCaptchaFailed, err)
		return
	}
	response.SuccExit(g.RequestFromCtx(ctx), g.MapStrStr{"idKeyC": idKeyC, "base64stringC": base64stringC})
	return
}
