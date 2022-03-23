package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeGetUserFailed     = gcode.New(1000, "获取用户信息失败", "") // 获取用户信息失败
	CodeUserNotExist      = gcode.New(1001, "用户信息不存在", "")  // 用户信息不存在
	CodeAddUserFailed     = gcode.New(1002, "新增用户信息失败", "") // 新增用户信息失败
	CodeGetDeptListFailed = gcode.New(1003, "获取部门列表失败", "") // 获取部门列表失败
	CodeAddDeptFailed     = gcode.New(1004, "新增部门信息失败", "") // 新增部门信息失败
	CodeDeleteDeptFailed  = gcode.New(1005, "删除部门信息失败", "") // 删除部门信息失败
)
