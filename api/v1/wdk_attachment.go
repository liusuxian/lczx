package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"lczx/internal/model/entity"
)

// WdkAttachmentRecordReq 文档库上传附件记录请求参数
type WdkAttachmentRecordReq struct {
	g.Meta    `path:"/record" tags:"WdkAttachmentRecord" method:"get" summary:"You first wdk/attachment/record api"`
	ProjectId uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
}

// WdkAttachmentRecordRes 文档库上传附件记录返回参数
type WdkAttachmentRecordRes struct {
	List []*WdkAttachmentInfo `json:"list" dc:"文档库上传附件信息列表"` // 文档库上传附件信息列表
}

// WdkAttachmentInfo 文档库上传附件信息
type WdkAttachmentInfo struct {
	Record     *entity.WdkAttachmentRecord `json:"record" dc:"文档库上传附件记录"`   // 文档库上传附件记录
	Attachment []*entity.WdkAttachmentFile `json:"attachment" dc:"文档库上传附件"` // 文档库上传附件
}

// WdkAttachmentAddReq 文档库新增上传附件记录请求参数
type WdkAttachmentAddReq struct {
	g.Meta            `path:"/add" tags:"WdkAttachmentAdd" method:"post" summary:"You first wdk/attachment/add api"`
	ProjectId         uint64              `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
	Remark            string              `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                // 备注
	UploadAttachment  *ghttp.UploadFile   `json:"uploadAttachment" dc:"上传附件"`                                                   // 上传附件
	UploadAttachments []*ghttp.UploadFile `json:"uploadAttachments" dc:"上传附件列表"`                                                // 上传附件列表
}

// WdkAttachmentAddRes 文档库新增上传附件记录返回参数
type WdkAttachmentAddRes struct {
}
