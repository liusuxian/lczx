package session

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

type sSession struct{}

func init() {
	service.RegisterSession(newSession())
}

// session服务
func newSession() *sSession {
	return &sSession{}
}

// SetUser 设置用户信息到session中
func (s *sSession) SetUser(ctx context.Context, user *entity.User) (err error) {
	var session *ghttp.Session
	session, err = service.Context().GetSession(ctx)
	if err != nil {
		return
	}
	err = session.Set(consts.SessionKey, user)
	return
}

// GetUser 从session中检索并返回用户
func (s *sSession) GetUser(ctx context.Context) (user *entity.User, err error) {
	var session *ghttp.Session
	session, err = service.Context().GetSession(ctx)
	if err != nil {
		return
	}
	sessionVal := session.MustGet(consts.SessionKey)
	if sessionVal.IsNil() {
		err = gerror.New("从session中获取用户信息失败")
		return
	}
	err = sessionVal.Struct(&user)
	return
}

// RemoveUser 从session中移除用户信息
func (s *sSession) RemoveUser(ctx context.Context) (err error) {
	var session *ghttp.Session
	session, err = service.Context().GetSession(ctx)
	if err != nil {
		return
	}
	err = session.Remove(consts.SessionKey)
	return
}
