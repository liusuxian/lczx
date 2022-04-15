package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/consts"
)

var (
	Cache = cCache{}
)

type cCache struct{}

// Clear 清除缓存
func (c *cCache) Clear(ctx context.Context, req *v1.CacheClearReq) (res *v1.CacheClearRes, err error) {
	err = g.DB().GetCache().Removes(ctx, []interface{}{consts.MenuKey, consts.RoleKey, consts.DeptKey})
	if err != nil {
		err = gerror.WrapCode(code.ClearCacheFailed, err)
		return
	}

	return
}
