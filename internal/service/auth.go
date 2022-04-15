package service

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/mssola/user_agent"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
	"lczx/utility/response"
	"lczx/utility/utils"
	"net/http"
	"strings"
)

type sAuth struct {
	token *gtoken.GfToken
}

var (
	authCtx  = gctx.New()
	insAuth  = sAuth{}
	insToken = &gtoken.GfToken{
		ServerName:       g.Cfg().MustGet(authCtx, "gToken.ServerName").String(),
		CacheMode:        g.Cfg().MustGet(authCtx, "gToken.CacheMode").Int8(),
		CacheKey:         g.Cfg().MustGet(authCtx, "gToken.CacheKey").String(),
		Timeout:          g.Cfg().MustGet(authCtx, "gToken.Timeout").Int(),
		MaxRefresh:       g.Cfg().MustGet(authCtx, "gToken.MaxRefresh").Int(),
		TokenDelimiter:   g.Cfg().MustGet(authCtx, "gToken.TokenDelimiter").String(),
		EncryptKey:       g.Cfg().MustGet(authCtx, "gToken.EncryptKey").Bytes(),
		AuthFailMsg:      g.Cfg().MustGet(authCtx, "gToken.AuthFailMsg").String(),
		MultiLogin:       g.Cfg().MustGet(authCtx, "gToken.MultiLogin").Bool(),
		LoginPath:        "/login",
		LoginBeforeFunc:  insAuth.loginBefore,
		LoginAfterFunc:   insAuth.loginAfter,
		LogoutPath:       "/logout",
		LogoutBeforeFunc: insAuth.logoutBefore,
		LogoutAfterFunc:  insAuth.logoutAfter,
		AuthPaths:        g.SliceStr{},         // 这里是按照前缀拦截，拦截
		AuthExcludePaths: g.SliceStr{"/login"}, // 不拦截路径
		AuthBeforeFunc:   insAuth.authBefore,
		AuthAfterFunc:    insAuth.authAfter,
	}
)

func init() {
	insAuth.token = insToken
}

// Auth 验证服务
func Auth() *sAuth {
	return &insAuth
}

// Token 获取 gfToken
func (s *sAuth) Token() *gtoken.GfToken {
	return s.token
}

// GetUuidUserKeyByToken 通过token获取uuid和userKey
func (s *sAuth) GetUuidUserKeyByToken(ctx context.Context, token string) (uuid, userKey string) {
	decryptToken := s.token.DecryptToken(ctx, token)
	if !decryptToken.Success() {
		return
	}
	uuid = decryptToken.GetString("uuid")
	userKey = decryptToken.GetString("userKey")
	return
}

// UserIsOnline 判断用户是否在线
func (s *sAuth) UserIsOnline(ctx context.Context, token string) bool {
	uuid, userKey := s.GetUuidUserKeyByToken(ctx, token)
	cacheKey := s.token.CacheKey + userKey
	switch s.token.CacheMode {
	case gtoken.CacheModeCache:
		userCacheValue, err := gcache.Get(ctx, cacheKey)
		if err != nil {
			logger.Error(ctx, "GetCache gcache Error: ", err.Error())
			return false
		}
		if userCacheValue == nil {
			return false
		}
		return true
	case gtoken.CacheModeRedis:
		userCacheValue, err := g.DB().GetCache().Get(ctx, cacheKey)
		if err != nil {
			logger.Error(ctx, "GetCache redis Error: ", err.Error())
			return false
		}
		if userCacheValue == nil {
			return false
		}
		if uuid != userCacheValue.Map()["uuid"] {
			return false
		}
		return true
	}
	return false
}

// CheckUserOnline 检查在线用户
func (s *sAuth) CheckUserOnline(ctx context.Context) {
	req := &v1.UserOnlineListReq{
		CurPage:  1,
		PageSize: 50,
	}
	var total int
	for {
		var list []*entity.UserOnline
		var err error
		total, list, err = UserOnline().GetOnlineList(ctx, req)
		if err != nil {
			logger.Error(ctx, "CheckUserOnline GetOnlineList Error: ", err.Error())
			break
		}
		if list == nil {
			break
		}
		for _, v := range list {
			if ok := s.UserIsOnline(ctx, v.Token); !ok {
				UserOnline().DeleteOnlineByToken(ctx, v.Token)
			}
		}
		if req.CurPage*req.PageSize >= total {
			break
		}
		req.CurPage++
	}
}

