package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	FileUrl = cFileUrl{}
)

type cFileUrl struct{}

// GetFileUrl 获取文件url
func (c *cFileUrl) GetFileUrl(ctx context.Context, req *v1.FileUrlGetReq) (res *v1.FileUrlGetRes, err error) {
	var fileUrl string
	fileUrl, err = service.AliyunOSS().GetAccessUrl(ctx, req.FilePath)
	if err != nil {
		err = gerror.WrapCode(code.GetFileUrlFailed, err)
		return
	}

	res = &v1.FileUrlGetRes{FileUrl: fileUrl}
	return
}
