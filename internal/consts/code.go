package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeGetUserFailed      = gcode.New(1000, "获取用户信息失败", "")   // 获取用户信息失败
	CodeUserNotExist       = gcode.New(1001, "用户信息不存在", "")    // 用户信息不存在
	CodeUserStatusDisabled = gcode.New(1002, "用户已被禁用", "")     // 用户已被禁用
	CodeGetUserListFailed  = gcode.New(1003, "获取用户列表失败", "")   // 获取用户列表失败
	CodeAddUserFailed      = gcode.New(1004, "新增用户信息失败", "")   // 新增用户信息失败
	CodeDeleteUserFailed   = gcode.New(1005, "删除用户信息失败", "")   // 删除用户信息失败
	CodeDeleteUserNotExist = gcode.New(1006, "删除的用户ID不存在", "") // 删除的用户ID不存在
	CodeUpdateUserFailed   = gcode.New(1007, "修改用户信息失败", "")   // 修改用户信息失败
	CodeUpdateUserNotExist = gcode.New(1008, "修改的用户ID不存在", "") // 修改的用户ID不存在
	CodeGetDeptListFailed  = gcode.New(1100, "获取部门列表失败", "")   // 获取部门列表失败
	CodeAddDeptFailed      = gcode.New(1101, "新增部门信息失败", "")   // 新增部门信息失败
	CodeDeleteDeptFailed   = gcode.New(1102, "删除部门信息失败", "")   // 删除部门信息失败
	CodeDeleteDeptNotExist = gcode.New(1103, "删除的部门ID不存在", "") // 删除的部门ID不存在
	CodeUpdateDeptFailed   = gcode.New(1104, "修改部门信息失败", "")   // 修改部门信息失败
	CodeUpdateDeptNotExist = gcode.New(1105, "修改的部门ID不存在", "") // 修改的部门ID不存在
)
