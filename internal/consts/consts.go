package consts

const (
	ContextKey          = "ContextKey" // 上下文变量存储键名，前后端系统共享
	SessionKey          = "SessionKey" // Session Key
	ContextKeyRequestId = "RequestId"  // 用于自定义Context RequestId打印
)

// 用户状态
const (
	UserStatusEnable   = 0 // 启用
	UserStatusDisabled = 1 // 禁用
)

// 角色
const (
	RoleUser       = 0    // 默认普通用户
	RoleAdmin      = 900  // 普通管理员
	RoleSuperAdmin = 1000 // 超级管理员
)
