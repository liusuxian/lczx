package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// WdkFiletypeListReq 文档库上传文件类型列表请求参数
type WdkFiletypeListReq struct {
	g.Meta    `path:"/list" tags:"WdkFiletypeList" method:"get" summary:"You first wdk/filetype/list api"`
	Name      string      `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#类型名称必须为中文且长度不能超过20" dc:"类型名称"`                           // 类型名称
	Multiple  string      `json:"multiple" v:"in:0,1#是否同时存在多个文件只能是0,1" dc:"是否同时存在多个文件 0:否 1:是"`                                  // 是否同时存在多个文件 0:否 1:是
	Audit     string      `json:"audit" v:"in:0,1#是否需要审核只能是0,1" dc:"是否需要审核 0:不需要 1:需要"`                                          // 是否需要审核 0:不需要 1:需要
	Step      string      `json:"step" v:"in:0,1,2,3,4,5#项目阶段只能是0,1,2,3,4,5" dc:"所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘"` // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	StartTime *gtime.Time `json:"startTime" v:"datetime#开始时间不是有效的日期时间" dc:"开始时间"`                                                // 开始时间
	EndTime   *gtime.Time `json:"endTime" v:"datetime#结束时间不是有效的日期时间" dc:"结束时间"`                                                  // 结束时间
	CurPage   int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`                          // 当前页码
	PageSize  int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`                         // 每页数量
}

// WdkFiletypeListRes 文档库上传文件类型列表返回参数
type WdkFiletypeListRes struct {
	CurPage int                `json:"curPage" dc:"当前页码"`       // 当前页码
	Total   int                `json:"total" dc:"数据总量"`         // 数据总量
	List    []*WdkFiletypeInfo `json:"list" dc:"文档库上传文件类型信息列表"` // 文档库上传文件类型信息列表
}

// WdkFiletypeInfo 文档库上传文件类型信息
type WdkFiletypeInfo struct {
	FiletypeCfg *entity.WdkFiletypeCfg `json:"filetypeCfg" dc:"文档库上传文件类型配置信息"` // 文档库上传文件类型配置信息
	AuditCfg    []*entity.WdkAuditCfg  `json:"auditCfg" dc:"文档库上传文件类型审核配置信息"`  // 文档库上传文件类型审核配置信息
}

// WdkFiletypeAddReq 文档库新增上传文件类型请求参数
type WdkFiletypeAddReq struct {
	g.Meta    `path:"/add" tags:"WdkFiletypeAdd" method:"post" summary:"You first wdk/filetype/add api"`
	Name      string   `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#类型名称不能为空|类型名称必须为中文且长度不能超过20" dc:"类型名称"`                               // 类型名称
	Multiple  uint     `json:"multiple" v:"required|in:0,1#是否同时存在多个文件不能为空|是否同时存在多个文件只能是0,1" dc:"是否同时存在多个文件 0:否 1:是"`                                // 是否同时存在多个文件 0:否 1:是
	Audit     uint     `json:"audit" v:"required|in:0,1#是否需要审核不能为空|是否需要审核只能是0,1" dc:"是否需要审核 0:不需要 1:需要"`                                            // 是否需要审核 0:不需要 1:需要
	Step      uint     `json:"step" v:"required|in:0,1,2,3,4,5#所属项目阶段不能为空|所属项目阶段只能是0,1,2,3,4,5" dc:"所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘"` // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	AuditUids []uint64 `json:"auditUids" v:"required-if:audit,1|slice_valid:uint64#文档库报告审核员用户ID列表不能为空" dc:"文档库报告审核员用户ID列表"`                         // 文档库报告审核员用户ID列表
}

// WdkFiletypeAddRes 文档库新增上传文件类型返回参数
type WdkFiletypeAddRes struct {
}

// WdkFiletypeInfoReq 文档库上传文件类型信息请求参数
type WdkFiletypeInfoReq struct {
	g.Meta `path:"/info" tags:"WdkFiletypeInfo" method:"get" summary:"You first wdk/filetype/info api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#文档库上传文件类型ID不能为空|文档库上传文件类型ID必须为正整数" dc:"文档库上传文件类型ID"` // 文档库上传文件类型ID
}

// WdkFiletypeInfoRes 文档库上传文件类型信息返回参数
type WdkFiletypeInfoRes struct {
	Info *WdkFiletypeInfo `json:"info" dc:"文档库上传文件类型信息"` // 文档库上传文件类型信息
}

// WdkFiletypeEditReq 文档库编辑上传文件类型请求参数
type WdkFiletypeEditReq struct {
	g.Meta    `path:"/edit" tags:"WdkFiletypeEdit" method:"put" summary:"You first wdk/filetype/edit api"`
	Id        uint64   `json:"id" v:"required|regex:^[1-9]\\d*$#文档库上传文件类型ID不能为空|文档库上传文件类型ID必须为正整数" dc:"文档库上传文件类型ID"`                                // 上传文件类型ID
	Name      string   `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#类型名称不能为空|类型名称必须为中文且长度不能超过20" dc:"类型名称"`                               // 类型名称
	Multiple  uint     `json:"multiple" v:"required|in:0,1#是否同时存在多个文件不能为空|是否同时存在多个文件只能是0,1" dc:"是否同时存在多个文件 0:否 1:是"`                                // 是否同时存在多个文件 0:否 1:是
	Audit     uint     `json:"audit" v:"required|in:0,1#是否需要审核不能为空|是否需要审核只能是0,1" dc:"是否需要审核 0:不需要 1:需要"`                                            // 是否需要审核 0:不需要 1:需要
	Step      uint     `json:"step" v:"required|in:0,1,2,3,4,5#所属项目阶段不能为空|所属项目阶段只能是0,1,2,3,4,5" dc:"所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘"` // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	AuditUids []uint64 `json:"auditUids" v:"required-if:audit,1|slice_valid:uint64#文档库报告审核员用户ID列表不能为空" dc:"文档库报告审核员用户ID列表"`                         // 文档库报告审核员用户ID列表
}

// WdkFiletypeEditRes 文档库编辑上传文件类型返回参数
type WdkFiletypeEditRes struct {
}

// WdkFiletypeDeleteReq 文档库删除上传文件类型请求参数
type WdkFiletypeDeleteReq struct {
	g.Meta `path:"/delete" tags:"WdkFiletypeDelete" method:"delete" summary:"You first wdk/filetype/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#文档库上传文件类型ID不能为空" dc:"文档库上传文件类型ID"` // 文档库上传文件类型ID
}

// WdkFiletypeDeleteRes 文档库删除上传文件类型返回参数
type WdkFiletypeDeleteRes struct {
}
