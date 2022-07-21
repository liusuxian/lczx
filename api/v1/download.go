package v1

import "github.com/gogf/gf/v2/frame/g"

// DownloadWdkReportFileReq 下载文档库报告文件请求参数
type DownloadWdkReportFileReq struct {
	g.Meta   `path:"/wdkReportFile" tags:"DownloadWdkReportFile" method:"get" summary:"You first download/wdkReportFile api"`
	FilePath string `json:"filePath" v:"required#文件路径不能为空" dc:"文件路径"` // 文件路径
}

// DownloadWdkReportFileRes 下载文档库报告文件返回参数
type DownloadWdkReportFileRes struct {
}
