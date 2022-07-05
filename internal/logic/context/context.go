package context

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"lczx/internal/consts"
	"lczx/internal/model"
	"lczx/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sContext struct{}

func init() {
	service.RegisterContext(newContext())
}

// 上下文服务
func newContext() *sContext {
	return &sContext{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(req *ghttp.Request, customCtx *model.Context) {
	req.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获得上下文变量
func (s *sContext) Get(ctx context.Context) (localCtx *model.Context, err error) {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		err = gerror.New("上下文信息不存在")
		return
	}
	err = gconv.Scan(value, &localCtx)
	return
}

// GetUser 获取上下文用户信息
func (s *sContext) GetUser(ctx context.Context) (user *model.ContextUser, err error) {
	var localCtx *model.Context
	localCtx, err = s.Get(ctx)
	if err != nil {
		return
	}
	user = localCtx.User
	if user == nil {
		err = gerror.New("上下文用户信息不存在")
		return
	}
	return
}

// GetSession 获取当前Session管理对象
func (s *sContext) GetSession(ctx context.Context) (session *ghttp.Session, err error) {
	var localCtx *model.Context
	localCtx, err = s.Get(ctx)
	if err != nil {
		return
	}
	session = localCtx.Session
	if session == nil {
		err = gerror.New("当前Session管理对象不存在")
		return
	}
	return
}

// SetUser 将上下文用户信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) (err error) {
	var localCtx *model.Context
	localCtx, err = s.Get(ctx)
	if err != nil {
		return
	}
	localCtx.User = ctxUser
	return
}

// SetData 将自定KV变量设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetData(ctx context.Context, data g.Map) (err error) {
	var localCtx *model.Context
	localCtx, err = s.Get(ctx)
	if err != nil {
		return
	}
	localCtx.Data = data
	return
}
