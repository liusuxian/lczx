package service

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/utils"
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
func (s *sRole) GetRoleList(ctx context.Context, req *v1.RoleListReq, fieldNames ...string) (total int, list []*entity.Role, err error) {
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
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
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
	_, err = g.Redis().Do(ctx, "DEL", utils.GetCacheRoleKey())
	return
}

// GetRoleById 通过角色ID获取角色信息
func (s *sRole) GetRoleById(ctx context.Context, id uint64, fieldNames ...string) (role *entity.Role, err error) {
	model := dao.Role.Ctx(ctx).Where(do.Role{Id: id})
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
	}
	err = model.Scan(&role)
	return
}

// GetMenuIdsByRoleId 获取角色ID关联的菜单ID列表
func (s *sRole) GetMenuIdsByRoleId(ctx context.Context, id uint64) (menuIds []uint64, err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}

	gp := enforcer.GetFilteredNamedPolicy("p", 0, fmt.Sprintf("%d", id))
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
		DataScope: req.DataScope,
		Remark:    req.Remark,
	}).InsertAndGetId()
	roleId = gconv.Uint64(id)
	return
}

// AddRoleRule 添加角色权限
func (s *sRole) AddRoleRule(ctx context.Context, iRule interface{}, roleId uint64) (err error) {
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = Casbin(ctx).GetEnforcer()
	if err != nil {
		return
	}
	rule := gconv.Strings(iRule)
	for _, v := range rule {
		_, err = enforcer.AddPolicy(fmt.Sprintf("%d", roleId), fmt.Sprintf("%s", v), "All")
		if err != nil {
			break
		}
	}
	return
}
