package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkAttachmentRecordReq 文档库上传附件记录请求参数
type WdkAttachmentRecordReq struct {
	g.Meta    `path:"/record" tags:"WdkAttachmentRecord" method:"get" summary:"You first wdk/attachment/record api"`
	ProjectId uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
}

// WdkAttachmentRecordRes 文档库上传附件记录返回参数
type WdkAttachmentRecordRes struct {
	List []*WdkAttachmentRecordInfo `json:"list" dc:"文档库上传附件记录信息列表"` // 文档库上传附件记录信息列表
}

// WdkAttachmentRecordInfo 文档库上传附件记录信息
type WdkAttachmentRecordInfo struct {
	Record     *entity.WdkAttachmentRecord `json:"record" dc:"文档库上传附件记录"`   // 文档库上传附件记录
	Attachment []*entity.WdkAttachmentFile `json:"attachment" dc:"文档库上传附件"` // 文档库上传附件
}

// WdkAttachmentRecordAddReq 文档库新增上传附件记录请求参数
type WdkAttachmentRecordAddReq struct {
	g.Meta     `path:"/add" tags:"WdkAttachmentRecordAdd" method:"post" summary:"You first wdk/attachment/add api"`
	UploadName string `json:"uploadName" v:"required#表单文件字段名不能为空" dc:"表单文件字段名"`                             // 表单文件字段名
	ProjectId  uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
	Remark     string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                // 备注
}

// WdkAttachmentRecordAddRes 文档库新增上传附件记录返回参数
type WdkAttachmentRecordAddRes struct {
}

// WdkAttachmentRecordDeleteReq 文档库删除上传附件记录请求参数
type WdkAttachmentRecordDeleteReq struct {
	g.Meta `path:"/delete" tags:"WdkAttachmentRecordDelete" method:"delete" summary:"You first wdk/attachment/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#附件上传记录ID列表不能为空" dc:"附件上传记录ID列表"` // 附件上传记录ID列表
}

// WdkAttachmentRecordDeleteRes 文档库删除上传附件记录返回参数
type WdkAttachmentRecordDeleteRes struct {
}
