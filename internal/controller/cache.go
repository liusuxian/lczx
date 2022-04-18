package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	Cache = cCache{}
)

type cCache struct{}

// Clear 清除缓存
func (c *cCache) Clear(ctx context.Context, req *v1.CacheClearReq) (res *v1.CacheClearRes, err error) {
	err = service.Cache().ClearAllCache(ctx)
	if err != nil {
		err = gerror.WrapCode(code.ClearCacheFailed, err)
		return
	}

	return
}
