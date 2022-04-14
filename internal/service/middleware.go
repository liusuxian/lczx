package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"lczx/internal/code"
	"lczx/internal/model"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
	"lczx/utility/response"
)

type sMiddleware struct{}

var (
	insMiddleware = sMiddleware{}
)

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
	respData := Auth().Token().GetTokenData(req)
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

// Auth 权限判断处理中间件
func (s *sMiddleware) Auth(req *ghttp.Request) {
	ctx := req.GetCtx()
	user := Context().Get(ctx).User
	accessParams := req.Get("accessParams").Strings()
	accessParamsStr := ""
	if len(accessParams) > 0 && accessParams[0] != "undefined" {
		accessParamsStr = "?" + gstr.Join(accessParams, "&")
	}
	// 获取无需验证权限的用户ID
	tagSuperUser := false
	UserManager().NotCheckAuthUserIds().Iterator(func(v interface{}) bool {
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
	menuList, err = Menu().GetStatusEnableMenus(ctx)
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
	logger.Debug(ctx, "menu: ", menu)
	// 只验证存在数据库中的规则
	if menu != nil {
		// 若存在不需要验证的条件则跳过
		if gstr.Equal(menu.Condition, "nocheck") {
			req.Middleware.Next()
			return
		}
		// 菜单没存数据库不验证权限
		if menu.Id != 0 {
			// 判断权限操作
			var enforcer *casbin.SyncedEnforcer
			enforcer, err = Casbin().GetEnforcer()
			logger.Debug(ctx, "enforcer: ", enforcer)
			if err != nil {
				response.RespJsonExitByGcode(req, code.GetAccessAuthFailed)
			}
			groupPolicy := enforcer.GetFilteredGroupingPolicy(0, gconv.String(user.Id))
			logger.Debug(ctx, "groupPolicy: ", groupPolicy)
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
