package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcmd"
	"lczx/internal/controller"
	"lczx/internal/service"
	"lczx/utility/logger"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 数据库缓存适配Redis
			g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(g.Redis()))
			// 自定义参数校验服务
			service.ParamsValid()
			s := g.Server()
			// 不认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				// 验证码
				group.Group("/captcha", func(group *ghttp.RouterGroup) {
					group.Bind(controller.Captcha)
				})
			})
			// 认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				err = service.Auth().Token().Middleware(ctx, group)
				if err != nil {
					logger.Panic(ctx, "Init GfToken Error: ", err.Error())
				}
				// 权限判断
				group.Middleware(
					service.Middleware().Auth,
				)
				// 后台操作日志记录
				group.Hook("/*", ghttp.HookAfterOutput, service.OperLog().Invoke)
				// 缓存相关
				group.Group("/cache", func(group *ghttp.RouterGroup) {
					group.Bind(controller.Cache)
				})
				// TODO 用户相关
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(controller.User)
				})
				// 权限管理
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					// 菜单管理
					group.Group("/menu", func(group *ghttp.RouterGroup) {
						group.Bind(controller.Menu)
					})
					// 角色管理
					group.Group("/role", func(group *ghttp.RouterGroup) {
						group.Bind(controller.Role)
					})
					// 部门管理
					group.Group("/dept", func(group *ghttp.RouterGroup) {
						group.Bind(controller.Dept)
					})
					// TODO 用户管理
					group.Group("/user", func(group *ghttp.RouterGroup) {
						group.Bind(controller.UserManager)
					})
				})
				// 系统监控
				group.Group("/monitor", func(group *ghttp.RouterGroup) {
					// 在线用户管理
					group.Group("/userOnline", func(group *ghttp.RouterGroup) {
						group.Bind(controller.UserOnline)
					})
					// TODO 服务监控
					group.Group("/server", func(group *ghttp.RouterGroup) {
					})
					// 登录日志
					group.Group("/loginLog", func(group *ghttp.RouterGroup) {
						group.Bind(controller.LoginLog)
					})
					// 操作日志
					group.Group("/operLog", func(group *ghttp.RouterGroup) {
						group.Bind(controller.OperLog)
					})
				})
			})
			// 启动
			s.Run()
			return nil
		},
	}
)
