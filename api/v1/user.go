package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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

// UserProfileReq 获取个人中心信息请求参数
type UserProfileReq struct {
	g.Meta `path:"/profile" tags:"UserProfile" method:"get" summary:"You first user/profile api"`
}

// UserProfileRes 获取个人中心信息返回参数
type UserProfileRes struct {
	ProfileInfo *UserProfileInfo `json:"profileInfo" dc:"个人中心信息"` // 个人中心信息
}

// UserProfileInfo 个人中心信息
type UserProfileInfo struct {
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

// UserListReq 用户列表请求参数
type UserListReq struct {
	g.Meta    `path:"/list" tags:"UserList" method:"get" summary:"You first auth/user/list api"`
	DeptId    string      `json:"deptId" v:"regex:^[1-9]\\d*$#部门ID必须为正整数" dc:"部门ID"`                             // 部门ID
	Passport  string      `json:"passport" v:"regex:^[a-zA-Z]\\w{0,18}$#账号以字母开头，只能包含字母、数字和下划线且长度不能超过18" dc:"账号"` // 账号
	Realname  string      `json:"realname" v:"regex:^[\u4e00-\u9fa5]{0,10}$#姓名必须为中文且长度不能超过10" dc:"姓名"`           // 姓名
	Mobile    string      `json:"mobile" v:"regex:^\\d{0,11}$#手机号必须为无符号整数且长度不能超过11" dc:"手机号"`                    // 手机号
	Status    string      `json:"status" v:"in:0,1#状态只能是0,1" dc:"状态 0:禁用 1:启用"`                                  // 状态 0:禁用 1:启用
	StartTime *gtime.Time `json:"startTime" v:"datetime#开始时间不是有效的日期时间" dc:"开始时间"`                                // 开始时间
	EndTime   *gtime.Time `json:"endTime" v:"datetime#结束时间不是有效的日期时间" dc:"结束时间"`                                  // 结束时间
	SortName  string      `json:"sortName" v:"regex:^[a-zA-Z]\\w*$#排序字段以字母开头，只能包含字母、数字和下划线" dc:"排序字段"`           // 排序字段
	SortOrder string      `json:"sortOrder" v:"regex:^[a-zA-Z]\\w*$#排序方式以字母开头，只能包含字母、数字和下划线" dc:"排序方式"`          // 排序方式
	CurPage   int         `json:"curPage" v:"required|regex:^[1-9]\\d*$#当前页码不能为空|当前页码必须为正整数" dc:"当前页码"`          // 当前页码
	PageSize  int         `json:"pageSize" v:"required|regex:^[1-9]\\d*$#每页数量不能为空|每页数量必须为正整数" dc:"每页数量"`         // 每页数量
}

// UserListRes 用户列表返回参数
type UserListRes struct {
	CurPage         int                `json:"curPage" dc:"当前页码"`             // 当前页码
	Total           int                `json:"total" dc:"数据总量"`               // 数据总量
	ProfileInfoList []*UserProfileInfo `json:"profileInfoList" dc:"个人中心信息列表"` // 个人中心信息列表
}

// UserAddReq 添加用户请求参数
type UserAddReq struct {
	g.Meta   `path:"/add" tags:"UserAdd" method:"post" summary:"You first user/add api"`
	Passport string   `json:"passport" v:"required|passport#账号不能为空|账号以字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"账号"`       // 账号
	Password string   `json:"password" v:"required|password#密码不能为空|密码为任意可见字符，长度在6~18之间" dc:"密码"`                   // 密码
	Realname string   `json:"realname" v:"required|regex:^[\u4e00-\u9fa5]{0,10}$#姓名不能为空|姓名必须为中文且长度不能超过10" dc:"姓名"` // 姓名
	Nickname string   `json:"nickname" v:"regex:^[\u4e00-\u9fa5]{0,20}$#昵称必须为中文且长度不能超过20" dc:"昵称"`                 // 昵称
	DeptId   uint64   `json:"deptId" v:"required|regex:^[1-9]\\d*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"`                 // 部门ID
	RoleIds  []uint64 `json:"roleIds" v:"required|slice_valid:uint64#角色ID列表不能为空" dc:"角色ID列表"`                      // 角色ID列表
	Gender   uint     `json:"gender" v:"required|in:1,2#性别不能为空|性别只能是1,2" dc:"性别 1: 男 2: 女"`                        // 性别 1: 男 2: 女
	Status   uint     `json:"status" v:"required|in:0,1#状态不能为空|状态只能是0,1" dc:"状态 0:禁用 1:启用"`                        // 状态 0:禁用 1:启用
	IsAdmin  uint     `json:"isAdmin" v:"required|in:0,1#是否后台管理员不能为空|是否后台管理员只能是0,1" dc:"是否后台管理员 0:否 1:是"`          // 是否后台管理员 0:否 1:是
	Mobile   string   `json:"mobile" v:"required|phone#手机号不能为空|不是有效的手机号" dc:"手机号"`                                 // 手机号
	Email    string   `json:"email" v:"email#不是有效的用户邮箱" dc:"用户邮箱"`                                                 // 用户邮箱
	Remark   string   `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                       // 备注
}

// UserAddRes 添加用户返回参数
type UserAddRes struct {
}
