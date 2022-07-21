package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	Download = cDownload{}
)

type cDownload struct{}

// File 下载文件
func (c *cDownload) File(ctx context.Context, req *v1.DownloadFileReq) (res *v1.DownloadFileRes, err error) {
	var body io.ReadCloser
	body, err = service.AliyunOSS().Download(ctx, req.FilePath)
	if err != nil {
		err = gerror.WrapCode(code.DownloadFileFailed, err)
		return
	}

	defer body.Close()
	_, err = io.Copy(g.RequestFromCtx(ctx).Response.Writer, body)
	if err != nil {
		err = gerror.WrapCode(code.DownloadFileFailed, err)
		return
	}

	return
}
