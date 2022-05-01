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
	Project         *entity.WdkProject           `json:"project" dc:"所属项目信息"`            // 所属项目信息
	ReportAudit     *entity.WdkReportAudit       `json:"reportAudit" dc:"文档库报告审核信息"`     // 文档库报告审核信息
	ReportAuditType []*entity.WdkReportAuditType `json:"reportAuditType" dc:"文档库报告审核类型"` // 文档库报告审核类型
}

// WdkReportBeAuditedListReq 文档库报告被审核记录列表请求参数
type WdkReportBeAuditedListReq struct {
	g.Meta   `path:"/beAuditedList" tags:"WdkReportBeAuditedList" method:"get" summary:"You first wdk/reportAudit/beAuditedList api"`
	Status   uint `json:"status" v:"required|in:0,1,2#审核状态不能为空|审核状态只能是0,1,2" dc:"审核状态 0:未通过 1:审核中 2:已通过"` // 审核状态 0:未通过 1:审核中 2:已通过
	CurPage  int  `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`           // 当前页码
	PageSize int  `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`          // 每页数量
}

// WdkReportBeAuditedListRes 文档库报告被审核记录列表返回参数
type WdkReportBeAuditedListRes struct {
	CurPage int                       `json:"curPage" dc:"当前页码"`        // 当前页码
	Total   int                       `json:"total" dc:"数据总量"`          // 数据总量
	List    []*WdkReportBeAuditedInfo `json:"list" dc:"文档库报告被审核记录信息列表"` // 文档库报告被审核记录信息列表
}

// WdkReportBeAuditedInfo 文档库报告被审核记录信息
type WdkReportBeAuditedInfo struct {
	Report          *entity.WdkReport            `json:"report" dc:"文档库报告信息"`              // 文档库报告信息
	Project         *entity.WdkProject           `json:"project" dc:"所属项目信息"`              // 所属项目信息
	ReportAudit     []*entity.WdkReportAudit     `json:"reportAudit" dc:"文档库报告审核信息列表"`     // 文档库报告审核信息列表
	ReportAuditType []*entity.WdkReportAuditType `json:"reportAuditType" dc:"文档库报告审核类型列表"` // 文档库报告审核类型列表
}
