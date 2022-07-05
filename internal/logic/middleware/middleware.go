package middleware

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"lczx/internal/code"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/response"
	"net/http"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(newMiddleware())
}

// 中间件服务
func newMiddleware() *sMiddleware {
	return &sMiddleware{}
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(req *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: req.Session,
		Data:    make(g.Map),
	}
	service.Context().Init(req, customCtx)
	ctx := req.GetCtx()
	user := &entity.User{}
	respData := service.Auth().Token().GetTokenData(req)
	if respData.Success() {
		err := gconv.Struct(respData.Get("data"), user)
		if err != nil {
			logger.Error(ctx, "Ctx GetUserData Error: ", err)
		}
		if user != nil {
			service.Context().SetUser(ctx, &model.ContextUser{
				Id:       user.Id,
				Passport: user.Passport,
				Realname: user.Realname,
				DeptId:   user.DeptId,
				IsAdmin:  user.IsAdmin,
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

// HandlerResponse 自定义返回中间件
func (s *sMiddleware) HandlerResponse(req *ghttp.Request) {
	req.Middleware.Next()
	// 如果有自定义缓冲区内容，则退出当前处理程序。
	if req.Response.BufferLength() > 0 {
		return
	}

	msg := "OK"
	err := req.GetError()
	res := req.GetHandlerResponse()
	rCode := gerror.Code(err)
	if err != nil {
		if rCode == gcode.CodeNil {
			rCode = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if req.Response.Status > 0 && req.Response.Status != http.StatusOK {
		msg = http.StatusText(req.Response.Status)
		switch req.Response.Status {
		case http.StatusNotFound:
			rCode = gcode.CodeNotFound
		case http.StatusForbidden:
			rCode = gcode.CodeNotAuthorized
		default:
			rCode = gcode.CodeUnknown
		}
	} else {
		rCode = gcode.CodeOK
	}

	resData := response.Resp{
		Code: rCode.Code(),
		Msg:  msg,
		Data: res,
	}
	req.SetParam("apiReturnRes", resData)
	req.Response.WriteJson(resData)
}

// Auth 权限判断处理中间件
func (s *sMiddleware) Auth(req *ghttp.Request) {
	ctx := req.GetCtx()
	user := service.Context().Get(ctx).User
	accessParams := req.Get("accessParams").Strings()
	accessParamsStr := ""
	if len(accessParams) > 0 && accessParams[0] != "undefined" {
		accessParamsStr = "?" + gstr.Join(accessParams, "&")
	}
	// 获取无需验证权限的用户ID
	tagSuperUser := false
	service.UserManager().NotCheckAuthUserIds().Iterator(func(v any) bool {
		if gconv.Uint64(v) == user.Id {
			tagSuperUser = true
			return false
		}
		return true
	})
	if tagSuperUser {
		req.Middleware.Next()
		// 不要再往后面执行
		return
	}
	url := gstr.TrimLeft(req.Request.URL.Path, "/") + accessParamsStr
	// 获取地址对应的菜单ID
	var err error
	var menuList []*entity.Menu
	menuList, err = service.Menu().GetAllMenus(ctx)
	if err != nil {
		response.RespJsonExitByGcode(req, code.RequestDataFailed)
	}
	var menu *entity.Menu
	for _, m := range menuList {
		ms := gstr.SubStr(m.Rule, 0, gstr.Pos(m.Rule, "?"))
		if m.Rule == url || ms == url {
			menu = m
			break
		}
	}
	// 只验证存在数据库中的规则
	if menu != nil {
		// 检查菜单状态是否为已停用
		if menu.Status == 0 {
			response.RespJsonExitByGcode(req, code.MenuStatusDisabled)
		}
		// 若存在不需要验证的条件则跳过
		if gstr.Equal(menu.Condition, "nocheck") {
			req.Middleware.Next()
			return
		}
		if gstr.Equal(menu.Condition, "check_dept") {
			s.CheckProjectUpload(ctx, req)
		}
		// 菜单没存数据库不验证权限
		if menu.Id != 0 {
			// 判断权限操作
			var enforcer *casbin.SyncedEnforcer
			enforcer, err = service.Casbin().GetEnforcer(ctx)
			if err != nil {
				response.RespJsonExitByGcode(req, code.GetAccessAuthFailed)
			}
			groupPolicy := enforcer.GetFilteredGroupingPolicy(0, gconv.String(user.Id))
			if len(groupPolicy) == 0 {
				response.RespJsonExitByGcode(req, code.NotAccessAuth)
			}
			hasAccess := false
			for _, v := range groupPolicy {
				if enforcer.HasPolicy(v[1], gconv.String(menu.Id), "All") {
					hasAccess = true
					break
				}
			}
			if !hasAccess {
				response.RespJsonExitByGcode(req, code.NotAccessAuth)
			}
		}
	} else if menu == nil && accessParamsStr != "" {
		response.RespJsonExitByGcode(req, code.NotAccessAuth)
	}
	req.Middleware.Next()
}

// CheckProjectUpload 检查项目相关文件上传权限
func (s *sMiddleware) CheckProjectUpload(ctx context.Context, req *ghttp.Request) {
	// 解析参数
	paramsMap := req.GetMap()
	projectIdStr := paramsMap["projectId"]
	projectId := gconv.Uint64(projectIdStr)
	// 通过文档库项目ID判断文档库项目信息是否存在
	var wdkProject *entity.WdkProject
	var err error
	err = dao.WdkProject.Ctx(ctx).Where(do.WdkProject{Id: projectId}).Scan(&wdkProject)
	if err != nil {
		response.RespJsonExit(req, code.GetWdkProjectFailed.Code(), code.GetWdkProjectFailed.Message()+": "+err.Error())
	}
	if wdkProject == nil {
		err = gerror.Newf(`文档库项目ID[%d]不存在`, projectId)
		response.RespJsonExit(req, code.GetWdkProjectFailed.Code(), code.GetWdkProjectFailed.Message()+": "+err.Error())
	}
	// 通过项目负责人，查询项目负责人是否是部门负责人
	var dept *entity.Dept
	err = dao.Dept.Ctx(ctx).Where(do.Dept{PrincipalUid: wdkProject.PrincipalUid}).Scan(&dept)
	if err != nil {
		response.RespJsonExit(req, code.GetDeptFailed.Code(), code.GetDeptFailed.Message()+": "+err.Error())
	}
	// 判断权限
	user := service.Context().Get(ctx).User
	if !(user.IsAdmin == 1 || user.DeptId == wdkProject.DeptId || (dept != nil && dept.Id == user.DeptId)) {
		response.RespJsonExitByGcode(req, code.NotAccessAuth)
	}
}
