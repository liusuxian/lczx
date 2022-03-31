package code

import "github.com/gogf/gf/v2/errors/gcode"

var (
	GetCaptchaDebugFailed = gcode.New(1000, "获取验证码调试配置失败", "") // 获取验证码调试配置失败
	GetCaptchaImgFailed   = gcode.New(1001, "获取验证码图片信息失败", "") // 获取验证码图片信息失败
	CaptchaPutErr         = gcode.New(1002, "验证码输入错误", "")     // 验证码输入错误

	GetUserOnlineListFailed     = gcode.New(1100, "获取在线用户列表失败", "") // 获取在线用户列表失败
	ForceLogoutUserOnlineFailed = gcode.New(1101, "强退在线用户失败", "")   // 强退在线用户失败
	GetLoginLogListFailed       = gcode.New(1102, "获取登录日志列表失败", "") // 获取登录日志列表失败
	DeleteLoginLogFailed        = gcode.New(1103, "删除登录日志失败", "")   // 删除登录日志失败
	ClearLoginLogFailed         = gcode.New(1104, "清除登录日志失败", "")   // 清除登录日志失败

	GetUserFailed     = gcode.New(1200, "获取用户信息失败", "") // 获取用户信息失败
	GetUserListFailed = gcode.New(1201, "获取用户列表失败", "") // 获取用户列表失败
	AddUserFailed     = gcode.New(1202, "新增用户信息失败", "") // 新增用户信息失败
	DeleteUserFailed  = gcode.New(1203, "删除用户信息失败", "") // 删除用户信息失败
	UpdateUserFailed  = gcode.New(1204, "修改用户信息失败", "") // 修改用户信息失败

	GetDeptListFailed = gcode.New(1300, "获取部门列表失败", "") // 获取部门列表失败
	AddDeptFailed     = gcode.New(1301, "新增部门信息失败", "") // 新增部门信息失败
	DeleteDeptFailed  = gcode.New(1302, "删除部门信息失败", "") // 删除部门信息失败
	UpdateDeptFailed  = gcode.New(1303, "修改部门信息失败", "") // 修改部门信息失败

)
