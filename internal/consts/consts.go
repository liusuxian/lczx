package consts

const (
	ContextKey          = "ContextKey" // 上下文变量存储键名，前后端系统共享
	SessionKey          = "SessionKey" // Session Key
	ContextKeyRequestId = "RequestId"  // 用于自定义Context RequestId打印
)

const (
	// CachePrefix 缓存前缀
	CachePrefix = "cache:lczx:"
	// CacheAccessUrlPrefix 访问授权URL缓存前缀
	CacheAccessUrlPrefix = "cache:lczx:accessurl:"
	// DeptKey 缓存部门key
	DeptKey = CachePrefix + "dept"
	// MenuKey 缓存菜单key
	MenuKey = CachePrefix + "menu"
	// RoleKey 缓存角色key
	RoleKey = CachePrefix + "role"
)
