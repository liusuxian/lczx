package utils

const (
	// cachePrefix 缓存前缀
	cachePrefix = "cache:lczx:"
	// menu 缓存菜单key
	menu = cachePrefix + "menu"
	// role 缓存角色key
	role = cachePrefix + "role"
)

// GetCacheMenuKey 获取缓存菜单key
func GetCacheMenuKey() string {
	return menu
}

// GetCacheRoleKey 获取缓存角色key
func GetCacheRoleKey() string {
	return role
}
