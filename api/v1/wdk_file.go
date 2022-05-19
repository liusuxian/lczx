package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"lczx/internal/model/entity"
)

// WdkFileRecordReq 文档库上传文件记录请求参数
type WdkFileRecordReq struct {
	g.Meta    `path:"/record" tags:"WdkFileRecord" method:"get" summary:"You first wdk/file/record api"`
	ProjectId uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
}

// WdkFileRecordRes 文档库上传文件记录返回参数
type WdkFileRecordRes struct {
	List []*entity.WdkFile `json:"list" dc:"文档库上传文件记录列表"` // 文档库上传文件记录列表
}

// WdkFileAddReq 文档库新增上传文件记录请求参数
type WdkFileAddReq struct {
	g.Meta      `path:"/add" tags:"WdkFileAdd" method:"post" summary:"You first wdk/file/add api"`
	ProjectId   uint64              `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"`                                                        // 所属项目ID
	Type        uint                `json:"type" v:"required|in:0,1,2,3,4,5,6#文件类型不能为空|文件类型只能是0,1,2,3,4,5,6" dc:"文件类型 0:合同扫描文件 1:年度服务计划书 2:总结报告 3:项目移交 4:复盘报告 5:文件签收单 6:满意度调查表"` // 文件类型 0:合同扫描文件 1:年度服务计划书 2:总结报告 3:项目移交 4:复盘报告 5:文件签收单 6:满意度调查表
	UploadFile  *ghttp.UploadFile   `json:"uploadFile" dc:"上传文件"`                                                                                                                // 上传文件
	UploadFiles []*ghttp.UploadFile `json:"uploadFiles" dc:"上传文件列表"`                                                                                                             // 上传文件列表
}

// WdkFileAddRes 文档库新增上传文件记录返回参数
type WdkFileAddRes struct {
}
