package cache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"lczx/internal/consts"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/utils"
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
	var err error
	err = g.DB().GetCache().Removes(ctx, []any{consts.MenuKey, consts.RoleKey, consts.DeptKey})
	if err != nil {
		logger.Error(ctx, "ClearAllCache Error: ", err)
	}
	s.scanClearCache(ctx, utils.RedisKey("GToken:lczx:", "*"))
	s.scanClearCache(ctx, utils.RedisKey(consts.CacheAccessUrlPrefix, "*"))
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

// 扫描删除缓存
func (s *sCache) scanClearCache(ctx context.Context, cacheKey string) {
	var err error
	cursor := 0
	for {
		var val *gvar.Var
		val, err = g.Redis().Do(ctx, "scan", cursor, "match", cacheKey, "count", "100")
		if err != nil {
			logger.Error(ctx, "scanClearCache Error: ", err)
		}
		data := gconv.SliceAny(val)
		delKeys := gconv.Interfaces(data[1])
		// 批量删除
		if len(delKeys) != 0 {
			_, err = g.Redis().Do(ctx, "DEL", delKeys...)
			if err != nil {
				logger.Error(ctx, "scanClearCache Error: ", err)
			}
		}
		cursor = gconv.Int(data[0])
		if cursor == 0 {
			break
		}
	}
}
