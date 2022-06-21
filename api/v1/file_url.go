package v1

import "github.com/gogf/gf/v2/frame/g"

// FileUrlGetReq 获取文件url请求参数
type FileUrlGetReq struct {
	g.Meta   `path:"/get" tags:"FileUrlGet" method:"get" summary:"You first fileUrl/get api"`
	FilePath string `json:"filePath" v:"required#文件路径不能为空" dc:"文件路径"` // 文件路径
}

// FileUrlGetRes 获取文件url返回参数
type FileUrlGetRes struct {
	FileUrl string `json:"fileUrl" dc:"文件url地址"` // 文件url地址
}
