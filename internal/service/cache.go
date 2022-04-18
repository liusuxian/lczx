package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/consts"
)

type sCache struct{}

var (
	insCache = sCache{}
)

// Cache 缓存管理服务
func Cache() *sCache {
	return &insCache
}

// ClearAllCache 清除所有缓存
func (s *sCache) ClearAllCache(ctx context.Context) (err error) {
	err = g.DB().GetCache().Removes(ctx, []interface{}{consts.MenuKey, consts.RoleKey, consts.DeptKey})
	return
}

// ClearCache 清除缓存
func (s *sCache) ClearCache(ctx context.Context, key interface{}) (lastVal *gvar.Var, err error) {
	lastVal, err = g.DB().GetCache().Remove(ctx, key)
	return
}

// GetCache 获取缓存
func (s *sCache) GetCache(ctx context.Context, key interface{}) (val *gvar.Var, err error) {
	val, err = g.DB().GetCache().Get(ctx, key)
	return
}
