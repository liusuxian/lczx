package controller

import (
	"context"
	"lczx/utility/response"

	"lczx/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	response.JsonOK(ctx, "Hello World!")
	return
}
