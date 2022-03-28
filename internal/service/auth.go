package service

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
	"lczx/utility/response"
	"lczx/utility/utils"
	"net/http"
)

var gfToken *gtoken.GfToken

// InitGfToken 初始化 gfToken
func InitGfToken(ctx context.Context) *gtoken.GfToken {
	if gfToken == nil {
		gfToken = &gtoken.GfToken{
			ServerName:       g.Cfg().MustGet(ctx, "gToken.ServerName").String(),
			CacheMode:        g.Cfg().MustGet(ctx, "gToken.CacheMode").Int8(),
			CacheKey:         g.Cfg().MustGet(ctx, "gToken.CacheKey").String(),
			Timeout:          g.Cfg().MustGet(ctx, "gToken.Timeout").Int(),
			MaxRefresh:       g.Cfg().MustGet(ctx, "gToken.MaxRefresh").Int(),
			TokenDelimiter:   g.Cfg().MustGet(ctx, "gToken.TokenDelimiter").String(),
			EncryptKey:       g.Cfg().MustGet(ctx, "gToken.EncryptKey").Bytes(),
			AuthFailMsg:      g.Cfg().MustGet(ctx, "gToken.AuthFailMsg").String(),
			MultiLogin:       g.Cfg().MustGet(ctx, "gToken.MultiLogin").Bool(),
			LoginPath:        "/login",
			LoginBeforeFunc:  loginBefore,
			LoginAfterFunc:   loginAfter,
			LogoutPath:       "/logout",
			LogoutBeforeFunc: logoutBefore,
			LogoutAfterFunc:  logoutAfter,
			AuthPaths:        g.SliceStr{}, // 这里是按照前缀拦截，拦截
			AuthExcludePaths: g.SliceStr{}, // 不拦截路径
			AuthBeforeFunc:   authBefore,
			AuthAfterFunc:    authAfter,
		}
	}

	return gfToken
}

// GfToken 获取 gfToken
func GfToken() *gtoken.GfToken {
	return gfToken
}

// 登录验证方法 return userKey 用户标识 如果userKey为空，结束执行
func loginBefore(req *ghttp.Request) (string, interface{}) {
	ctx := req.GetCtx()
	var err error
	var loginReq *v1.LoginReq
	if err = req.Parse(&loginReq); err != nil {
		errCode := err.(gvalid.Error).Code().Code()
		errMsg := err.(gvalid.Error).Current().Error()
		response.RespJsonExit(req, errCode, errMsg)
	}
	// 判断验证码是否正确
	captchaDebugVar, err := g.Cfg().Get(ctx, "captcha.debug", false)
	if err != nil {
		response.RespJsonExit(req, code.GetCaptchaCfgFailed.Code(), code.GetCaptchaCfgFailed.Message()+": "+err.Error())
	}
	if !captchaDebugVar.Bool() {
		if !Captcha().VerifyString(loginReq.VerifyKey, loginReq.VerifyCode) {
			response.RespJsonExit(req, code.CaptchaPutErr.Code(), code.CaptchaPutErr.Message())
		}
	}
	// 通过账号和密码获取用户信息
	var user *entity.User
	user, err = User().GetUserByPassportAndPassword(ctx, loginReq.Passport, loginReq.Password)
	if err != nil {
		// 保存登录日志（异步）
		LoginLog().Invoke(req, &entity.LoginLog{Passport: loginReq.Passport, Msg: err.Error()})
		response.RespJsonExit(req, code.GetUserFailed.Code(), code.GetUserFailed.Message()+": "+err.Error())
	}
	// 设置用户信息到session中
	err = Session().SetUser(ctx, user)
	if err != nil {
		// 保存登录日志（异步）
		LoginLog().Invoke(req, &entity.LoginLog{Passport: loginReq.Passport, Msg: "内部错误: " + err.Error()})
		response.RespJsonExit(req, gcode.CodeInternalError.Code(), "内部错误: "+err.Error())
	}
	// 保存登录日志（异步）
	LoginLog().Invoke(req, &entity.LoginLog{Passport: user.Passport, Status: 1, Msg: "登录成功"})
	// 更新用户登录信息
	User().UpdateUserLogin(ctx, user.Id, utils.GetClientIp(req))
	req.SetParam("user", user)
	return user.Passport, user
}

// 登录返回方法
func loginAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "loginAfter: ", respData)
	if !respData.Success() {
		response.RespJson(req, respData.Code, respData.Msg, respData.Data)
	} else {
		token := respData.GetString("token")
		uuid := respData.GetString("uuid")
		var user *entity.User
		_ = req.GetParam("user").Struct(&user)
		// 保存用户在线状态token到数据库
		logger.Debug(ctx, "11: ", uuid, user)
		response.Succ(req, &v1.LoginRes{Token: token})
	}
}

// 登出验证方法 return true 继续执行，否则结束执行
func logoutBefore(req *ghttp.Request) bool {
	ctx := req.GetCtx()
	logger.Debug(ctx, "logoutBefore: ", req.GetClientIp())
	return true
}

// 登出返回方法
func logoutAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "logoutAfter: ", respData)
	_ = Session().RemoveUser(ctx)
	Context().Init(req, nil)
	response.RespJson(req, respData.Code, respData.Msg, respData.Data)
}

// 认证验证方法 return true 继续执行，否则结束执行
func authBefore(req *ghttp.Request) bool {
	ctx := req.GetCtx()
	logger.Debug(ctx, "authBefore: ", req.GetClientIp())
	return true
}

// 认证返回方法
func authAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "authAfter: ", respData)
	if req.Method == "OPTIONS" || respData.Success() {
		req.Middleware.Next()
	} else if respData.Code == gtoken.UNAUTHORIZED {
		response.RespJson(req, http.StatusUnauthorized, respData.Msg, respData.Data)
	} else {
		response.RespJson(req, respData.Code, respData.Msg, respData.Data)
	}
}
