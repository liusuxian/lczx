package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// WdkProjectListReq 文档库项目列表请求参数
type WdkProjectListReq struct {
	g.Meta           `path:"/list" tags:"WdkProjectList" method:"get" summary:"You first wdk/project/list api"`
	Name             string      `json:"name" v:"regex:^[\u4e00-\u9fa5\\w]{0,100}#项目名称只能包含中文、英文、数字和下划线且长度不能超过100" dc:"项目名称"`                      // 项目名称
	Type             string      `json:"type" v:"in:0,1#项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                        // 项目性质 0:蓝绿体系 1:非绿
	Origin           string      `json:"origin" v:"in:0,1,2,3#项目来源只能是0,1,2,3" dc:"项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓"`                                 // 项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓
	Step             string      `json:"step" v:"in:0,1,2,3,4,5#项目阶段只能是0,1,2,3,4,5" dc:"项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘"`             // 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	FileUploadStatus string      `json:"fileUploadStatus" v:"in:0,1#项目文件上传状态只能是0,1" dc:"项目文件上传状态 0:未传完 1:已传完"`                                    // 项目文件上传状态 0:未传完 1:已传完
	BusinessType     string      `json:"businessType" v:"in:0,1,2#业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                        // 业务类型 0:物业 1:专项 2:全过程
	DeepCulture      string      `json:"deepCulture" v:"in:0,1#是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                   // 是否为深耕 0:否 1:是
	Status           string      `json:"status" v:"in:0,1,2,3#服务状态只能是0,1,2,3" dc:"服务状态 0:服务中 1:暂停 2:提前终止 3:正常结束"`                                 // 服务状态 0:服务中 1:暂停 2:提前终止 3:正常结束
	EntrustCompany   string      `json:"entrustCompany" v:"regex:^[\u4e00-\u9fa5\\da-zA-Z]{0,50}#委托方公司只能包含中文、英文、数字且长度不能超过50" dc:"委托方公司"`          // 委托方公司
	SignCompany      string      `json:"signCompany" v:"in:0,1,2#我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"` // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalName    string      `json:"principalName" v:"regex:^[\u4e00-\u9fa5]{0,10}$#负责人姓名必须为中文且长度不能超过10" dc:"负责人姓名"`                          // 负责人姓名
	DeptId           string      `json:"deptId" v:"regex:^[1-9]\\d*$#项目所属部门ID必须为正整数" dc:"项目所属部门ID"`                                               // 项目所属部门ID
	Region           string      `json:"region" v:"regex:^[\u4e00-\u9fa5]{0,50}$#地区(省/市/县)必须为中文且长度不能超过50" dc:"地区(省/市/县)"`                         // 地区(省/市/县)
	StartTime        *gtime.Time `json:"startTime" v:"datetime#项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                      // 项目开始时间
	EndTime          *gtime.Time `json:"endTime" v:"datetime#项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                        // 项目结束时间
	SortName         string      `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`                                     // 排序字段
	SortOrder        string      `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`                                    // 排序方式
	CurPage          int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`                                    // 当前页码
	PageSize         int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`                                   // 每页数量
}

// WdkProjectListRes 文档库项目列表返回参数
type WdkProjectListRes struct {
	CurPage int                  `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int                  `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.WdkProject `json:"list" dc:"文档库项目列表"` // 文档库项目列表
}

// WdkProjectAddReq 文档库新增项目请求参数
type WdkProjectAddReq struct {
	g.Meta         `path:"/add" tags:"WdkProjectAdd" method:"post" summary:"You first wdk/project/add api"`
	Name           string      `json:"name" v:"required|regex:^[\u4e00-\u9fa5\\w]{0,100}#项目名称不能为空|项目名称只能包含中文、英文、数字和下划线且长度不能超过100" dc:"项目名称"`                        // 项目名称
	Type           uint        `json:"type" v:"required|in:0,1#项目性质不能为空|项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                          // 项目性质 0:蓝绿体系 1:非绿
	Origin         uint        `json:"origin" v:"required|in:0,1,2,3#项目来源不能为空|项目来源只能是0,1,2,3" dc:"项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓"`                                   // 项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓
	BusinessType   uint        `json:"businessType" v:"required|in:0,1,2#业务类型不能为空|业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                          // 业务类型 0:物业 1:专项 2:全过程
	DeepCulture    uint        `json:"deepCulture" v:"required|in:0,1#是否为深耕不能为空|是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                    // 是否为深耕 0:否 1:是
	EntrustCompany string      `json:"entrustCompany" v:"required|regex:^[\u4e00-\u9fa5\\da-zA-Z]{0,50}#委托方公司不能为空|委托方公司只能包含中文、英文、数字且长度不能超过50" dc:"委托方公司"`           // 委托方公司
	SignCompany    uint        `json:"signCompany" v:"required|in:0,1,2#我方签订公司不能为空|我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"` // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalUid   uint64      `json:"principalUid" v:"required|regex:^[1-9]\\d*$#负责人用户ID不能为空|负责人用户ID必须为正整数" dc:"负责人用户ID"`                                          // 负责人用户ID
	DeptId         uint64      `json:"deptId" v:"required|regex:^[1-9]\\d*$#项目所属部门ID不能为空|项目所属部门ID必须为正整数" dc:"项目所属部门ID"`                                             // 项目所属部门ID
	Region         string      `json:"region" v:"required|regex:^[\u4e00-\u9fa5\\s]{0,50}$#地区(省/市/县)不能为空|地区(省/市/县)必须为中文且长度不能超过50" dc:"地区(省/市/县)"`                   // 地区(省/市/县)
	StartTime      *gtime.Time `json:"startTime" v:"required|datetime#项目开始时间不能为空|项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                      // 项目开始时间
	EndTime        *gtime.Time `json:"endTime" v:"required|datetime#项目结束时间不能为空|项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                        // 项目结束时间
	Remark         string      `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                               // 备注
}

// WdkProjectAddRes 文档库新增项目返回参数
type WdkProjectAddRes struct {
}

// WdkProjectEditReq 文档库编辑项目请求参数
type WdkProjectEditReq struct {
	g.Meta         `path:"/add" tags:"WdkProjectEdit" method:"put" summary:"You first wdk/project/edit api"`
	Name           string      `json:"name" v:"required|regex:^[\u4e00-\u9fa5\\w]{0,100}#项目名称不能为空|项目名称只能包含中文、英文、数字和下划线且长度不能超过100" dc:"项目名称"`                        // 项目名称
	Type           uint        `json:"type" v:"required|in:0,1#项目性质不能为空|项目性质只能是0,1" dc:"项目性质 0:蓝绿体系 1:非绿"`                                                          // 项目性质 0:蓝绿体系 1:非绿
	Origin         uint        `json:"origin" v:"required|in:0,1,2,3#项目来源不能为空|项目来源只能是0,1,2,3" dc:"项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓"`                                   // 项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓
	BusinessType   uint        `json:"businessType" v:"required|in:0,1,2#业务类型不能为空|业务类型只能是0,1,2" dc:"业务类型 0:物业 1:专项 2:全过程"`                                          // 业务类型 0:物业 1:专项 2:全过程
	DeepCulture    uint        `json:"deepCulture" v:"required|in:0,1#是否为深耕不能为空|是否为深耕只能是0,1" dc:"是否为深耕 0:否 1:是"`                                                    // 是否为深耕 0:否 1:是
	Status         uint        `json:"status" v:"required|in:0,1,2,3#服务状态不能为空|服务状态只能是0,1,2,3" dc:"服务状态 0:服务中 1:暂停 2:提前终止 3:正常结束"`                                   // 服务状态 0:服务中 1:暂停 2:提前终止 3:正常结束
	EntrustCompany string      `json:"entrustCompany" v:"required|regex:^[\u4e00-\u9fa5\\da-zA-Z]{0,50}#委托方公司不能为空|委托方公司只能包含中文、英文、数字且长度不能超过50" dc:"委托方公司"`           // 委托方公司
	SignCompany    uint        `json:"signCompany" v:"required|in:0,1,2#我方签订公司不能为空|我方签订公司只能是0,1,2" dc:"我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司"` // 我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司
	PrincipalUid   uint64      `json:"principalUid" v:"required|regex:^[1-9]\\d*$#负责人用户ID不能为空|负责人用户ID必须为正整数" dc:"负责人用户ID"`                                          // 负责人用户ID
	DeptId         uint64      `json:"deptId" v:"required|regex:^[1-9]\\d*$#项目所属部门ID不能为空|项目所属部门ID必须为正整数" dc:"项目所属部门ID"`                                             // 项目所属部门ID
	Region         string      `json:"region" v:"required|regex:^[\u4e00-\u9fa5]{0,50}$#地区(省/市/县)不能为空|地区(省/市/县)必须为中文且长度不能超过50" dc:"地区(省/市/县)"`                      // 地区(省/市/县)
	StartTime      *gtime.Time `json:"startTime" v:"required|datetime#项目开始时间不能为空|项目开始时间不是有效的日期时间" dc:"项目开始时间"`                                                      // 项目开始时间
	EndTime        *gtime.Time `json:"endTime" v:"required|datetime#项目结束时间不能为空|项目结束时间不是有效的日期时间" dc:"项目结束时间"`                                                        // 项目结束时间
	Remark         string      `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                               // 备注
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
