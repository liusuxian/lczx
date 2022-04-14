package code

import "github.com/gogf/gf/v2/errors/gcode"

var (
	GetCaptchaImgFailed = gcode.New(1000, "获取验证码图片信息失败", "") // 获取验证码图片信息失败
	CaptchaPutErr       = gcode.New(1001, "验证码输入错误", "")     // 验证码输入错误
	GetAccessAuthFailed = gcode.New(1002, "获取权限失败", "")      // 获取权限失败
	NotAccessAuth       = gcode.New(1003, "没有访问权限", "")      // 没有访问权限

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

	GetDeptListFailed = gcode.New(1300, "获取部门列表失败", "")  // 获取部门列表失败
	GetDeptTreeFailed = gcode.New(1301, "获取部门树信息失败", "") // 获取部门树信息失败
	AddDeptFailed     = gcode.New(1302, "新增部门信息失败", "")  // 新增部门信息失败
	GetDeptFailed     = gcode.New(1303, "获取部门信息失败", "")  // 获取部门信息失败
	EditDeptFailed    = gcode.New(1304, "修改部门信息失败", "")  // 修改部门信息失败
	DeleteDeptFailed  = gcode.New(1305, "删除部门信息失败", "")  // 删除部门信息失败

	GetMenuListFailed = gcode.New(1400, "获取菜单列表失败", "")            // 获取菜单列表失败
	GetIsMenusFailed  = gcode.New(1401, "获取菜单类型为目录和菜单的菜单列表失败", "") // 获取菜单类型为目录和菜单的菜单列表失败
	AddMenuFailed     = gcode.New(1402, "新增菜单信息失败", "")            // 新增菜单信息失败
	EditMenuFailed    = gcode.New(1403, "修改菜单信息失败", "")            // 修改菜单信息失败
	DeleteMenuFailed  = gcode.New(1305, "删除菜单信息失败", "")            // 删除菜单信息失败
)
