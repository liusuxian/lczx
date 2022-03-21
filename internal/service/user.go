package service

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"lczx/utility/logger"
)

// LoginBeforeFunc 登录验证方法 return userKey 用户标识 如果userKey为空，结束执行
func LoginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	passport := r.Get("passport").String()
	password := r.Get("password").String()
	logger.Info(r.GetCtx(), "LoginBeforeFunc: ", passport, password)

	if passport == "" || password == "" {
		_ = r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return passport, "1"
}

// LoginAfterFunc 登录返回方法
func LoginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	logger.Info(r.GetCtx(), "LoginAfterFunc: ", respData)
	_ = r.Response.WriteJson(respData)
}

// LogoutBeforeFunc 登出验证方法 return true 继续执行，否则结束执行
func LogoutBeforeFunc(r *ghttp.Request) bool {
	logger.Info(r.GetCtx(), "LogoutBeforeFunc: ", r.GetBody())
	return true
}

// LogoutAfterFunc 登出返回方法
func LogoutAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	logger.Info(r.GetCtx(), "LogoutAfterFunc: ", respData)
	_ = r.Response.WriteJson(respData)
}

// AuthBeforeFunc 认证验证方法 return true 继续执行，否则结束执行
func AuthBeforeFunc(r *ghttp.Request) bool {
	logger.Info(r.GetCtx(), "AuthBeforeFunc: ", r.GetBody())
	return true
}

// AuthAfterFunc 认证返回方法
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	logger.Info(r.GetCtx(), "AuthAfterFunc: ", respData)
	_ = r.Response.WriteJson(respData)
}
