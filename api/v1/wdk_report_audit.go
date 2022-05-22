package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkReportAuditListReq 文档库报告审核记录列表请求参数
type WdkReportAuditListReq struct {
	g.Meta   `path:"/list" tags:"WdkReportAuditList" method:"get" summary:"You first wdk/reportAudit/list api"`
	Status   uint `json:"status" v:"required|in:0,1,2#审核状态不能为空|审核状态只能是0,1,2" dc:"审核状态 0:未通过 1:审核中 2:已通过"` // 审核状态 0:未通过 1:审核中 2:已通过
	CurPage  int  `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`           // 当前页码
	PageSize int  `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`          // 每页数量
}

// WdkReportAuditListRes 文档库报告审核记录列表返回参数
type WdkReportAuditListRes struct {
	CurPage int                   `json:"curPage" dc:"当前页码"`       // 当前页码
	Total   int                   `json:"total" dc:"数据总量"`         // 数据总量
	List    []*WdkReportAuditInfo `json:"list" dc:"文档库报告审核记录信息列表"` // 文档库报告审核记录信息列表
}

// WdkReportAuditInfo 文档库报告审核记录信息
type WdkReportAuditInfo struct {
	Report          *entity.WdkReport            `json:"report" dc:"文档库报告信息"`            // 文档库报告信息
	ReportAudit     *entity.WdkReportAudit       `json:"reportAudit" dc:"文档库报告审核信息"`     // 文档库报告审核信息
	ReportAuditType []*entity.WdkReportAuditType `json:"reportAuditType" dc:"文档库报告审核类型"` // 文档库报告审核类型
}

// WdkReportAuditReq 文档库报告审核请求参数
type WdkReportAuditReq struct {
	g.Meta      `path:"/audit" tags:"WdkReportAudit" method:"put" summary:"You first wdk/reportAudit/audit api"`
	Id          uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#报告ID不能为空|报告ID必须为正整数" dc:"报告ID"`                                        // 报告ID
	AuditStatus uint   `json:"auditStatus" v:"required|in:0,2#审核状态不能为空|审核状态只能是0,2" dc:"审核状态 0:未通过 2:已通过"`                              // 审核状态 0:未通过 2:已通过
	Excellence  uint   `json:"excellence" v:"required|in:0,1#是否被推荐为优秀报告不能为空|是否被推荐为优秀报告只能是0,1" dc:"是否被推荐为优秀报告 0:未被推荐为优秀报告 1:已被推荐为优秀报告"` // 是否被推荐为优秀报告 0:未被推荐为优秀报告 1:已被推荐为优秀报告
}

// WdkReportAuditRes 文档库报告审核返回参数
type WdkReportAuditRes struct {
}

// WdkReportUploadAuditListReq 文档库报告上传审核列表请求参数
type WdkReportUploadAuditListReq struct {
	g.Meta   `path:"/uploadAuditList" tags:"WdkReportUploadAuditList" method:"get" summary:"You first wdk/reportAudit/uploadAuditList api"`
	Status   uint `json:"status" v:"required|in:0,1,2#审核状态不能为空|审核状态只能是0,1,2" dc:"审核状态 0:未通过 1:审核中 2:已通过"`   // 审核状态 0:未通过 1:审核中 2:已通过
	Rescind  uint `json:"rescind" v:"required-if:Status,1|in:0,1#是否已撤销不能为空|是否已撤销只能是0,1" dc:"是否已撤销 0:否 1:是"` // 是否已撤销 0:否 1:是
	CurPage  int  `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`             // 当前页码
	PageSize int  `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`            // 每页数量
}

// WdkReportUploadAuditListRes 文档库报告上传审核列表返回参数
type WdkReportUploadAuditListRes struct {
	CurPage int                         `json:"curPage" dc:"当前页码"`       // 当前页码
	Total   int                         `json:"total" dc:"数据总量"`         // 数据总量
	List    []*WdkReportUploadAuditInfo `json:"list" dc:"文档库报告上传审核信息列表"` // 文档库报告上传审核信息列表
}

// WdkReportUploadAuditInfo 文档库报告上传审核信息
type WdkReportUploadAuditInfo struct {
	Report     *entity.WdkReport       `json:"report" dc:"文档库报告信息"`     // 文档库报告信息
	ReportType []*entity.WdkReportType `json:"reportType" dc:"文档库报告类型"` // 文档库报告类型
}

// WdkReportRescindAuditReq 文档库报告撤销审核请求参数
type WdkReportRescindAuditReq struct {
	g.Meta `path:"/rescindAudit" tags:"WdkReportRescindAudit" method:"put" summary:"You first wdk/reportAudit/rescindAudit api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#报告ID不能为空|报告ID必须为正整数" dc:"报告ID"` // 部门ID
}

// WdkReportRescindAuditRes 文档库报告撤销审核返回参数
type WdkReportRescindAuditRes struct {
}

// WdkReportAuditProcessReq 文档库报告审核流程请求参数
type WdkReportAuditProcessReq struct {
	g.Meta `path:"/auditProcess" tags:"WdkReportAuditProcess" method:"get" summary:"You first wdk/reportAudit/auditProcess api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#报告ID不能为空|报告ID必须为正整数" dc:"报告ID"` // 部门ID
}

// WdkReportAuditProcessRes 文档库报告审核流程返回参数
type WdkReportAuditProcessRes struct {
	List []*WdkReportAuditProcessInfo `json:"list" dc:"文档库报告审核流程信息列表"` // 文档库报告审核流程信息列表
}

// WdkReportAuditProcessInfo 文档库报告审核流程信息
type WdkReportAuditProcessInfo struct {
	ReportAudit     *entity.WdkReportAudit       `json:"reportAudit" dc:"文档库报告审核信息"`       // 文档库报告审核信息
	ReportAuditType []*entity.WdkReportAuditType `json:"reportAuditType" dc:"文档库报告审核类型列表"` // 文档库报告审核类型列表
}
