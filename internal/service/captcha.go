// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

type ICaptcha interface {
	GetVerifyImgString() (idKeyC, base64stringC string, err error)
	VerifyString(id, answer string) bool
}

var localCaptcha ICaptcha

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
