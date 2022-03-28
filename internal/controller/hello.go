package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"lczx/api/v1"
	"lczx/utility/response"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	response.SuccExit(g.RequestFromCtx(ctx), "Hello World!")
	return
}
