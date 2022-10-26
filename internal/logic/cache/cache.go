package cache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"lczx/internal/service"
	"lczx/utility/logger"
	"time"
)

type sCache struct{}

func init() {
	service.RegisterCache(newCache())
}

// 缓存管理服务
func newCache() *sCache {
	// 设置数据库缓存适配Redis
	g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(g.Redis()))
	return &sCache{}
}

// ClearAllCache 清除所有缓存
func (s *sCache) ClearAllCache(ctx context.Context) {
	keys, err := g.DB().GetCache().Keys(ctx)
	if err != nil {
		logger.Error(ctx, "ClearAllCache Get All Keys Error: ", err)
	}
	err = g.DB().GetCache().Removes(ctx, keys)
	if err != nil {
		logger.Error(ctx, "ClearAllCache Removes All Keys Error: ", err)
	}
}

// ClearCache 清除缓存
func (s *sCache) ClearCache(ctx context.Context, keys ...any) (err error) {
	if len(keys) != 0 {
		err = g.DB().GetCache().Removes(ctx, keys)
	}
	return
}

// GetCache 获取缓存
func (s *sCache) GetCache(ctx context.Context, cacheKey any) (cacheVal *gvar.Var) {
	var err error
	cacheVal, err = g.DB().GetCache().Get(ctx, cacheKey)
	if err != nil {
		logger.Errorf(ctx, "CacheKey: %s, GetCache Error: %s", cacheKey, err)
		return nil
	}
	return
}

// SetCache 设置缓存
func (s *sCache) SetCache(ctx context.Context, key, value any, duration time.Duration) (err error) {
	err = g.DB().GetCache().Set(ctx, key, value, duration)
	return
}
