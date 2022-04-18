package v1

import "github.com/gogf/gf/v2/frame/g"

// CacheClearReq 清除缓存请求参数
type CacheClearReq struct {
	g.Meta `path:"/clear" tags:"CacheClear" method:"delete" summary:"You first cache/clear api"`
}

// CacheClearRes 清除缓存返回参数
type CacheClearRes struct {
}
