package code

import "github.com/gogf/gf/v2/errors/gcode"

var (
	GetCaptchaOnOffCfgFailed = gcode.New(1000, "获取验证码开关配置失败", "") // 获取验证码开关配置失败
	GetCaptchaFailed         = gcode.New(1001, "获取验证码失败", "")     // 获取验证码失败
	CaptchaPutErr            = gcode.New(1002, "验证码输入错误", "")     // 验证码输入错误
	GetUserFailed            = gcode.New(2000, "获取用户信息失败", "")    // 获取用户信息失败
	GetUserListFailed        = gcode.New(2001, "获取用户列表失败", "")    // 获取用户列表失败
	AddUserFailed            = gcode.New(2002, "新增用户信息失败", "")    // 新增用户信息失败
	DeleteUserFailed         = gcode.New(2003, "删除用户信息失败", "")    // 删除用户信息失败
	UpdateUserFailed         = gcode.New(2004, "修改用户信息失败", "")    // 修改用户信息失败
	GetDeptListFailed        = gcode.New(2100, "获取部门列表失败", "")    // 获取部门列表失败
	AddDeptFailed            = gcode.New(2101, "新增部门信息失败", "")    // 新增部门信息失败
	DeleteDeptFailed         = gcode.New(2102, "删除部门信息失败", "")    // 删除部门信息失败
	UpdateDeptFailed         = gcode.New(2103, "修改部门信息失败", "")    // 修改部门信息失败
	GetUserOnlineListFailed  = gcode.New(2200, "获取在线用户列表失败", "")  // 获取在线用户列表失败
)
