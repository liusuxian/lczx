package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkFileRecordReq 文档库上传文件记录请求参数
type WdkFileRecordReq struct {
	g.Meta    `path:"/record" tags:"WdkFileRecord" method:"get" summary:"You first wdk/file/record api"`
	ProjectId uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
}

// WdkFileRecordRes 文档库上传文件记录返回参数
type WdkFileRecordRes struct {
	List []*WdkFileRecordInfo `json:"list" dc:"文档库上传文件记录信息列表"` // 文档库上传文件记录信息列表
}

// WdkFileRecordInfo 文档库上传文件记录信息
type WdkFileRecordInfo struct {
	File     *entity.WdkFile       `json:"file" dc:"文档库上传文件信息"`     // 文档库上传文件信息
	Filetype []*entity.WdkFiletype `json:"filetype" dc:"文档库上传文件类型"` // 文档库上传文件类型
}
