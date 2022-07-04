package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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

// WdkReportAddReq 文档库新增上传报告请求参数
type WdkReportAddReq struct {
	g.Meta         `path:"/add" tags:"WdkReportAdd" method:"post" summary:"You first wdk/report/add api"`
	ProjectId      uint64                    `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"`                                                                                         // 所属项目ID
	ProjectName    string                    `json:"projectName" v:"required#项目名称不能为空" dc:"所属项目名称"`                                                                                                                        // 所属项目名称
	Step           uint                      `json:"step" v:"required|in:3,4,5,6,7,8,9#项目阶段不能为空|项目阶段只能是3,4,5,6,7,8,9" dc:"项目阶段 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收"` // 项目阶段 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收
	AuditTypeInfos []*WdkReportAuditTypeInfo `json:"auditTypeInfos" v:"required#文档库报告审核类型信息列表不能为空" dc:"文档库报告审核类型信息列表"`                                                                                                     // 文档库报告审核类型信息列表
	UploadReport   *ghttp.UploadFile         `json:"uploadReport" dc:"上传报告"`                                                                                                                                               // 上传报告
}

// WdkReportAddRes 文档库新增上传报告返回参数
type WdkReportAddRes struct {
}

// WdkReportAuditTypeInfo 文档库报告审核类型信息
type WdkReportAuditTypeInfo struct {
	TypeId    uint64   `json:"typeId" dc:"报告类型ID"`            // 报告类型ID
	AuditUids []uint64 `json:"auditUids" dc:"文档库报告审核员用户ID列表"` // 文档库报告审核员用户ID列表
}

// WdkReportDeleteReq 文档库删除报告请求参数
type WdkReportDeleteReq struct {
	g.Meta `path:"/delete" tags:"WdkReportDelete" method:"delete" summary:"You first wdk/report/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#文档库报告ID列表不能为空" dc:"文档库报告ID列表"` // 文档库报告ID列表
}

// WdkReportDeleteRes 文档库删除报告返回参数
type WdkReportDeleteRes struct {
}

// WdkReportExcellenceListReq 文档库优秀报告列表请求参数
type WdkReportExcellenceListReq struct {
	g.Meta      `path:"/excellenceList" tags:"WdkReportExcellenceList" method:"get" summary:"You first wdk/report/excellenceList api"`
	Excellence  uint   `json:"excellence" v:"required|in:1,2#是否是优秀报告不能为空|是否是优秀报告只能是1,2" dc:"是否是优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告"` // 是否是优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告
	TypeId      string `json:"typeId" v:"regex:^[1-9]\\d*$#报告类型ID必须为正整数" dc:"报告类型ID"`                                        // 报告类型ID
	ReportName  string `json:"reportName" dc:"报告名称"`                                                                         // 报告名称
	ProjectName string `json:"projectName" dc:"所属项目名称"`                                                                      // 所属项目名称
	CurPage     int    `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`                         // 当前页码
	PageSize    int    `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`                        // 每页数量
}

// WdkReportExcellenceListRes 文档库优秀报告列表返回参数
type WdkReportExcellenceListRes struct {
	CurPage int              `json:"curPage" dc:"当前页码"`     // 当前页码
	Total   int              `json:"total" dc:"数据总量"`       // 数据总量
	List    []*WdkReportInfo `json:"list" dc:"文档库上传报告信息列表"` // 文档库上传报告信息列表
}

// WdkReportChooseExcellenceReq 文档库报告评选优秀报告请求参数
type WdkReportChooseExcellenceReq struct {
	g.Meta     `path:"/chooseExcellence" tags:"WdkReportChooseExcellence" method:"put" summary:"You first wdk/report/chooseExcellence api"`
	Id         uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#报告ID不能为空|报告ID必须为正整数" dc:"报告ID"`
	Excellence uint   `json:"excellence" v:"required|in:0,2#是否是优秀报告不能为空|是否是优秀报告只能是0,2" dc:"是否是优秀报告 0:未被评选为优秀报告 2:已被评选为优秀报告"` // 是否是优秀报告 0:未被评选为优秀报告 2:已被评选为优秀报告
}

// WdkReportChooseExcellenceRes 文档库报告评选优秀报告返回参数
type WdkReportChooseExcellenceRes struct {
}
