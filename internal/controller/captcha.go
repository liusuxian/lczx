package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/service"
	"lczx/utility/response"
)

var (
	Captcha = cCaptcha{}
)

type cCaptcha struct{}

func (c *cCaptcha) Img(ctx context.Context, req *v1.CaptchaImgReq) (res *v1.CaptchaImgRes, err error) {
	idKeyC, base64stringC := service.Captcha().GetVerifyImgString(ctx)
	response.SuccExit(g.RequestFromCtx(ctx), g.MapStrStr{"idKeyC": idKeyC, "base64stringC": base64stringC})
	return
}
