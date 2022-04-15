package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// RoleListReq 获取角色列表请求参数
type RoleListReq struct {
	g.Meta    `path:"/list" tags:"RoleList" method:"get" summary:"You first auth/role/list api"`
	Name      string      `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#角色名称必须为中文且长度不能超过20" dc:"角色名称"`   // 角色名称
	Status    string      `json:"status" v:"in:0,1#角色状态只能是0,1" dc:"状态 0:停用 1:正常"`                        // 状态 0:停用 1:正常
	StartTime *gtime.Time `json:"startTime" v:"datetime#开始时间不是有效的日期时间" dc:"开始时间"`                        // 开始时间
	EndTime   *gtime.Time `json:"endTime" v:"datetime#结束时间不是有效的日期时间" dc:"结束时间"`                          // 结束时间
	CurPage   int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`  // 当前页码
	PageSize  int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"` // 每页数量
}

// RoleListRes 获取角色列表返回参数
type RoleListRes struct {
	CurPage int            `json:"curPage" dc:"当前页码"` // 当前页码
	Total   int            `json:"total" dc:"数据总量"`   // 数据总量
	List    []*entity.Role `json:"list" dc:"角色列表"`    // 角色列表
}

// RoleAddReq 新增角色请求参数
type RoleAddReq struct {
	g.Meta  `path:"/add" tags:"RoleAdd" method:"post" summary:"You first auth/role/add api"`
	Name    string   `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#角色名称不能为空|角色名称必须为中文且长度不能超过20" dc:"角色名称"` // 角色名称
	Status  uint     `json:"status" v:"required|in:0,1#角色状态不能为空|角色状态只能是0,1" dc:"角色状态 0:停用 1:正常"`                    // 角色状态 0:停用 1:正常
	MenuIds []uint64 `json:"menuIds" v:"required|slice_valid:uint64#菜单ID列表不能为空" dc:"菜单ID列表"`                        // 菜单ID列表
	Remark  string   `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                         // 备注
}

// RoleAddRes 新增角色返回参数
type RoleAddRes struct {
}

// RoleInfoReq 获取角色信息请求参数
type RoleInfoReq struct {
	g.Meta `path:"/info" tags:"RoleInfo" method:"get" summary:"You first auth/role/info api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#角色ID不能为空|角色ID必须为正整数" dc:"角色ID"` // 角色ID
}

// RoleInfoRes 获取角色信息返回参数
type RoleInfoRes struct {
	Info     *entity.Role    `json:"info" dc:"角色信息"`           // 角色信息
	MenuList []*MenuTreeInfo `json:"list" dc:"菜单树列表"`          // 菜单树列表
	MenuIds  []uint64        `json:"menuIds" dc:"角色关联的菜单ID列表"` // 角色关联的菜单ID列表
}

// RoleEditReq 编辑角色请求参数
type RoleEditReq struct {
	g.Meta  `path:"/edit" tags:"RoleEdit" method:"put" summary:"You first auth/role/edit api"`
	Id      uint64   `json:"id" v:"required|regex:^[1-9]\\d*$#角色ID不能为空|角色ID必须为正整数" dc:"角色ID"`                       // 角色ID
	Name    string   `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#角色名称不能为空|角色名称必须为中文且长度不能超过20" dc:"角色名称"` // 角色名称
	Status  uint     `json:"status" v:"required|in:0,1#角色状态不能为空|角色状态只能是0,1" dc:"角色状态 0:停用 1:正常"`                    // 角色状态 0:停用 1:正常
	MenuIds []uint64 `json:"menuIds" v:"required|slice_valid:uint64#菜单ID列表不能为空" dc:"菜单ID列表"`                        // 菜单ID列表
	Remark  string   `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                         // 备注
}

// RoleEditRes 编辑角色返回参数
type RoleEditRes struct {
}

// RoleSetStatusReq 设置角色状态请求参数
type RoleSetStatusReq struct {
	g.Meta `path:"/setStatus" tags:"RoleSetStatus" method:"put" summary:"You first auth/role/setStatus api"`
	Id     uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#角色ID不能为空|角色ID必须为正整数" dc:"角色ID"`    // 角色ID
	Status uint   `json:"status" v:"required|in:0,1#角色状态不能为空|角色状态只能是0,1" dc:"角色状态 0:停用 1:正常"` // 角色状态 0:停用 1:正常
}

// RoleSetStatusRes 设置角色状态返回参数
type RoleSetStatusRes struct {
}

// RoleSetDataScopeReq 设置数据权限请求参数
type RoleSetDataScopeReq struct {
	g.Meta    `path:"/setDataScope" tags:"RoleSetDataScope" method:"put" summary:"You first auth/role/setDataScope api"`
	Id        uint64   `json:"id" v:"required|regex:^[1-9]\\d*$#角色ID不能为空|角色ID必须为正整数" dc:"角色ID"`                                                             // 角色ID
	DataScope uint     `json:"dataScope" v:"required|in:1,2,3,4,5#数据范围不能为空|数据范围只能是1,2,3,4,5" dc:"数据范围 1:全部数据权限 2:自定义数据权限 3:本部门数据权限 4:本部门及以下数据权限 5:仅本人数据权限"` // 数据范围 1:全部数据权限 2:自定义数据权限 3:本部门数据权限 4:本部门及以下数据权限 5:仅本人数据权限
	DeptIds   []uint64 `json:"deptIds" v:"required|slice_valid:uint64#部门ID列表不能为空" dc:"部门ID列表"`                                                              // 部门ID列表
}

// RoleSetDataScopeRes 设置数据权限返回参数
type RoleSetDataScopeRes struct {
}

// RoleDeleteReq 删除角色请求参数
type RoleDeleteReq struct {
	g.Meta `path:"/delete" tags:"RoleDelete" method:"delete" summary:"You first auth/role/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#角色ID列表不能为空" dc:"角色ID列表"` // 角色ID列表
}

// RoleDeleteRes 删除角色返回参数
type RoleDeleteRes struct {
}
