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

// WdkReportFile 下载文档库报告文件
func (c *cDownload) WdkReportFile(ctx context.Context, req *v1.DownloadWdkReportFileReq) (res *v1.DownloadWdkReportFileRes, err error) {
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
