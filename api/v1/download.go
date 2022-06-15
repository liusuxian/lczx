package v1

import "github.com/gogf/gf/v2/frame/g"

// DownloadWdkProjectFileReq 下载文档库项目信息文件请求参数
type DownloadWdkProjectFileReq struct {
	g.Meta   `path:"/wdkProjectFile" tags:"DownloadWdkProjectFile" method:"get" summary:"You first download/wdkProjectFile api"`
	FilePath string `json:"filePath" v:"required#文件在服务器的路径不能为空" dc:"文件在服务器的路径"` // 文件在服务器的路径
}

// DownloadWdkProjectFileRes 下载文档库项目信息文件返回参数
type DownloadWdkProjectFileRes struct {
}
