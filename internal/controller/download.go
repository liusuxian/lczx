package controller

import (
	"context"
	v1 "lczx/api/v1"
)

var (
	Download = cDownload{}
)

type cDownload struct{}

// WdkReportFile 下载文档库报告文件
func (c *cDownload) WdkReportFile(ctx context.Context, req *v1.DownloadWdkReportFileReq) (res *v1.DownloadWdkReportFileRes, err error) {
	return
}
