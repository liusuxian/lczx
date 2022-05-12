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
	// WdkReportCfgKey 文档库报告类型配置缓存key
	WdkReportCfgKey = CachePrefix + "wdkreportcfg"
)
