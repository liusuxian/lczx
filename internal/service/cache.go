// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

type ICache interface {
	ClearAllCache(ctx context.Context)
	ClearCache(ctx context.Context, keys ...any) (err error)
	GetCache(ctx context.Context, cacheKey any) (cacheVal *gvar.Var)
	SetCache(ctx context.Context, key, value any, duration time.Duration) (err error)
}

var localCache ICache

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
