package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// WdkFiletypeListReq 文档库上传文件类型列表请求参数
type WdkFiletypeListReq struct {
	g.Meta    `path:"/list" tags:"WdkProjectList" method:"get" summary:"You first wdk/project/list api"`
	Name      string `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#类型名称必须为中文且长度不能超过20" dc:"类型名称"`                           // 类型名称
	Multiple  uint   `json:"multiple" v:"in:0,1#是否同时存在多个文件只能是0,1" dc:"是否同时存在多个文件 0:否 1:是"`                                  // 是否同时存在多个文件 0:否 1:是
	Audit     uint   `json:"audit" v:"in:0,1#是否需要审核只能是0,1" dc:"是否需要审核 0:不需要 1:需要"`                                          // 是否需要审核 0:不需要 1:需要
	Step      uint   `json:"step" v:"in:0,1,2,3,4,5#项目阶段只能是0,1,2,3,4,5" dc:"所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘"` // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	SortName  string `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`                           // 排序字段
	SortOrder string `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`                          // 排序方式
	CurPage   int    `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`                          // 当前页码
	PageSize  int    `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`                         // 每页数量
}

// WdkFiletypeListRes 文档库上传文件类型列表返回参数
type WdkFiletypeListRes struct {
	CurPage int                `json:"curPage" dc:"当前页码"`       // 当前页码
	Total   int                `json:"total" dc:"数据总量"`         // 数据总量
	List    []*WdkFiletypeInfo `json:"list" dc:"文档库上传文件类型信息列表"` // 文档库上传文件类型信息列表
}

// WdkFiletypeInfo 文档库上传文件类型信息
type WdkFiletypeInfo struct {
	FiletypeCfg *entity.WdkFiletypeCfg `json:"filetypeCfg" dc:"上传文件类型配置信息"` // 上传文件类型配置信息
	AuditCfg    []*entity.WdkAuditCfg  `json:"auditCfg" dc:"上传文件类型审核配置信息"`  // 上传文件类型审核配置信息
}
