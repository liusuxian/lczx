package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkReportRecordReq 文档库上传报告记录请求参数
type WdkReportRecordReq struct {
	g.Meta    `path:"/record" tags:"WdkReportRecord" method:"get" summary:"You first wdk/report/record api"`
	ProjectId uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
}

// WdkReportRecordRes 文档库上传报告记录返回参数
type WdkReportRecordRes struct {
	List []*WdkReportInfo `json:"list" dc:"文档库上传报告信息列表"` // 文档库上传报告信息列表
}

// WdkReportInfo 文档库上传报告信息
type WdkReportInfo struct {
	Report     *entity.WdkReport       `json:"report" dc:"文档库上传报告记录"`     // 文档库上传报告记录
	ReportType []*entity.WdkReportType `json:"reportType" dc:"文档库上传报告类型"` // 文档库上传报告类型
}
