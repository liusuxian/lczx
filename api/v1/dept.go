package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// DeptListReq 部门树列表请求参数
type DeptListReq struct {
	g.Meta `path:"/list" tags:"DeptList" method:"get" summary:"You first auth/dept/list api"`
	Name   string `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#部门名称必须为中文且长度不能超过20" dc:"部门名称"` // 部门名称
	Status string `json:"status" v:"in:0,1#部门状态只能是0,1" dc:"部门状态 0:停用 1:正常"`                    // 部门状态 0:停用 1:正常
}

// DeptListRes 部门树列表返回参数
type DeptListRes struct {
	List []*DeptTreeInfo `json:"list" dc:"部门树列表"` // 部门树列表
}

// DeptTreeInfo 部门树信息
type DeptTreeInfo struct {
	Dept     *entity.Dept    `json:"dept" dc:"部门信息"`        // 部门信息
	Children []*DeptTreeInfo `json:"children" dc:"子部门信息列表"` // 子部门信息列表
}

// DeptAddReq 新增部门请求参数
type DeptAddReq struct {
	g.Meta   `path:"/add" tags:"DeptAdd" method:"post" summary:"You first auth/dept/add api"`
	ParentId uint64 `json:"parentId" v:"required|regex:^\\d*$#父部门ID不能为空|父部门ID必须为无符号整数" dc:"父部门ID"`                 // 父部门ID
	Name     string `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#部门名称不能为空|部门名称必须为中文且长度不能超过20" dc:"部门名称"` // 部门名称
	Status   uint   `json:"status" v:"required|in:0,1#部门状态不能为空|部门状态只能是0,1" dc:"部门状态 0:停用 1:正常"`                    // 部门状态 0:停用 1:正常
}

// DeptAddRes 新增部门返回参数
type DeptAddRes struct {
}

// DeptInfoReq 获取部门信息请求参数
type DeptInfoReq struct {
	g.Meta `path:"/info" tags:"DeptInfo" method:"get" summary:"You first auth/dept/info api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"` // 部门ID
}

// DeptInfoRes 获取部门信息返回参数
type DeptInfoRes struct {
	Info           *entity.Dept    `json:"info" dc:"部门信息"`            // 部门信息
	List           []*DeptTreeInfo `json:"list" dc:"部门树列表"`           // 部门树列表
	ParentIdStatus uint            `json:"parentIdStatus" dc:"父部门状态"` // 父部门状态
}

// DeptEditReq 编辑部门请求参数
type DeptEditReq struct {
	g.Meta   `path:"/edit" tags:"DeptEdit" method:"put" summary:"You first auth/dept/edit api"`
	Id       uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"`                                   // 部门ID
	ParentId uint64 `json:"parentId" v:"required|regex:^\\d*$|different:Id#父部门ID不能为空|父部门ID必须为无符号整数|父部门ID和部门ID不能相同" dc:"父部门ID"` // 父部门ID
	Name     string `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#部门名称不能为空|部门名称必须为中文且长度不能超过20" dc:"部门名称"`             // 部门名称
	Status   uint   `json:"status" v:"required|in:0,1#部门状态不能为空|部门状态只能是0,1" dc:"部门状态 0:停用 1:正常"`                                // 部门状态 0:停用 1:正常
}

// DeptEditRes 编辑部门返回参数
type DeptEditRes struct {
}

// DeptDeleteReq 删除部门请求参数
type DeptDeleteReq struct {
	g.Meta `path:"/delete" tags:"DeptDelete" method:"delete" summary:"You first auth/dept/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#部门ID列表不能为空" dc:"部门ID列表"` // 部门ID列表
}

// DeptDeleteRes 删除部门返回参数
type DeptDeleteRes struct {
}

// DeptTreeReq 部门树信息请求参数
type DeptTreeReq struct {
	g.Meta `path:"/tree" tags:"DeptTree" method:"get" summary:"You first auth/dept/tree api"`
}

// DeptTreeRes 部门树信息返回参数
type DeptTreeRes struct {
	List []*DeptTreeInfo `json:"list" dc:"部门树列表"` // 部门树列表
}

// DeptRoleDeptTreeReq 角色部门树信息请求参数
type DeptRoleDeptTreeReq struct {
	g.Meta `path:"/roleDeptTree" tags:"DeptRoleDeptTree" method:"get" summary:"You first auth/dept/roleDeptTree api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#角色ID不能为空|角色ID必须为正整数" dc:"角色ID"` // 角色ID
}

// DeptRoleDeptTreeRes 角色部门树信息返回参数
type DeptRoleDeptTreeRes struct {
	List    []*DeptTreeInfo `json:"list" dc:"部门树列表"`          // 部门树列表
	DeptIds []uint64        `json:"deptIds" dc:"角色关联的部门ID列表"` // 角色关联的部门ID列表
}
