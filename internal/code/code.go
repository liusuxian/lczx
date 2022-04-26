package code

import "github.com/gogf/gf/v2/errors/gcode"

var (
	GetCaptchaImgFailed = gcode.New(1000, "获取验证码图片信息失败", "") // 获取验证码图片信息失败
	CaptchaPutErr       = gcode.New(1001, "验证码输入错误", "")     // 验证码输入错误
	RequestDataFailed   = gcode.New(1002, "请求数据失败", "")      // 请求数据失败
	GetAccessAuthFailed = gcode.New(1003, "获取权限失败", "")      // 获取权限失败
	NotAccessAuth       = gcode.New(1004, "没有访问权限", "")      // 没有访问权限

	GetUserOnlineListFailed     = gcode.New(1100, "获取在线用户列表失败", "") // 获取在线用户列表失败
	ForceLogoutUserOnlineFailed = gcode.New(1101, "强退在线用户失败", "")   // 强退在线用户失败
	GetLoginLogListFailed       = gcode.New(1102, "获取登录日志列表失败", "") // 获取登录日志列表失败
	DeleteLoginLogFailed        = gcode.New(1103, "删除登录日志失败", "")   // 删除登录日志失败
	ClearLoginLogFailed         = gcode.New(1104, "清除登录日志失败", "")   // 清除登录日志失败
	GetOperLogListFailed        = gcode.New(1105, "获取操作日志列表失败", "") // 获取操作日志列表失败
	DeleteOperLogFailed         = gcode.New(1106, "删除操作日志失败", "")   // 删除操作日志失败
	ClearOperLogFailed          = gcode.New(1107, "清除操作日志失败", "")   // 清除操作日志失败

	GetUserFailed         = gcode.New(1200, "获取用户信息失败", "")     // 获取用户信息失败
	GetUserProfileFailed  = gcode.New(1201, "获取用户个人中心信息失败", "") // 获取用户个人中心信息失败
	EditUserProfileFailed = gcode.New(1202, "修改用户个人中心信息失败", "") // 修改用户个人中心信息失败
	GetUserListFailed     = gcode.New(1203, "获取用户列表失败", "")     // 获取用户列表失败
	AddUserFailed         = gcode.New(1204, "添加用户信息失败", "")     // 添加用户信息失败
	DeleteUserFailed      = gcode.New(1205, "删除用户信息失败", "")     // 删除用户信息失败
	EditUserFailed        = gcode.New(1206, "修改用户信息失败", "")     // 修改用户信息失败
	EditPwdFailed         = gcode.New(1207, "修改用户密码失败", "")     // 修改用户密码失败
	ResetPwdFailed        = gcode.New(1208, "重置用户密码失败", "")     // 重置用户密码失败
	GetEditUserFailed     = gcode.New(1209, "获取被修改用户的信息失败", "") // 获取被修改用户的信息失败
	SetUserStatusFailed   = gcode.New(1210, "设置用户状态失败", "")     // 设置用户状态失败
	SetUserAvatarFailed   = gcode.New(1211, "设置用户头像失败", "")     // 设置用户头像失败

	GetDeptListFailed     = gcode.New(1300, "获取部门列表失败", "")    // 获取部门列表失败
	GetDeptTreeFailed     = gcode.New(1301, "获取部门树信息失败", "")   // 获取部门树信息失败
	GetRoleDeptTreeFailed = gcode.New(1302, "获取角色部门树信息失败", "") // 获取角色部门树信息失败
	AddDeptFailed         = gcode.New(1303, "添加部门信息失败", "")    // 添加部门信息失败
	GetDeptFailed         = gcode.New(1304, "获取部门信息失败", "")    // 获取部门信息失败
	EditDeptFailed        = gcode.New(1305, "修改部门信息失败", "")    // 修改部门信息失败
	DeleteDeptFailed      = gcode.New(1306, "删除部门信息失败", "")    // 删除部门信息失败

	GetMenuListFailed  = gcode.New(1400, "获取菜单列表失败", "")            // 获取菜单列表失败
	GetIsMenusFailed   = gcode.New(1401, "获取菜单类型为目录和菜单的菜单列表失败", "") // 获取菜单类型为目录和菜单的菜单列表失败
	AddMenuFailed      = gcode.New(1402, "添加菜单信息失败", "")            // 添加菜单信息失败
	EditMenuFailed     = gcode.New(1403, "修改菜单信息失败", "")            // 修改菜单信息失败
	DeleteMenuFailed   = gcode.New(1404, "删除菜单信息失败", "")            // 删除菜单信息失败
	GetMenuTreeFailed  = gcode.New(1405, "获取菜单树信息失败", "")           // 获取菜单树信息失败
	MenuStatusDisabled = gcode.New(1406, "菜单已停用", "")               // 菜单已停用

	GetRoleListFailed      = gcode.New(1500, "获取角色列表失败", "")     // 获取角色列表失败
	AddRoleFailed          = gcode.New(1501, "添加角色信息失败", "")     // 添加角色信息失败
	GetRoleFailed          = gcode.New(1502, "获取角色信息失败", "")     // 获取角色信息失败
	EditRoleFailed         = gcode.New(1503, "修改角色信息失败", "")     // 修改角色信息失败
	SetRoleStatusFailed    = gcode.New(1504, "设置角色状态失败", "")     // 设置角色状态失败
	SetRoleDataScopeFailed = gcode.New(1505, "设置数据权限失败", "")     // 设置数据权限失败
	DeleteRoleFailed       = gcode.New(1506, "删除角色信息失败", "")     // 删除角色信息失败
	GetEnableRolesFailed   = gcode.New(1507, "获取所有可用角色列表失败", "") // 获取所有可用角色列表失败

	GetWdkProjectListFailed = gcode.New(1600, "获取文档库项目列表失败", "") // 获取文档库项目列表失败
	AddWdkProjectFailed     = gcode.New(1601, "添加文档库项目失败", "")   // 添加文档库项目失败
	GetWdkProjectFailed     = gcode.New(1602, "获取文档库项目信息失败", "") // 获取文档库项目信息失败
	EditWdkProjectFailed    = gcode.New(1603, "修改文档库项目失败", "")   // 修改文档库项目失败
	DeleteWdkProjectFailed  = gcode.New(1604, "删除文档库项目失败", "")   // 删除文档库项目失败

	GetWdkReportCfgListFailed = gcode.New(1700, "获取文档库报告类型配置列表失败", "") // 获取文档库报告类型配置列表失败
	AddWdkReportCfgFailed     = gcode.New(1701, "添加文档库报告类型配置失败", "")   // 添加文档库报告类型配置失败
	GetWdkReportCfgFailed     = gcode.New(1702, "获取文档库报告类型配置信息失败", "") // 获取文档库报告类型配置信息失败
	EditWdkReportCfgFailed    = gcode.New(1703, "修改文档库报告类型配置失败", "")   // 修改文档库报告类型配置失败
	DeleteWdkReportCfgFailed  = gcode.New(1704, "删除文档库报告类型配置失败", "")   // 删除文档库报告类型配置失败

	GetWdkAttachmentRecordFailed = gcode.New(1800, "获取文档库上传附件记录失败", "") // 获取文档库上传附件记录失败
	AddWdkAttachmentRecordFailed = gcode.New(1801, "添加文档库上传附件记录失败", "") // 添加文档库上传附件记录失败

	GetWdkServiceRecordFailed = gcode.New(1900, "获取文档库服务记录失败", "") // 获取文档库服务记录失败
	AddWdkServiceRecordFailed = gcode.New(1901, "添加文档库服务记录失败", "") // 添加文档库服务记录失败
)
