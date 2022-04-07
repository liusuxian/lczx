package service

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
)

type sCaptcha struct{}

var (
	insCaptcha = sCaptcha{}
)

// Captcha 验证码服务
func Captcha() *sCaptcha {
	return &insCaptcha
}

// GetVerifyImgString 获取字母数字混合验证码
func (s *sCaptcha) GetVerifyImgString() (idKeyC, base64stringC string, err error) {
	driver := &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: 20,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"chromohv.ttf"},
	}
	driver = driver.ConvertFonts()
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	idKeyC, base64stringC, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *sCaptcha) VerifyString(id, answer string) bool {
	driver := &base64Captcha.DriverString{}
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}
