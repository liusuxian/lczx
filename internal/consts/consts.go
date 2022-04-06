package consts

const (
	ContextKey          = "ContextKey" // 上下文变量存储键名，前后端系统共享
	SessionKey          = "SessionKey" // Session Key
	ContextKeyRequestId = "RequestId"  // 用于自定义Context RequestId打印
)

// 用户状态
const (
	UserStatusDisabled = 0 // 禁用
	UserStatusEnable   = 1 // 启用
)

// 用户登录状态
const (
	LoginFailed = 0 // 失败
	LoginSucc   = 1 // 成功
)

// 部门状态
const (
	DeptStatusDisabled = 0 // 停用
	DeptStatusEnable   = 1 // 正常
)

// 角色
const (
	RoleUser       = 0    // 默认普通用户
	RoleAdmin      = 900  // 普通管理员
	RoleSuperAdmin = 1000 // 超级管理员
)
