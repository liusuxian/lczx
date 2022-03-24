package controller

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"lczx/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteJson(gtoken.Succ("Hello World!"))
	return
}
