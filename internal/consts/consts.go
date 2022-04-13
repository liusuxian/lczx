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

// 菜单类型
const (
	MenuTypeDir    = 0 // 目录
	MenuTypeMenu   = 1 // 菜单
	MenuTypeButton = 2 // 按钮
)
