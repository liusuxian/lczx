package v1

import "github.com/gogf/gf/v2/frame/g"

// DownloadFileReq 下载文件请求参数
type DownloadFileReq struct {
	g.Meta   `path:"/file" tags:"DownloadFile" method:"get" summary:"You first download/file api"`
	FilePath string `json:"filePath" v:"required#文件路径不能为空" dc:"文件路径"` // 文件路径
}

// DownloadFileRes 下载文件返回参数
type DownloadFileRes struct {
}
