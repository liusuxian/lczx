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
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
	"net/http"
)

var gfToken *gtoken.GfToken

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
		_ = req.Response.WriteJsonExit(gtoken.Resp{
			Code: errCode,
			Msg:  errMsg,
		})
	}
	// 通过账号和密码获取用户信息
	logger.Debug(ctx, "loginReq: ", loginReq)
	var user *entity.User
	user, err = User().GetUserByPassportAndPassword(ctx, loginReq.Passport, loginReq.Password)
	if err != nil {
		_ = req.Response.WriteJsonExit(gtoken.Resp{
			Code: code.GetUserFailed.Code(),
			Msg:  code.GetUserFailed.Message(),
		})
	}
	// 判读用户信息
	logger.Debug(ctx, "user: ", user)
	if user == nil {
		_ = req.Response.WriteJsonExit(gtoken.Resp{
			Code: code.UserNotExist.Code(),
			Msg:  code.UserNotExist.Message(),
		})
	}
	// 判断用户状态
	if user.Status == consts.UserStatusDisabled {
		_ = req.Response.WriteJsonExit(gtoken.Resp{
			Code: code.UserDisabled.Code(),
			Msg:  code.UserDisabled.Message(),
		})
	}
	// 设置用户信息到session中
	err = Session().SetUser(ctx, user)
	if err != nil {
		_ = req.Response.WriteJsonExit(gtoken.Resp{
			Code: gcode.CodeInternalError.Code(),
			Msg:  gcode.CodeInternalError.Message(),
		})
	}
	// 唯一标识，扩展参数user data
	return user.Passport, user
}

// 登录返回方法
func loginAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "loginAfter respData: ", respData)
	if !respData.Success() {
		_ = req.Response.WriteJson(respData)
	} else {
		_ = req.Response.WriteJson(&v1.LoginRes{
			Token:   respData.GetString("token"),
			UserKey: respData.GetString("userKey"),
			Uuid:    respData.GetString("uuid"),
		})
	}
}

// 登出验证方法 return true 继续执行，否则结束执行
func logoutBefore(req *ghttp.Request) bool {
	ctx := req.GetCtx()
	logger.Debugf(ctx, "logoutBefore req: %v", req)
	return true
}

// 登出返回方法
func logoutAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "logoutAfter respData: ", respData)
	_ = req.Response.WriteJson(respData)
}

// 认证验证方法 return true 继续执行，否则结束执行
func authBefore(req *ghttp.Request) bool {
	ctx := req.GetCtx()
	logger.Debugf(ctx, "authBefore req: %v", req)
	return true
}

// 认证返回方法
func authAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "authAfter respData: ", respData)
	if req.Method == "OPTIONS" || respData.Success() {
		req.Middleware.Next()
	} else if respData.Code == gtoken.UNAUTHORIZED {
		_ = req.Response.WriteJson(gtoken.Resp{
			Code: http.StatusUnauthorized,
			Msg:  respData.Msg,
			Data: respData.Data,
		})
	} else {
		_ = req.Response.WriteJson(respData)
	}
}
