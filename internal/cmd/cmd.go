package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	v1 "lczx/api/v1"
	"lczx/internal/controller"
	"lczx/internal/service"
	"lczx/internal/upload"
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
			// 自定义参数校验服务
			service.ParamsValid()
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
				// 测试上传文件
				group.Group("/upload", func(group *ghttp.RouterGroup) {
					// 测试上传头像
					group.POST("/avatar", func(req *ghttp.Request) {
						// 获取上传文件信息
						avatar := req.GetUploadFile("upload-avatar")
						if avatar != nil {
							f, e := upload.Upload.UploadImg(avatar, "user/avatar")
							if e != nil {
								fmt.Println("upload avatar err1: ", e)
							} else {
								fmt.Println("upload avatar res: ", f)
								// 设置用户头像
								e = service.User().SetAvatar(ctx, 1, f.OriginFileUrl)
								if e != nil {
									fmt.Println("upload avatar err2: ", e)
								}
							}
						}
					})
					// 测试上传附件
					group.POST("/attachment", func(req *ghttp.Request) {
						// 获取上传文件信息
						files := req.GetUploadFiles("upload-attachment")
						if len(files) > 0 && len(files) <= 4 {
							fs, e := upload.Upload.UploadFiles(files, "wdk/attachment")
							if e != nil {
								fmt.Println("upload attachment err1: ", e)
							} else {
								fmt.Println("upload attachment res: ", fs)
								// 新增文档库上传附件记录
								e = service.WdkAttachment().AddWdkAttachmentRecord(ctx, &v1.WdkAttachmentRecordAddReq{
									ProjectId: 1,
								}, fs)
								if e != nil {
									fmt.Println("upload attachment err2: ", e)
								}
							}
						}
					})
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
					logger.Panic(ctx, "Init GfToken Error: ", err.Error())
				}
				// 权限判断
				group.Middleware(
					service.Middleware().Auth,
				)
				// 后台操作日志记录
				group.Hook("/*", ghttp.HookAfterOutput, service.OperLog().Invoke)
				// 用户相关
				group.Group("/user", func(group *ghttp.RouterGroup) {
					// 上传头像未完成
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
				})
				// 文档库
				group.Group("/wdk", func(group *ghttp.RouterGroup) {
					// 项目管理
					group.Group("/project", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkProject)
					})
					// 上传文件类型管理
					group.Group("/filetype", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkFiletype)
					})
					// 上传附件记录管理
					group.Group("/attachment", func(group *ghttp.RouterGroup) {
						group.Bind(controller.WdkAttachment)
					})
				})
			})
			// 每2小时执行一次检查在线用户
			_, err = gcron.Add(ctx, "0 0 */2 * * *", service.Auth().CheckUserOnline)
			if err != nil {
				return err
			}
			// 启动
			s.Run()
			return nil
		},
	}
)
