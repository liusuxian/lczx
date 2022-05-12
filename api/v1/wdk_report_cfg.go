package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkReportCfgListReq 文档库报告类型配置信息列表请求参数
type WdkReportCfgListReq struct {
	g.Meta `path:"/list" tags:"WdkReportCfgList" method:"get" summary:"You first wdk/reportTypeCfg/list api"`
	Name   string `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#报告类型名称必须为中文且长度不能超过20" dc:"报告类型名称"` // 报告类型名称
}

// WdkReportCfgListRes 文档库报告类型配置信息列表返回参数
type WdkReportCfgListRes struct {
	List []*WdkReportCfgInfo `json:"list" dc:"文档库报告类型配置信息列表"` // 文档库报告类型配置信息列表
}

// WdkReportCfgInfo 文档库报告类型配置信息
type WdkReportCfgInfo struct {
	ReportCfg      *entity.WdkReportCfg        `json:"reportCfg" dc:"文档库报告类型配置信息"`        // 文档库报告类型配置信息
	ReportAuditCfg []*entity.WdkReportAuditCfg `json:"reportAuditCfg" dc:"文档库报告类型审核配置信息"` // 文档库报告类型审核配置信息
}

// WdkAllReportCfgReq 文档库全部报告类型配置信息列表请求参数
type WdkAllReportCfgReq struct {
	g.Meta `path:"/all" tags:"WdkAllReportCfg" method:"get" summary:"You first wdk/reportTypeCfg/all api"`
}

// WdkAllReportCfgRes 文档库全部报告类型配置信息列表返回参数
type WdkAllReportCfgRes struct {
	List []*entity.WdkReportCfg `json:"list" dc:"文档库报告类型配置信息列表"` // 文档库报告类型配置信息列表
}

// WdkReportCfgAddReq 文档库新增报告类型配置请求参数
type WdkReportCfgAddReq struct {
	g.Meta    `path:"/add" tags:"WdkReportCfgAdd" method:"post" summary:"You first wdk/reportTypeCfg/add api"`
	Name      string   `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#报告类型名称不能为空|报告类型名称必须为中文且长度不能超过20" dc:"报告类型名称"` // 报告类型名称
	AuditUids []uint64 `json:"auditUids" v:"required|slice_valid:uint64#文档库报告审核员用户ID列表不能为空" dc:"文档库报告审核员用户ID列表"`            // 文档库报告审核员用户ID列表
}

// WdkReportCfgAddRes 文档库新增报告类型配置返回参数
type WdkReportCfgAddRes struct {
}

// WdkReportCfgInfoReq 文档库报告类型配置信息请求参数
type WdkReportCfgInfoReq struct {
	g.Meta `path:"/info" tags:"WdkReportCfgInfo" method:"get" summary:"You first wdk/reportTypeCfg/info api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#报告类型ID不能为空|报告类型ID必须为正整数" dc:"报告类型ID"` // 报告类型ID
}

// WdkReportCfgInfoRes 文档库报告类型配置信息返回参数
type WdkReportCfgInfoRes struct {
	Info *WdkReportCfgInfo `json:"info" dc:"文档库报告类型配置信息"` // 文档库报告类型配置信息
}

// WdkReportCfgEditReq 文档库编辑报告类型配置请求参数
type WdkReportCfgEditReq struct {
	g.Meta    `path:"/edit" tags:"WdkReportCfgEdit" method:"put" summary:"You first wdk/reportTypeCfg/edit api"`
	Id        uint64   `json:"id" v:"required|regex:^[1-9]\\d*$#报告类型ID不能为空|报告类型ID必须为正整数" dc:"报告类型ID"`                       // 报告类型ID
	Name      string   `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#报告类型名称不能为空|报告类型名称必须为中文且长度不能超过20" dc:"报告类型名称"` // 报告类型名称
	AuditUids []uint64 `json:"auditUids" v:"required|slice_valid:uint64#文档库报告审核员用户ID列表不能为空" dc:"文档库报告审核员用户ID列表"`            // 文档库报告审核员用户ID列表
}

// WdkReportCfgEditRes 文档库编辑报告类型配置返回参数
type WdkReportCfgEditRes struct {
}

// WdkReportCfgDeleteReq 文档库删除报告类型配置请求参数
type WdkReportCfgDeleteReq struct {
	g.Meta `path:"/delete" tags:"WdkReportCfgDelete" method:"delete" summary:"You first wdk/reportTypeCfg/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#报告类型ID列表不能为空" dc:"报告类型ID列表"` // 报告类型ID列表
}

// WdkReportCfgDeleteRes 文档库删除报告类型配置返回参数
type WdkReportCfgDeleteRes struct {
}
