package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeGetUserInfoFailed    = gcode.New(1000, "获取用户信息失败", "") // 获取用户信息失败
	CodeCreateUserInfoFailed = gcode.New(1001, "创建用户信息失败", "") // 创建用户信息失败
	CodeGetDeptListFailed    = gcode.New(1002, "获取部门列表失败", "") // 获取部门列表失败
)
