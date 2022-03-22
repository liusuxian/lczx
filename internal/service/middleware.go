package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"lczx/internal/model"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
)

type sMiddleware struct{}

var insMiddleware = sMiddleware{}

// Middleware 中间件管理服务
func Middleware() *sMiddleware {
	return &insMiddleware
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(req *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: req.Session,
		Data:    make(g.Map),
	}
	Context().Init(req, customCtx)
	ctx := req.GetCtx()
	userInfo := &entity.User{}
	resp := GfToken().GetTokenData(req)
	if resp.Success() {
		err := gconv.Struct(resp.Get("data"), userInfo)
		if err != nil {
			logger.Error(ctx, "Ctx GetUserData Error: ", err.Error())
		}
		if userInfo != nil {
			customCtx.User = &model.ContextUser{
				Id:       userInfo.Id,
				Passport: userInfo.Passport,
				Realname: userInfo.Realname,
				Nickname: userInfo.Nickname,
				Gender:   userInfo.Gender,
				Avatar:   userInfo.Avatar,
				Mobile:   userInfo.Mobile,
				DeptId:   userInfo.DeptId,
				RoleId:   userInfo.RoleId,
			}
		}
	}
	// 执行下一步请求逻辑
	req.Middleware.Next()
}

// CORS 允许接口跨域请求
func (s *sMiddleware) CORS(req *ghttp.Request) {
	req.Response.CORSDefault()
	req.Middleware.Next()
}
