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
			// 清除所有缓存
			service.Cache().ClearAllCache(ctx)
			// 注册参数校验规则
			service.ParamsValid().RegisterRule()
			// 注册并启动所有任务
			service.Crontab().RegisterAndStartAllTask(ctx)

			s := g.Server()
			// 不认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					service.Middleware().HandlerResponse,
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
					service.Middleware().HandlerResponse,
				)
				err = service.Auth().Token().Middleware(ctx, group)
				if err != nil {
					logger.Panic(ctx, "Init GfToken Error: ", err)
				}
				// 权限判断
				group.Middleware(
					service.Middleware().Auth,
				)
				// 后台操作日志记录
				group.Hook("/*", ghttp.HookAfterOutput, service.OperLog().Invoke)
				// 用户相关
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(controller.User)
				})
				// 文件Url
				group.Group("/fileUrl", func(group *ghttp.RouterGroup) {
					group.Bind(controller.FileUrl)
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
					// 用户管理
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
					// 服务监控
					group.Group("/server_monitor", func(group *ghttp.RouterGroup) {
						group.Bind(controller.ServerMonitor)
					})
					// 登录日志
					group.Group("/loginLog", func(group *ghttp.RouterGroup) {
						group.Bind(controller.LoginLog)
					})
					// 操作日志
					group.Group("/operLog", func(group *ghttp.RouterGroup) {
						group.Bind(controller.OperLog)
					})
					// 定时任务管理
					group.Group("/crontab", func(group *ghttp.RouterGroup) {
						group.Bind(controller.Crontab)
					})
				})
				// 文档库
				group.Group("/wdk", func(group *ghttp.RouterGroup) {
					// 项目管理
					group.Group("/project", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkProject)
					})
					// 报告类型配置管理
					group.Group("/reportTypeCfg", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkReportCfg)
					})
					// 上传附件记录管理
					group.Group("/attachment", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkAttachment)
					})
					// 服务记录管理
					group.Group("/service", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkService)
					})
					// 上传文件记录管理
					group.Group("/file", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkFile)
					})
					// 上传报告记录管理
					group.Group("/report", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkReport)
					})
					// 报告审核记录管理
					group.Group("/reportAudit", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkReportAudit)
					})
				})
			})
			// 启动
			s.Run()
			return nil
		},
	}
)
