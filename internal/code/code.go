package code

import "github.com/gogf/gf/v2/errors/gcode"

var (
	GetUserFailed     = gcode.New(1000, "获取用户信息失败", "") // 获取用户信息失败
	UserNotExist      = gcode.New(1001, "用户信息不存在", "")  // 用户信息不存在
	UserDisabled      = gcode.New(1002, "用户已被禁用", "")   // 用户已被禁用
	GetUserListFailed = gcode.New(1003, "获取用户列表失败", "") // 获取用户列表失败
	AddUserFailed     = gcode.New(1004, "新增用户信息失败", "") // 新增用户信息失败
	DeleteUserFailed  = gcode.New(1005, "删除用户信息失败", "") // 删除用户信息失败
	UpdateUserFailed  = gcode.New(1006, "修改用户信息失败", "") // 修改用户信息失败
	GetDeptListFailed = gcode.New(1100, "获取部门列表失败", "") // 获取部门列表失败
	AddDeptFailed     = gcode.New(1101, "新增部门信息失败", "") // 新增部门信息失败
	DeleteDeptFailed  = gcode.New(1102, "删除部门信息失败", "") // 删除部门信息失败
	UpdateDeptFailed  = gcode.New(1103, "修改部门信息失败", "") // 修改部门信息失败
)
