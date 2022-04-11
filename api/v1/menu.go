package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// MenuListReq 菜单树列表请求参数
type MenuListReq struct {
	g.Meta `path:"/list" tags:"MenuList" method:"get" summary:"You first auth/menu/list api"`
	Name   string `json:"name" v:"regex:^[\u4e00-\u9fa5]{0,20}$#菜单名称必须为纯中文且长度不能超过20" dc:"菜单名称"` // 菜单名称
	Status string `json:"status" v:"in:0,1#菜单状态只能是0,1" dc:"菜单状态 0:停用 1:正常"`                     // 菜单状态 0:停用 1:正常
}

// MenuListRes 菜单树列表返回参数
type MenuListRes struct {
	List []*MenuTreeInfo `json:"list" dc:"菜单树列表"` // 菜单树列表
}

// MenuTreeInfo 菜单树信息
type MenuTreeInfo struct {
	Menu     *entity.AuthRule `json:"menu" dc:"菜单信息"`        // 菜单信息
	Children []*MenuTreeInfo  `json:"children" dc:"子菜单信息列表"` // 子菜单信息列表
}

// MenuParentListReq 获取菜单类型为目录和菜单的菜单列表请求参数
type MenuParentListReq struct {
	g.Meta `path:"/parentList" tags:"MenuParentList" method:"get" summary:"You first auth/menu/parentList api"`
}

// MenuParentListRes 获取菜单类型为目录和菜单的菜单列表返回参数
type MenuParentListRes struct {
	List []*entity.AuthRule `json:"list" dc:"菜单列表"` // 菜单列表
}

// MenuAddReq 添加菜单请求参数
type MenuAddReq struct {
	ParentId   uint64 `json:"parentId" v:"required|regex:^\\d*$#父规则ID不能为空|父规则ID必须为非负整数" dc:"父规则ID"`                                                               // 父规则ID
	Rule       string `json:"rule" v:"required|regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#权限规则不能为空|权限规则以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"权限规则"`                    // 权限规则
	Name       string `json:"name" v:"required|regex:^[\u4e00-\u9fa5]{0,20}$#菜单名称不能为空|菜单名称必须为纯中文且长度不能超过20" dc:"菜单名称"`                                             // 菜单名称
	Condition  string `json:"condition" v:"max-length:255#条件长度不能超过255" dc:"条件"`                                                                                   // 条件
	MenuType   uint   `json:"menuType" v:"required|in:0,1,2#类型不能为空|类型只能是0,1,2" dc:"类型 0:目录 1:菜单 2:按钮"`                                                            // 类型 0:目录 1:菜单 2:按钮
	Status     uint   `json:"status" v:"required|in:0,1#菜单状态不能为空|菜单状态只能是0,1" dc:"菜单状态 0:停用 1:正常"`                                                                 // 菜单状态 0:停用 1:正常
	Show       uint   `json:"show" v:"required|in:0,1#显示状态不能为空|显示状态只能是0,1" dc:"显示状态 0:隐藏 1:显示"`                                                                   // 显示状态 0:隐藏 1:显示
	Path       string `json:"path" v:"required|regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#路由地址不能为空|路由地址以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"路由地址"`                    // 路由地址
	JumpPath   string `json:"jumpPath" v:"regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#跳转路由以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"跳转路由"`                                  // 跳转路由
	Component  string `json:"component" v:"required-if:menuType,1|regex:^[A-Za-z][A-Za-z0-9/_]{0,100}$#组件路径不能为空|组件路径以字母开头，只能包含字母、数字、下划线和反斜杠且长度不能超过100" dc:"组件路径"` // 组件路径
	IsFrame    uint   `json:"isFrame" v:"required|in:0,1#是否外链不能为空|是否外链只能是0,1" dc:"是否外链 1是 0否"`                                                                    // 是否外链 1是 0否
	ModuleType string `json:"moduleType" v:"required|max-length:30#所属模块不能为空|所属模块长度不能超过30" dc:"所属模块"`                                                              // 所属模块
	ModelId    uint   `json:"modelId" v:"regex:^\\d*$#模型ID必须为非负整数" dc:"模型ID"`                                                                                     // 模型ID
	Remark     string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                                                                      // 备注
}

// MenuAddRes 添加菜单返回参数
type MenuAddRes struct {
}
