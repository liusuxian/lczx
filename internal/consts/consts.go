package consts

const (
	ContextKey          = "ContextKey" // 上下文变量存储键名，前后端系统共享
	SessionKey          = "SessionKey" // Session Key
	ContextKeyRequestId = "RequestId"  // 用于自定义Context RequestId打印
)

const (
	// CachePrefix 缓存前缀
	CachePrefix = "cache:lczx:"
	// DeptKey 缓存部门key
	DeptKey = CachePrefix + "dept"
	// MenuKey 缓存菜单key
	MenuKey = CachePrefix + "menu"
	// RoleKey 缓存角色key
	RoleKey = CachePrefix + "role"
	// WdkProjectKey 缓存角色key
	WdkProjectKey = CachePrefix + "wdkproject"
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

// 角色状态
const (
	RoleStatusDisabled = 0 // 停用
	RoleStatusEnable   = 1 // 正常
)

// 菜单状态
const (
	MenuStatusDisabled = 0 // 停用
	MenuStatusEnable   = 1 // 正常
)

// 菜单类型
const (
	MenuTypeDir    = 0 // 目录
	MenuTypeMenu   = 1 // 菜单
	MenuTypeButton = 2 // 按钮
)

// 数据范围
const (
	DataScopeAll          = 1 // 全部数据权限
	DataScopeCustom       = 2 // 自定义数据权限
	DataScopeDept         = 3 // 本部门数据权限
	DataScopeDeptAndBelow = 4 // 本部门及以下数据权限
	DataScopeSelf         = 5 // 仅本人数据权限
)

// 文档库项目阶段
const (
	WdkPStepNotStart      = 0 // 未开始
	WdkPStepContractAward = 1 // 合同签约
	WdkPStepProjectStart  = 2 // 项目启动会
	WdkPStepInService     = 3 // 服务中
	WdkPStepContractStop  = 4 // 合同结束
	WdkPStepReview        = 5 // 复盘
)

// 文档库项目文件上传状态
const (
	WdkPFileUpStatusNotComplete = 0 // 未传完
	WdkPFileUpStatusComplete    = 1 // 已传完
)

// 文档库项目服务状态
const (
	WdkPSStatusInService = 0 // 服务中
	WdkPSStatusPause     = 1 // 暂停
	WdkPSStatusStop      = 2 // 提前终止
	WdkPSStatusComplete  = 3 // 正常结束
)
