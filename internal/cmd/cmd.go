package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"lczx/internal/controller"
	"lczx/internal/service"
	"lczx/utility/logger"
)

var gfToken *gtoken.GfToken

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 不认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				// 调试路由
				group.Bind(
					controller.Hello,
				)
			})
			// 认证接口
			// 启动 gtoken
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
				LoginBeforeFunc:  service.LoginBeforeFunc,
				LoginAfterFunc:   service.LoginAfterFunc,
				LogoutPath:       "/logout",
				LogoutBeforeFunc: service.LogoutBeforeFunc,
				LogoutAfterFunc:  service.LogoutAfterFunc,
				AuthPaths:        g.SliceStr{}, // 这里是按照前缀拦截，拦截
				AuthExcludePaths: g.SliceStr{}, // 不拦截路径
				AuthBeforeFunc:   service.AuthBeforeFunc,
				AuthAfterFunc:    service.AuthAfterFunc,
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					logger.Panic(ctx, "GfToken Start Error: ", err.Error())
				}
			})
			// 启动
			s.Run()
			return nil
		},
	}
)
