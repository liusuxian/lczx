package service

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"time"
)

type sRole struct{}

var (
	insRole = sRole{}
)

// Role 角色管理服务
func Role() *sRole {
	return &insRole
}

// GetRoleList 获取角色列表
func (s *sRole) GetRoleList(ctx context.Context, req *v1.RoleListReq) (total int, list []*entity.Role, err error) {
	model := dao.Role.Ctx(ctx)
	columns := dao.Role.Columns()
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if req.StartTime.String() != "" {
		model = model.WhereGTE(columns.CreateAt, req.StartTime)
	}
	if req.EndTime.String() != "" {
		model = model.WhereLTE(columns.CreateAt, req.EndTime)
	}
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderAsc(columns.Id).Scan(&list)
	return
}

// AddRole 新增角色
func (s *sRole) AddRole(ctx context.Context, req *v1.RoleAddReq) (err error) {
	// 检查角色名称是否可用
	var available bool
	available, err = s.IsRoleNameAvailable(ctx, req.Name)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`角色名称[%s]已存在`, req.Name)
		return
	}
	// 开启事务
	err = dao.Role.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 保存角色
		roleId, err := s.SaveRole(ctx, req)
		if err != nil {
			return err
		}
		// 添加角色权限
		err = s.AddRoleRule(ctx, req.MenuIds, roleId)
		return err
	})
	if err != nil {
		return
	}
	// 清除角色缓存
	_, err = Cache().ClearCache(ctx, consts.RoleKey)
	return
}

// EditRole 编辑角色
func (s *sRole) EditRole(ctx context.Context, req *v1.RoleEditReq) (err error) {
	// 检查角色信息是否存在
	var role *entity.Role
	role, err = s.GetRoleById(ctx, req.Id)
	if err != nil {
		return
	}
	if role == nil {
		err = gerror.Newf(`角色ID[%d]不存在`, req.Id)
		return
	}
	// 检查角色名称是否可用
	if role.Name != req.Name {
		var available bool
		available, err = s.IsRoleNameAvailable(ctx, req.Name)
		if err != nil {
			return
		}
		if !available {
			err = gerror.Newf(`角色名称[%s]已存在`, req.Name)
			return
		}
	}
	// 开启事务
	err = dao.Role.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 更新角色
		err := s.UpdateRole(ctx, req)
		if err != nil {
			return err
		}
		// 修改角色权限
		err = s.EditRoleRule(ctx, req.MenuIds, req.Id)
		return err
	})
	if err != nil {
		return
	}
	// 清除角色缓存
	_, err = Cache().ClearCache(ctx, consts.RoleKey)
	return
}

// SetRoleStatus 设置角色状态
func (s *sRole) SetRoleStatus(ctx context.Context, req *v1.RoleSetStatusReq) (err error) {
	_, err = dao.Role.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.RoleKey,
		Force:    false,
	}).Data(do.Role{Status: req.Status}).Where(do.Role{Id: req.Id}).Update()
	return
}

