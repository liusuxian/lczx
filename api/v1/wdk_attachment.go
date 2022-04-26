package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkAttachmentListReq 文档库上传附件列表请求参数
type WdkAttachmentListReq struct {
	g.Meta `path:"/list" tags:"WdkAttachmentList" method:"get" summary:"You first wdk/filetype/list api"`
}

// WdkAttachmentListRes 文档库上传附件列表返回参数
type WdkAttachmentListRes struct {
	List []*WdkFiletypeInfo `json:"list" dc:"文档库上传文件类型信息列表"` // 文档库上传文件类型信息列表
}

// WdkFiletypeInfo 文档库上传文件类型信息
type WdkFiletypeInfo struct {
	FiletypeCfg *entity.WdkAttachmentRecord `json:"filetypeCfg" dc:"文档库上传文件类型配置信息"` // 文档库上传文件类型配置信息
	AuditCfg    []*entity.WdkAuditCfg       `json:"auditCfg" dc:"文档库上传文件类型审核配置信息"`  // 文档库上传文件类型审核配置信息
}