// 登录验证方法 return userKey 用户标识 如果userKey为空，结束执行
func (s *sAuth) loginBefore(req *ghttp.Request) (string, any) {
	ctx := req.GetCtx()
	var err error
	var loginReq *v1.LoginReq
	if err = req.Parse(&loginReq); err != nil {
		errCode := err.(gvalid.Error).Code().Code()
		errMsg := err.(gvalid.Error).Current().Error()
		response.RespJsonExit(req, errCode, errMsg)
	}
	// 判断验证码是否正确
	if !Captcha().VerifyString(loginReq.VerifyKey, loginReq.VerifyCode) {
		response.RespJsonExitByGcode(req, code.CaptchaPutErr)
	}
	// 获取客户端IP
	ip := utils.GetClientIp(req)
	// 获取ip所属城市
	location := utils.GetCityByIp(ctx, ip)
	// 获取 User-Agent
	userAgent := req.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	// 获取浏览器类型
	explorer, _ := ua.Browser()
	// 获取操作系统
	os := ua.OS()
	// 通过账号和密码获取用户信息
	var user *entity.User
	user, err = User().GetUserByPassportAndPassword(ctx, loginReq.Passport, loginReq.Password)
	if err != nil {
		// 保存登录日志（异步）
		LoginLog().Invoke(ctx, &entity.LoginLog{
			Passport: loginReq.Passport,
			Ip:       ip,
			Location: location,
			Browser:  explorer,
			Os:       os,
			Status:   consts.LoginFailed,
			Msg:      err.Error(),
			Time:     gtime.Now(),
			Module:   "系统登录",
		})
		response.RespJsonExit(req, code.GetUserFailed.Code(), code.GetUserFailed.Message()+": "+err.Error())
	}
	// 设置用户信息到session中
	err = Session().SetUser(ctx, user)
	if err != nil {
		// 保存登录日志（异步）
		LoginLog().Invoke(ctx, &entity.LoginLog{
			Passport: loginReq.Passport,
			Ip:       ip,
			Location: location,
			Browser:  explorer,
			Os:       os,
			Status:   consts.LoginFailed,
			Msg:      "内部错误: " + err.Error(),
			Time:     gtime.Now(),
			Module:   "系统登录",
		})
		response.RespJsonExit(req, gcode.CodeInternalError.Code(), "内部错误: "+err.Error())
	}
	// 更新用户登录信息
	User().UpdateUserLogin(ctx, user.Id, ip)
	// 保存登录日志（异步）
	LoginLog().Invoke(ctx, &entity.LoginLog{
		Passport: loginReq.Passport,
		Ip:       ip,
		Location: location,
		Browser:  explorer,
		Os:       os,
		Status:   consts.LoginSucc,
		Msg:      "登录成功",
		Time:     gtime.Now(),
		Module:   "系统登录",
	})
	req.SetParam("user", user)
	return user.Passport, user
}

// 登录返回方法
func (s *sAuth) loginAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "loginAfter: ", respData)
	if !respData.Success() {
		response.RespJson(req, respData.Code, respData.Msg, respData.Data)
	} else {
		token := respData.GetString("token")
		uuid := respData.GetString("uuid")
		var user *entity.User
		err := req.GetParam("user").Struct(&user)
		if err != nil {
			logger.Error(ctx, "GetParam User Error: ", err.Error())
		}
		// 获取 User-Agent
		userAgent := req.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		// 获取浏览器类型
		explorer, _ := ua.Browser()
		// 获取操作系统
		os := ua.OS()
		// 获取客户端IP
		ip := utils.GetClientIp(req)
		// 保存用户在线状态（异步）
		UserOnline().Invoke(ctx, &entity.UserOnline{
			Uuid:     uuid,
			Token:    token,
			Passport: user.Passport,
			Browser:  explorer,
			Os:       os,
			Ip:       ip,
			Time:     gtime.Now(),
		})
		response.SuccExit(req, &v1.LoginRes{Token: token})
	}
}

// 登出验证方法 return true 继续执行，否则结束执行
func (s *sAuth) logoutBefore(req *ghttp.Request) bool {
	ctx := req.GetCtx()
	logger.Debug(ctx, "logoutBefore: ", req)
	// 删除在线用户状态
	var authHeader string
	authHeader = req.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" && parts[1] != "" {
			// 删除在线用户状态操作
			UserOnline().DeleteOnlineByToken(ctx, parts[1])
		}
	}
	authHeader = req.Get("token").String()
	if authHeader != "" {
		// 删除在线用户状态操作
		UserOnline().DeleteOnlineByToken(ctx, authHeader)
	}

	Context().Init(req, nil)
	err := Session().RemoveUser(ctx)
	if err != nil {
		logger.Error(ctx, "Session RemoveUser Error: ", err.Error())
	}
	return true
}

// 登出返回方法
func (s *sAuth) logoutAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "logoutAfter: ", respData)
	response.RespJson(req, respData.Code, respData.Msg, respData.Data)
}

// 认证验证方法 return true 继续执行，否则结束执行
func (s *sAuth) authBefore(req *ghttp.Request) bool {
	ctx := req.GetCtx()
	logger.Debug(ctx, "authBefore: ", req)
	return true
}

// 认证返回方法
func (s *sAuth) authAfter(req *ghttp.Request, respData gtoken.Resp) {
	ctx := req.GetCtx()
	logger.Debug(ctx, "authAfter: ", respData)
	if req.Method == "OPTIONS" || respData.Success() {
		req.Middleware.Next()
	} else if respData.Code == gtoken.UNAUTHORIZED {
		response.RespJsonExit(req, http.StatusUnauthorized, s.token.AuthFailMsg, respData.Data)
	} else {
		response.RespJsonExit(req, respData.Code, respData.Msg, respData.Data)
	}
}