// SetRoleDataScope 设置数据权限
func (s *sRole) SetRoleDataScope(ctx context.Context, req *v1.RoleSetDataScopeReq) (err error) {
	err = dao.Role.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 更新数据权限
		_, err := dao.Role.Ctx(ctx).Data(do.Role{DataScope: req.DataScope}).Where(do.Role{Id: req.Id}).Update()
		if err != nil {
			return err
		}
		// 处理自定义数据权限
		if req.DataScope == consts.DataScopeCustom {
			_, err = dao.RoleDept.Ctx(ctx).Where(do.RoleDept{RoleId: req.Id}).Delete()
			if err != nil {
				return err
			}
			data := g.List{}
			for _, deptId := range req.DeptIds {
				data = append(data, g.Map{"role_id": req.Id, "dept_id": deptId})
			}
			_, err = dao.RoleDept.Ctx(ctx).Data(data).Insert()
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	// 清除角色缓存
	_, err = Cache().ClearCache(ctx, consts.RoleKey)
	return
}

// DeleteRoleByIds 通过角色ID列表删除角色
func (s *sRole) DeleteRoleByIds(ctx context.Context, ids []uint64) (err error) {
	err = dao.Role.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除角色
		_, err := dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, ids).Delete()
		if err != nil {
			return err
		}
		// 删除角色的权限和管理的部门数据权限
		var enforcer *casbin.SyncedEnforcer
		enforcer, err = Casbin(ctx).GetEnforcer()
		if err != nil {
			return err
		}
		// 删除角色的权限
		for _, id := range ids {
			_, err = enforcer.RemoveFilteredPolicy(0, gconv.String(id))
			if err != nil {
				return err
			}
		}
		// 删除管理的部门数据权限
		_, err = dao.RoleDept.Ctx(ctx).WhereIn(dao.RoleDept.Columns().RoleId, ids).Delete()
		return err
	})
	if err != nil {
		return
	}
	// 清除角色缓存
	_, err = Cache().ClearCache(ctx, consts.RoleKey)
	return
}

// GetEnableRoles 获取全部可用的角色
func (s *sRole) GetEnableRoles(ctx context.Context) (roles []*entity.Role, err error) {
	// 获取所有角色
	var allRoles []*entity.Role
	allRoles, err = s.GetAllRoles(ctx)
	if err != nil {
		return
	}

	roles = make([]*entity.Role, 0, len(allRoles))
	for _, r := range allRoles {
		if r.Status == consts.RoleStatusEnable {
			roles = append(roles, r)
		}
	}
	return
}

// GetAllRoles 获取所有角色
func (s *sRole) GetAllRoles(ctx context.Context) (roles []*entity.Role, err error) {
	// 从缓存获取
	rolesCacheVal := Cache().GetCache(ctx, consts.RoleKey)
	if rolesCacheVal != nil {
		err = gconv.Structs(rolesCacheVal, &roles)
		if err != nil {
			return
		}
		if roles != nil {
			return
		}
	}
	// 从数据库获取
	err = dao.Role.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 2,
		Name:     consts.RoleKey,
		Force:    false,
	}).OrderAsc(dao.Role.Columns().Id).Scan(&roles)
	return
}

// GetRoleById 通过角色ID获取角色信息
func (s *sRole) GetRoleById(ctx context.Context, id uint64) (role *entity.Role, err error) {
	err = dao.Role.Ctx(ctx).Where(do.Role{Id: id}).Scan(&role)
	return
}

// GetMenuIdsByRoleId 获取角色ID关联的菜单ID列表
func (s *sRole) GetMenuIdsByRoleId(ctx context.Context, id uint64) (menuIds []uint64, err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}

	gp := enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))
	menuIds = make([]uint64, len(gp))
	for k, v := range gp {
		menuIds[k] = gconv.Uint64(v[1])
	}
	return
}

// IsRoleNameAvailable 角色名称是否可用
func (s *sRole) IsRoleNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.Role.Ctx(ctx).Where(do.Role{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// SaveRole 保存角色
func (s *sRole) SaveRole(ctx context.Context, req *v1.RoleAddReq) (roleId uint64, err error) {
	var id int64
	id, err = dao.Role.Ctx(ctx).Data(do.Role{
		Name:      req.Name,
		Status:    req.Status,
		DataScope: consts.DataScopeDept,
		Remark:    req.Remark,
	}).InsertAndGetId()
	roleId = gconv.Uint64(id)
	return
}

// UpdateRole 更新角色
func (s *sRole) UpdateRole(ctx context.Context, req *v1.RoleEditReq) (err error) {
	_, err = dao.Role.Ctx(ctx).Data(do.Role{
		Name:   req.Name,
		Status: req.Status,
		Remark: req.Remark,
	}).Where(do.Role{Id: req.Id}).Update()
	return
}

// AddRoleRule 添加角色权限
func (s *sRole) AddRoleRule(ctx context.Context, iRule any, roleId uint64) (err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}
	rule := gconv.Strings(iRule)
	for _, v := range rule {
		_, err = enforcer.AddPolicy(gconv.String(roleId), gconv.String(v), "All")
		if err != nil {
			break
		}
	}
	return
}

// EditRoleRule 修改角色权限
func (s *sRole) EditRoleRule(ctx context.Context, iRule any, roleId uint64) (err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}
	// 删除旧权限
	_, err = enforcer.RemoveFilteredPolicy(0, gconv.String(roleId))
	if err != nil {
		return
	}
	// 添加新权限
	rule := gconv.Strings(iRule)
	for _, v := range rule {
		_, err = enforcer.AddPolicy(gconv.String(roleId), gconv.String(v), "All")
		if err != nil {
			break
		}
	}
	return
}

