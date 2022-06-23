package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// WdkProjectListReq 文档库项目列表请求参数
type WdkProjectListReq struct {
	g.Meta           `path:"/list" tags:"WdkProjectList" method:"get" summary:"You first wdk/project/list api"`
	Name             string      `json:"name" v:"max-length:100#项目名称长度不能超过100" dc:"项目名称"`                                                                                                                                                               // 项目名称
	Type             string      `json:"type" v:"in:0,1#项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                                                                                                                              // 项目性质 0:蓝绿体系 1:非绿
	Origin           string      `json:"origin" v:"in:0,1,2,3,4,5,6,7#项目来源只能是0,1,2,3,4,5,6,7" dc:"项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他"`                                                                                                    // 项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他
	Step             string      `json:"step" v:"in:0,1,2,3,4,5,6,7,8,9,30,31#项目阶段只能是0,1,2,3,4,5,6,7,8,9,30,31" dc:"项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘"` // 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘
	FileUploadStatus string      `json:"fileUploadStatus" v:"in:0,1,2#项目文件上传状态只能是0,1,2" dc:"项目文件上传状态 0:异常 1:正常 2:已完成"`                                                                                                                                  // 项目文件上传状态 0:异常 1:正常 2:已完成
	BusinessType     string      `json:"businessType" v:"in:0,1,2#业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                                                                                                                              // 业务类型 0:物业 1:专项 2:全过程
	BusinessForms    []uint      `json:"businessForms" v:"slice_valid:uint#业态ID列表不能为空" dc:"业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房"`                       // 业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房
	ContractStatus   string      `json:"contractStatus" v:"in:0,1,2#签约状态只能是0,1,2" dc:"签约状态 0:新签 1:续签 2:未签"`                                                                                                                                             // 签约状态 0:新签 1:续签 2:未签
	ContractSum      string      `json:"contractSum" v:"regex:^\\d*\\.*\\d*$#合同金额必须大于等于0" dc:"合同金额"`                                                                                                                                                    // 合同金额
	DeepCulture      string      `json:"deepCulture" v:"in:0,1#是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                                                                                                                         // 是否为深耕 0:否 1:是
	Status           string      `json:"status" v:"in:0,1,2,3,4,5#服务状态只能是0,1,2,3,4,5" dc:"服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束"`                                                                                                                   // 服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束
	EntrustCompany   string      `json:"entrustCompany" v:"max-length:50#委托方公司长度不能超过50" dc:"委托方公司"`                                                                                                                                                     // 委托方公司
	SignCompany      string      `json:"signCompany" v:"in:0,1,2#我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"`                                                                                                       // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalName    string      `json:"principalName" v:"regex:^[\u4e00-\u9fa5]{0,10}$#负责人姓名必须为中文且长度不能超过10" dc:"负责人姓名"`                                                                                                                                // 负责人姓名
	DeptId           string      `json:"deptId" v:"regex:^[1-9]\\d*$#项目所属部门ID必须为正整数" dc:"项目所属部门ID"`                                                                                                                                                     // 项目所属部门ID
	Region           string      `json:"region" v:"max-length:50#地区(省/市/区县)长度不能超过50" dc:"地区(省/市/区县)"`                                                                                                                                                   // 地区(省/市/区县)
	StartTime        *gtime.Time `json:"startTime" v:"date#项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                                                                                                                                // 项目开始时间
	EndTime          *gtime.Time `json:"endTime" v:"date#项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                                                                                                                                  // 项目结束时间
	SortName         string      `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`                                                                                                                                           // 排序字段
	SortOrder        string      `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`                                                                                                                                          // 排序方式
	CurPage          int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`                                                                                                                                          // 当前页码
	PageSize         int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`                                                                                                                                         // 每页数量
}

// WdkProjectListRes 文档库项目列表返回参数
type WdkProjectListRes struct {
	CurPage int               `json:"curPage" dc:"当前页码"`   // 当前页码
	Total   int               `json:"total" dc:"数据总量"`     // 数据总量
	List    []*WdkProjectInfo `json:"list" dc:"文档库项目信息列表"` // 文档库项目信息列表
}

// WdkProjectInfo 文档库项目信息
type WdkProjectInfo struct {
	ProjectInfo   *entity.WdkProject                `json:"projectInfo" dc:"文档库项目信息"`       // 文档库项目信息
	Businessforms []*entity.WdkProjectBusinessforms `json:"businessforms" dc:"文档库项目业态信息列表"` // 文档库项目业态信息列表
}

// WdkProjectAddReq 文档库新增项目请求参数
type WdkProjectAddReq struct {
	g.Meta         `path:"/add" tags:"WdkProjectAdd" method:"post" summary:"You first wdk/project/add api"`
	Name           string      `json:"name" v:"required|max-length:100#项目名称不能为空|项目名称长度不能超过100" dc:"项目名称"`                                                                                                                                       // 项目名称
	Type           uint        `json:"type" v:"required|in:0,1#项目性质不能为空|项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                                                                                                      // 项目性质 0:蓝绿体系 1:非绿
	Origin         uint        `json:"origin" v:"required|in:0,1,2,3,4,5,6,7#项目来源不能为空|项目来源只能是0,1,2,3,4,5,6,7" dc:"项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他"`                                                                            // 项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他
	BusinessType   uint        `json:"businessType" v:"required|in:0,1,2#业务类型不能为空|业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                                                                                                      // 业务类型 0:物业 1:专项 2:全过程
	BusinessForms  []uint      `json:"businessForms" v:"required|slice_valid:uint#业态不能为空|业态ID列表不能为空" dc:"业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房"` // 业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房
	ContractStatus uint        `json:"contractStatus" v:"required|in:0,1,2#签约状态不能为空|签约状态只能是0,1,2" dc:"签约状态 0:新签 1:续签 2:未签"`                                                                                                                     // 签约状态 0:新签 1:续签 2:未签
	ContractSum    float64     `json:"contractSum" v:"required|regex:^(\\d+)(\\.\\d+)?$#合同金额不能为空|合同金额必须大于等于0" dc:"合同金额"`                                                                                                                        // 合同金额
	DeepCulture    uint        `json:"deepCulture" v:"required|in:0,1#是否为深耕不能为空|是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                                                                                                // 是否为深耕 0:否 1:是
	Status         string      `json:"status" v:"required|in:0,1,2,3,4,5#服务状态不能为空|服务状态只能是0,1,2,3,4,5" dc:"服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束"`                                                                                           // 服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束
	EntrustCompany string      `json:"entrustCompany" v:"required|max-length:50#委托方公司不能为空|委托方公司长度不能超过50" dc:"委托方公司"`                                                                                                                            // 委托方公司
	SignCompany    uint        `json:"signCompany" v:"required|in:0,1,2#我方签订公司不能为空|我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"`                                                                             // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalUid   uint64      `json:"principalUid" v:"required|regex:^[1-9]\\d*$#负责人用户ID不能为空|负责人用户ID必须为正整数" dc:"负责人用户ID"`                                                                                                                      // 负责人用户ID
	DeptId         uint64      `json:"deptId" v:"required|regex:^[1-9]\\d*$#所属部门ID不能为空|所属部门ID必须为正整数" dc:"所属部门ID"`                                                                                                                               // 所属部门ID
	Region         string      `json:"region" v:"required|max-length:50#地区(省/市/区县)不能为空|地区(省/市/区县)长度不能超过50" dc:"地区(省/市/区县)"`                                                                                                                     // 地区(省/市/区县)
	StartTime      *gtime.Time `json:"startTime" v:"required|date#项目开始时间不能为空|项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                                                                                                      // 项目开始时间
	EndTime        *gtime.Time `json:"endTime" v:"required|date#项目结束时间不能为空|项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                                                                                                        // 项目结束时间
	Remark         string      `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                                                                                                           // 备注
}

// WdkProjectAddRes 文档库新增项目返回参数
type WdkProjectAddRes struct {
}

// WdkProjectInfoReq 文档库项目信息请求参数
type WdkProjectInfoReq struct {
	g.Meta `path:"/info" tags:"WdkProjectInfo" method:"get" summary:"You first wdk/project/info api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#文档库项目ID不能为空|文档库项目ID必须为正整数" dc:"文档库项目ID"` // 文档库项目ID
}

// WdkProjectInfoRes 文档库项目信息返回参数
type WdkProjectInfoRes struct {
	Info *WdkProjectInfo `json:"info" dc:"文档库项目信息"` // 文档库项目信息
}

// WdkProjectEditReq 文档库编辑项目请求参数
type WdkProjectEditReq struct {
	g.Meta         `path:"/edit" tags:"WdkProjectEdit" method:"put" summary:"You first wdk/project/edit api"`
	Id             uint64      `json:"id" v:"required|regex:^[1-9]\\d*$#文档库项目ID不能为空|文档库项目ID必须为正整数" dc:"文档库项目ID"`                                                                                                                                // 文档库项目ID
	Name           string      `json:"name" v:"required|max-length:100#项目名称不能为空|项目名称长度不能超过100" dc:"项目名称"`                                                                                                                                       // 项目名称
	Type           uint        `json:"type" v:"required|in:0,1#项目性质不能为空|项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                                                                                                      // 项目性质 0:蓝绿体系 1:非绿
	Origin         uint        `json:"origin" v:"required|in:0,1,2,3,4,5,6,7#项目来源不能为空|项目来源只能是0,1,2,3,4,5,6,7" dc:"项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他"`                                                                            // 项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他
	BusinessType   uint        `json:"businessType" v:"required|in:0,1,2#业务类型不能为空|业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                                                                                                      // 业务类型 0:物业 1:专项 2:全过程
	BusinessForms  []uint      `json:"businessForms" v:"required|slice_valid:uint#业态不能为空|业态ID列表不能为空" dc:"业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房"` // 业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房
	ContractStatus uint        `json:"contractStatus" v:"required|in:0,1,2#签约状态不能为空|签约状态只能是0,1,2" dc:"签约状态 0:新签 1:续签 2:未签"`                                                                                                                     // 签约状态 0:新签 1:续签 2:未签
	ContractSum    float64     `json:"contractSum" v:"required|regex:^(\\d+)(\\.\\d+)?$#合同金额不能为空|合同金额必须大于等于0" dc:"合同金额"`                                                                                                                        // 合同金额
	DeepCulture    uint        `json:"deepCulture" v:"required|in:0,1#是否为深耕不能为空|是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                                                                                                // 是否为深耕 0:否 1:是
	Status         string      `json:"status" v:"required|in:0,1,2,3,4,5#服务状态不能为空|服务状态只能是0,1,2,3,4,5" dc:"服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束"`                                                                                           // 服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束
	EntrustCompany string      `json:"entrustCompany" v:"required|max-length:50#委托方公司不能为空|委托方公司长度不能超过50" dc:"委托方公司"`                                                                                                                            // 委托方公司
	SignCompany    uint        `json:"signCompany" v:"required|in:0,1,2#我方签订公司不能为空|我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"`                                                                             // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalUid   uint64      `json:"principalUid" v:"required|regex:^[1-9]\\d*$#负责人用户ID不能为空|负责人用户ID必须为正整数" dc:"负责人用户ID"`                                                                                                                      // 负责人用户ID
	DeptId         uint64      `json:"deptId" v:"required|regex:^[1-9]\\d*$#所属部门ID不能为空|所属部门ID必须为正整数" dc:"所属部门ID"`                                                                                                                               // 所属部门ID
	Region         string      `json:"region" v:"required|max-length:50#地区(省/市/区县)不能为空|地区(省/市/区县)长度不能超过50" dc:"地区(省/市/区县)"`                                                                                                                     // 地区(省/市/区县
	StartTime      *gtime.Time `json:"startTime" v:"required|date#项目开始时间不能为空|项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                                                                                                      // 项目开始时间
	EndTime        *gtime.Time `json:"endTime" v:"required|date#项目结束时间不能为空|项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                                                                                                        // 项目结束时间
	Remark         string      `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                                                                                                           // 备注
}

// WdkProjectEditRes 文档库编辑项目返回参数
type WdkProjectEditRes struct {
}

// WdkProjectDeleteReq 文档库删除项目请求参数
type WdkProjectDeleteReq struct {
	g.Meta `path:"/delete" tags:"WdkProjectDelete" method:"delete" summary:"You first wdk/project/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#文档库项目ID列表不能为空" dc:"文档库项目ID列表"` // 文档库项目ID列表
}

// WdkProjectDeleteRes 文档库删除项目返回参数
type WdkProjectDeleteRes struct {
}

// WdkProjectExportReq 文档库项目信息导出请求参数
type WdkProjectExportReq struct {
	g.Meta           `path:"/export" tags:"WdkProjectExport" method:"get" summary:"You first wdk/project/export api"`
	Name             string      `json:"name" v:"max-length:100#项目名称长度不能超过100" dc:"项目名称"`                                                                                                                                                               // 项目名称
	Type             string      `json:"type" v:"in:0,1#项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                                                                                                                              // 项目性质 0:蓝绿体系 1:非绿
	Origin           string      `json:"origin" v:"in:0,1,2,3,4,5,6,7#项目来源只能是0,1,2,3,4,5,6,7" dc:"项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他"`                                                                                                    // 项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他
	Step             string      `json:"step" v:"in:0,1,2,3,4,5,6,7,8,9,30,31#项目阶段只能是0,1,2,3,4,5,6,7,8,9,30,31" dc:"项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘"` // 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘
	FileUploadStatus string      `json:"fileUploadStatus" v:"in:0,1,2#项目文件上传状态只能是0,1,2" dc:"项目文件上传状态 0:异常 1:正常 2:已完成"`                                                                                                                                  // 项目文件上传状态 0:异常 1:正常 2:已完成
	BusinessType     string      `json:"businessType" v:"in:0,1,2#业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                                                                                                                              // 业务类型 0:物业 1:专项 2:全过程
	BusinessForms    []uint      `json:"businessForms" v:"slice_valid:uint#业态ID列表不能为空" dc:"业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房"`                       // 业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房
	ContractStatus   string      `json:"contractStatus" v:"in:0,1,2#签约状态只能是0,1,2" dc:"签约状态 0:新签 1:续签 2:未签"`                                                                                                                                             // 签约状态 0:新签 1:续签 2:未签
	ContractSum      string      `json:"contractSum" v:"regex:^\\d*\\.*\\d*$#合同金额必须大于等于0" dc:"合同金额"`                                                                                                                                                    // 合同金额
	DeepCulture      string      `json:"deepCulture" v:"in:0,1#是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                                                                                                                         // 是否为深耕 0:否 1:是
	Status           string      `json:"status" v:"in:0,1,2,3,4,5#服务状态只能是0,1,2,3,4,5" dc:"服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束"`                                                                                                                   // 服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束
	EntrustCompany   string      `json:"entrustCompany" v:"max-length:50#委托方公司长度不能超过50" dc:"委托方公司"`                                                                                                                                                     // 委托方公司
	SignCompany      string      `json:"signCompany" v:"in:0,1,2#我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"`                                                                                                       // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalName    string      `json:"principalName" v:"regex:^[\u4e00-\u9fa5]{0,10}$#负责人姓名必须为中文且长度不能超过10" dc:"负责人姓名"`                                                                                                                                // 负责人姓名
	DeptId           string      `json:"deptId" v:"regex:^[1-9]\\d*$#项目所属部门ID必须为正整数" dc:"项目所属部门ID"`                                                                                                                                                     // 项目所属部门ID
	Region           string      `json:"region" v:"max-length:50#地区(省/市/区县)长度不能超过50" dc:"地区(省/市/区县)"`                                                                                                                                                   // 地区(省/市/区县)
	StartTime        *gtime.Time `json:"startTime" v:"date#项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                                                                                                                                // 项目开始时间
	EndTime          *gtime.Time `json:"endTime" v:"date#项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                                                                                                                                  // 项目结束时间
}

// WdkProjectExportRes 文档库项目信息导出返回参数
type WdkProjectExportRes struct {
	FileInfo *WdkProjectExportFile `json:"fileInfo" dc:"导出的文件信息"` // 导出的文件信息
}

// WdkProjectExportFile 文档库项目信息导出的文件信息
type WdkProjectExportFile struct {
	FileName string `json:"fileName" dc:"文件名"`       // 文件名
	FilePath string `json:"filePath" dc:"文件在服务器的路径"` // 文件在服务器的路径
}
