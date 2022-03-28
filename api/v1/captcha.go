package v1

import "github.com/gogf/gf/v2/frame/g"

type CaptchaImgReq struct {
	g.Meta `path:"/img" tags:"CaptchaImg" method:"get" summary:"You first captcha/img api"`
}
type CaptchaImgRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