// GetUserRoleIds 获取用户角色ID列表
func (s *sRole) GetUserRoleIds(ctx context.Context, id uint64) (roleIds []uint64, err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}
	// 查询关联角色规则
	groupPolicy := enforcer.GetFilteredGroupingPolicy(0, gconv.String(id))
	if len(groupPolicy) > 0 {
		roleIds = make([]uint64, len(groupPolicy))
		// 得到角色ID的切片
		for k, v := range groupPolicy {
			roleIds[k] = gconv.Uint64(v[1])
		}
	}
	return
}

// GetUserRoles 获取用户角色
func (s *sRole) GetUserRoles(ctx context.Context, id uint64) (roles []*entity.Role, err error) {
	// 获取所有角色
	var allRoles []*entity.Role
	allRoles, err = s.GetAllRoles(ctx)
	if err != nil {
		return
	}
	// 获取用户角色ID列表
	var roleIds []uint64
	roleIds, err = s.GetUserRoleIds(ctx, id)
	if err != nil {
		return
	}

	roles = make([]*entity.Role, 0, len(allRoles))
	for _, v := range allRoles {
		for _, rid := range roleIds {
			if rid == v.Id {
				roles = append(roles, v)
			}
		}
		if len(roles) == len(roleIds) {
			break
		}
	}
	return
}

// GetUserMenuList 获取用户菜单列表
func (s *sRole) GetUserMenuList(ctx context.Context, ids []uint64) (menuList []string, err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}

	menuIds := map[uint64]uint64{}
	for _, id := range ids {
		// 查询当前权限
		gp := enforcer.GetFilteredPolicy(0, gconv.String(id))
		for _, p := range gp {
			mid := gconv.Uint64(p[1])
			menuIds[mid] = mid
		}
	}
	// 获取所有正常状态的菜单列表
	var menus []*entity.Menu
	menus, err = Menu().GetStatusEnableMenus(ctx)
	menuList = make([]string, 0, len(menus))
	for _, m := range menus {
		if _, ok := menuIds[gconv.Uint64(m.Id)]; gstr.Equal(m.Condition, "nocheck") || ok {
			menuList = append(menuList, m.Rule)
		}
	}
	return
}

// AddUserRoles 添加用户角色信息
func (s *sRole) AddUserRoles(ctx context.Context, roleIds []uint64, userId uint64) (err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}

	for _, rid := range roleIds {
		_, err = enforcer.AddGroupingPolicy(gconv.String(userId), gconv.String(rid))
		if err != nil {
			return
		}
	}
	return
}

// EditUserRoles 修改用户角色信息
func (s *sRole) EditUserRoles(ctx context.Context, roleIds []uint64, userId uint64) (err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}
	// 删除用户旧角色信息
	_, err = enforcer.RemoveFilteredGroupingPolicy(0, gconv.String(userId))
	if err != nil {
		return
	}
	for _, rid := range roleIds {
		_, err = enforcer.AddGroupingPolicy(gconv.String(userId), gconv.String(rid))
		if err != nil {
			return
		}
	}
	return
}
