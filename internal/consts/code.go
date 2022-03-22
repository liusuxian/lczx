package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeErrorReqArgs      = gcode.New(1000, "请求参数错误", "")   // 请求参数错误
	CodeGetUserInfoFailed = gcode.New(1001, "获取用户信息失败", "") // 获取用户信息失败
)
