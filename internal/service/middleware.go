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

// Middleware 中间件服务
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
	user := &entity.User{}
	respData := Auth(ctx).Token().GetTokenData(req)
	if respData.Success() {
		err := gconv.Struct(respData.Get("data"), user)
		if err != nil {
			logger.Error(ctx, "Ctx GetUserData Error: ", err.Error())
		}
		if user != nil {
			Context().SetUser(ctx, &model.ContextUser{
				Id:            user.Id,
				Passport:      user.Passport,
				Realname:      user.Realname,
				Nickname:      user.Nickname,
				Gender:        user.Gender,
				Avatar:        user.Avatar,
				Mobile:        user.Mobile,
				DeptId:        user.DeptId,
				Status:        user.Status,
				Email:         user.Email,
				LastLoginIp:   user.LastLoginIp,
				LastLoginTime: user.LastLoginTime,
				CreateAt:      user.CreateAt,
			})
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
