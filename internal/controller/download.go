package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
)

var (
	Download = cDownload{}
)

type cDownload struct{}

// WdkProjectFile 下载文档库项目信息文件
func (c *cDownload) WdkProjectFile(ctx context.Context, req *v1.DownloadWdkProjectFileReq) (res *v1.DownloadWdkProjectFileRes, err error) {
	g.RequestFromCtx(ctx).Response.ServeFileDownload(req.FilePath)
	if gfile.Exists(req.FilePath) {
		_ = gfile.Remove(req.FilePath)
	}
	return
}
