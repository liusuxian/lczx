package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"lczx/internal/consts"
	"lczx/utility/logger"
	"time"
)

type sCache struct{}

var (
	insCache = sCache{}
)

func init() {
	// 设置数据库缓存适配Redis
	g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(g.Redis()))
}

// Cache 缓存管理服务
func Cache() *sCache {
	return &insCache
}

// ClearAllCache 清除所有缓存
func (s *sCache) ClearAllCache(ctx context.Context) {
	var err error
	err = g.DB().GetCache().Removes(ctx, []any{consts.MenuKey, consts.RoleKey, consts.DeptKey})
	if err != nil {
		logger.Error(ctx, "ClearAllCache Error: ", err.Error())
	}
	s.scanClearCache(ctx, "GToken:lczx:*")
	s.scanClearCache(ctx, "cache:lczx:accessUrl:*")
}

// ClearCache 清除缓存
func (s *sCache) ClearCache(ctx context.Context, keys ...any) (err error) {
	if len(keys) != 0 {
		err = g.DB().GetCache().Removes(ctx, keys)
	}
	return
}

// GetCache 获取缓存
func (s *sCache) GetCache(ctx context.Context, cacheKey any, mapKey ...string) any {
	var cacheVal *gvar.Var
	var err error
	cacheVal, err = g.DB().GetCache().Get(ctx, cacheKey)
	if err != nil {
		logger.Errorf(ctx, "CacheKey: %s, GetCache Error: %s", cacheKey, err.Error())
		return nil
	}
	if cacheVal != nil {
		if len(mapKey) == 0 {
			return cacheVal.Map()["Result"]
		}
		return cacheVal.Map()[mapKey[0]]
	}
	return nil
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
			logger.Error(ctx, "scanClearCache Error: ", err.Error())
		}
		data := gconv.SliceAny(val)
		delKeys := gconv.Interfaces(data[1])
		// 批量删除
		if len(delKeys) != 0 {
			_, err = g.Redis().Do(ctx, "DEL", delKeys...)
			if err != nil {
				logger.Error(ctx, "scanClearCache Error: ", err.Error())
			}
		}
		cursor = gconv.Int(data[0])
		if cursor == 0 {
			break
		}
	}
}
