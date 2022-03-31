package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
				err = service.Auth(ctx).Token().Middleware(ctx, group)
				if err != nil {
					logger.Panic(ctx, "Init GfToken Error: ", err.Error())
				}
				// TODO 用户相关
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(controller.User)
				})
				// TODO 权限管理
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(
						// 菜单管理
						// 角色管理
						// 部门管理
						controller.Dept,
						// 岗位管理
						// 用户管理
					)
				})
				// TODO 系统监控
				group.Group("/monitor", func(group *ghttp.RouterGroup) {
					group.Group("/userOnline", func(group *ghttp.RouterGroup) {
						group.Bind(controller.UserOnline)
					})
					// TODO 定时任务管理
					group.Group("/job", func(group *ghttp.RouterGroup) {
					})
					// TODO 服务监控
					group.Group("/server", func(group *ghttp.RouterGroup) {
					})
					// TODO 登录日志
					group.Group("/loginLog", func(group *ghttp.RouterGroup) {
						group.Bind(controller.LoginLog)
					})
					// TODO 操作日志
					group.Group("/operationLog", func(group *ghttp.RouterGroup) {
					})
				})
			})
			// 启动
			s.Run()
			return nil
		},
	}
)
