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
	User     *entity.User `json:"user" dc:"用户信息"`       // 用户信息
	MenuList []string     `json:"menuList" dc:"用户菜单列表"` // 用户菜单列表
}

// UserProfileReq 获取个人中心请求参数
type UserProfileReq struct {
	g.Meta `path:"/profile" tags:"UserProfile" method:"get" summary:"You first user/profile api"`
}

// UserProfileRes 获取个人中心返回参数
type UserProfileRes struct {
	User  *entity.User   `json:"user" dc:"用户信息"`  // 用户信息
	Dept  *entity.Dept   `json:"dept" dc:"部门信息"`  // 部门信息
	Roles []*entity.Role `json:"roles" dc:"角色信息"` // 角色信息
}

// UserUploadAvatarReq 用户上传头像请求参数
type UserUploadAvatarReq struct {
	g.Meta     `path:"/uploadAvatar" tags:"UserUploadAvatar" method:"post" summary:"You first user/uploadAvatar api"`
	UploadName string `json:"uploadName" dc:"表单的文件字段名"` // 表单的文件字段名
}

type UserUploadAvatarRes struct {
	FileInfo *FileInfo `json:"fileInfo" dc:"文件信息"` // 文件信息
}

// UserProfileEditReq 编辑个人中心信息请求参数
type UserProfileEditReq struct {
	g.Meta   `path:"/profileEdit" tags:"UserProfileEdit" method:"put" summary:"You first user/profileEdit api"`
	Realname string `json:"realname" v:"required|regex:^[\u4e00-\u9fa5]{0,10}$#姓名不能为空|姓名必须为中文且长度不能超过10" dc:"姓名"` // 姓名
	Nickname string `json:"nickname" v:"regex:^[\u4e00-\u9fa5]{0,20}$#昵称必须为中文且长度不能超过20" dc:"昵称"`                 // 昵称
	Mobile   string `json:"mobile" v:"phone#不是有效的手机号码" dc:"手机号"`                                                 // 手机号
	Email    string `json:"email" v:"email#不是有效的用户邮箱" dc:"用户邮箱"`                                                 // 用户邮箱
	Gender   uint   `json:"gender" v:"required|in:1,2#性别不能为空|性别只能是1,2" dc:"性别 1: 男 2: 女"`                        // 性别 1: 男 2: 女
}

// UserProfileEditRes 编辑个人中心信息返回参数
type UserProfileEditRes struct {
}

// UserPwdEditReq 修改用户密码请求参数
type UserPwdEditReq struct {
	g.Meta      `path:"/pwdEdit" tags:"UserPwdEdit" method:"put" summary:"You first user/pwdEdit api"`
	OldPassword string `json:"oldPassword" v:"required|password#旧密码不能为空|旧密码为任意可见字符，长度在6~18之间" dc:"旧密码"` // 旧密码
	NewPassword string `json:"newPassword" v:"required|password#新密码不能为空|新密码为任意可见字符，长度在6~18之间" dc:"新密码"` // 新密码
}

// UserPwdEditRes 修改用户密码返回参数
type UserPwdEditRes struct {
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
