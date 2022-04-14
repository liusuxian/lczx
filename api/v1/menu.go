package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// MenuListReq 菜单树列表请求参数
type MenuListReq struct {
	g.Meta `path:"/list" tags:"MenuList" method:"get" summary:"You first auth/menu/list api"`
	Name   string `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#菜单名称必须为中文且长度不能超过20" dc:"菜单名称"` // 菜单名称
	Status string `json:"status" v:"in:0,1#菜单状态只能是0,1" dc:"菜单状态 0:停用 1:正常"`                    // 菜单状态 0:停用 1:正常
}

// MenuListRes 菜单树列表返回参数
type MenuListRes struct {
	List []*MenuTreeInfo `json:"list" dc:"菜单树列表"` // 菜单树列表
}

// MenuTreeInfo 菜单树信息
type MenuTreeInfo struct {
	Menu     *entity.Menu    `json:"menu" dc:"菜单信息"`        // 菜单信息
	Children []*MenuTreeInfo `json:"children" dc:"子菜单信息列表"` // 子菜单信息列表
}

// MenuIsMenusReq 获取菜单类型为目录和菜单的菜单树列表请求参数
type MenuIsMenusReq struct {
	g.Meta `path:"/isMenus" tags:"MenuIsMenus" method:"get" summary:"You first auth/menu/isMenus api"`
}

// MenuIsMenusRes 获取菜单类型为目录和菜单的菜单树列表返回参数
type MenuIsMenusRes struct {
	List []*MenuTreeInfo `json:"list" dc:"菜单树列表"` // 菜单树列表
}

// MenuAddReq 添加菜单请求参数
type MenuAddReq struct {
	g.Meta     `path:"/add" tags:"MenuAdd" method:"post" summary:"You first auth/menu/add api"`
	ParentId   uint64 `json:"parentId" v:"required|regex:^\\d*$#父规则ID不能为空|父规则ID必须为无符号整数" dc:"父规则ID"`                                           // 父规则ID
	Rule       string `json:"rule" v:"required|regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#权限规则不能为空|权限规则以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"权限规则"` // 权限规则
	Name       string `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#菜单名称不能为空|菜单名称必须为中文且长度不能超过20" dc:"菜单名称"`                           // 菜单名称
	Condition  string `json:"condition" v:"max-length:255#条件长度不能超过255" dc:"条件"`                                                                // 条件
	MenuType   uint   `json:"menuType" v:"required|in:0,1,2#类型不能为空|类型只能是0,1,2" dc:"类型 0:目录 1:菜单 2:按钮"`                                         // 类型 0:目录 1:菜单 2:按钮
	Status     uint   `json:"status" v:"required|in:0,1#菜单状态不能为空|菜单状态只能是0,1" dc:"菜单状态 0:停用 1:正常"`                                              // 菜单状态 0:停用 1:正常
	JumpPath   string `json:"jumpPath" v:"regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#跳转路由以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"跳转路由"`               // 跳转路由
	IsFrame    uint   `json:"isFrame" v:"required|in:0,1#是否外链不能为空|是否外链只能是0,1" dc:"是否外链 1是 0否"`                                                 // 是否外链 1是 0否
	ModuleType string `json:"moduleType" v:"required|max-length:30#所属模块不能为空|所属模块长度不能超过30" dc:"所属模块"`                                           // 所属模块
	Remark     string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                   // 备注
}

// MenuAddRes 添加菜单返回参数
type MenuAddRes struct {
}

// MenuEditReq 编辑菜单请求参数
type MenuEditReq struct {
	g.Meta     `path:"/edit" tags:"MenuEdit" method:"put" summary:"You first auth/menu/edit api"`
	Id         uint64 `json:"id" v:"required|regex:^[1-9]\\d*$#规则ID不能为空|规则ID必须为正整数" dc:"规则ID"`                                                 // 规则ID
	ParentId   uint64 `json:"parentId" v:"required|regex:^\\d*|different:Id$#父规则ID不能为空|父规则ID必须为无符号整数|父规则ID和规则ID不能相同" dc:"父规则ID"`               // 父规则ID
	Rule       string `json:"rule" v:"required|regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#权限规则不能为空|权限规则以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"权限规则"` // 权限规则
	Name       string `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#菜单名称不能为空|菜单名称必须为中文且长度不能超过20" dc:"菜单名称"`                           // 菜单名称
	Condition  string `json:"condition" v:"max-length:255#条件长度不能超过255" dc:"条件"`                                                                // 条件
	MenuType   uint   `json:"menuType" v:"required|in:0,1,2#类型不能为空|类型只能是0,1,2" dc:"类型 0:目录 1:菜单 2:按钮"`                                         // 类型 0:目录 1:菜单 2:按钮
	Status     uint   `json:"status" v:"required|in:0,1#菜单状态不能为空|菜单状态只能是0,1" dc:"菜单状态 0:停用 1:正常"`                                              // 菜单状态 0:停用 1:正常
	JumpPath   string `json:"jumpPath" v:"regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#跳转路由以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"跳转路由"`               // 跳转路由
	IsFrame    uint   `json:"isFrame" v:"required|in:0,1#是否外链不能为空|是否外链只能是0,1" dc:"是否外链 1是 0否"`                                                 // 是否外链 1是 0否
	ModuleType string `json:"moduleType" v:"required|max-length:30#所属模块不能为空|所属模块长度不能超过30" dc:"所属模块"`                                           // 所属模块
	Remark     string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                   // 备注
}

// MenuEditRes 编辑菜单返回参数
type MenuEditRes struct {
}

// MenuDeleteReq 删除菜单请求参数
type MenuDeleteReq struct {
	g.Meta `path:"/delete" tags:"MenuDelete" method:"delete" summary:"You first auth/menu/delete api"`
	Ids    []uint64 `json:"ids" v:"required|slice_valid:uint64#菜单ID列表不能为空" dc:"菜单ID列表"` // 菜单ID列表
}

// MenuDeleteRes 删除菜单返回参数
type MenuDeleteRes struct {
}

// MenuTreeReq 全部菜单树请求参数
type MenuTreeReq struct {
	g.Meta `path:"/tree" tags:"MenuTree" method:"get" summary:"You first auth/menu/tree api"`
}

// MenuTreeRes 全部菜单树返回参数
type MenuTreeRes struct {
	List []*MenuTreeInfo `json:"list" dc:"菜单树列表"` // 菜单树列表
}
