package session

import (
	"context"
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
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return service.Context().Get(ctx).Session.Set(consts.SessionKey, user)
}

// GetUser 从session中检索并返回用户，如果用户没有登录，则返回nil
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := service.Context().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.SessionKey); !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// RemoveUser 从session中移除用户信息
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.SessionKey)
	}
	return nil
}
