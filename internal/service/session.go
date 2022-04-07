package service

import (
	"context"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
)

type sSession struct{}

var (
	insSession = sSession{}
)

// Session session服务
func Session() *sSession {
	return &insSession
}

// SetUser 设置用户信息到session中
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return Context().Get(ctx).Session.Set(consts.SessionKey, user)
}

// GetUser 从session中检索并返回用户，如果用户没有登录，则返回nil
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := Context().Get(ctx)
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
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.SessionKey)
	}
	return nil
}
