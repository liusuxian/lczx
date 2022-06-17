package v1

import "github.com/gogf/gf/v2/frame/g"

// FileContentGetReq 获取文件内容请求参数
type FileContentGetReq struct {
	g.Meta  `path:"/get" tags:"FileContentGet" method:"get" summary:"You first fileContent/get api"`
	FileUrl string `json:"fileUrl" v:"required|url#文件url地址不能为空|文件url地址无效" dc:"文件url地址"` // 文件url地址
}

// FileContentGetRes 获取文件内容返回参数
type FileContentGetRes struct {
}
