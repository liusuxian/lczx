package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

// UserInfoReq 获取用户信息请求参数
type UserInfoReq struct {
	g.Meta `path:"/info" tags:"UserInfo" method:"post" summary:"You first user/info api"`
}

// UserInfoRes 获取用户信息返回参数
type UserInfoRes struct {
	User         *entity.User `json:"user" dc:"用户信息"`           // 用户信息
	RoleNameList []string     `json:"roleNameList" dc:"角色名称列表"` // 用户角色名称列表
	MenuList     []string     `json:"menuList" dc:"用户菜单列表"`     // 用户菜单列表
}

// UserProfileReq 获取个人中心请求参数
type UserProfileReq struct {
	g.Meta `path:"/profile" tags:"UserProfile" method:"get" summary:"You first user/profile api"`
}

// UserProfileRes 获取个人中心返回参数
type UserProfileRes struct {
	User *entity.User `json:"user" dc:"用户信息"` // 用户信息
	Dept *entity.Dept `json:"dept" dc:"部门信息"` // 部门信息
	Role *entity.Role `json:"role" dc:"角色信息"` // 角色信息
}

type UserAddReq struct {
	g.Meta   `path:"/add" tags:"UserAdd" method:"post" summary:"You first user/add api"`
	Passport string `json:"passport" v:"required|passport#账号不能为空|账号以字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"账号"`       // 账号
	Realname string `json:"realname" v:"required|regex:^[\u4e00-\u9fa5]{0,10}$#姓名不能为空|姓名必须为中文且长度不能超过10" dc:"姓名"` // 姓名
	Gender   uint   `json:"gender" v:"required|in:1,2#性别不能为空|性别只能是1,2" dc:"性别 1: 男 2: 女"`                        // 性别 1: 男 2: 女
	Mobile   string `json:"mobile" v:"phone#不是有效的手机号码" dc:"手机号"`                                                 // 手机号
	DeptId   uint64 `json:"deptId" v:"required|regex:^[1-9]\\d*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"`                 // 部门ID
}
type UserAddRes struct {
	User *entity.User `json:"user" dc:"用户信息"` // 用户信息
}
