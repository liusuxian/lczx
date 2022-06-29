package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	Role = cRole{}
)

type cRole struct{}

// ClientOptions 客户端选项
func (c *cRole) ClientOptions(ctx context.Context, req *v1.RoleClientOptionsReq) (res *v1.RoleClientOptionsRes, err error) {
	// 获取客户端选项Map
	clientOptionMap := service.Role().GetClientOptionMap()

	res = &v1.RoleClientOptionsRes{
		StatusList:    clientOptionMap["statusList"],
		DataScopeList: clientOptionMap["dataScopeList"],
	}
	return
}

// List 角色列表
func (c *cRole) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	var total int
	var list []*entity.Role
	total, list, err = service.Role().GetRoleList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetRoleListFailed, err)
		return
	}

	res = &v1.RoleListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// Add 新增角色
func (c *cRole) Add(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleAddRes, err error) {
	err = service.Role().AddRole(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddRoleFailed, err)
		return
	}

	return
}

// Info 获取角色信息
func (c *cRole) Info(ctx context.Context, req *v1.RoleInfoReq) (res *v1.RoleInfoRes, err error) {
	var role *entity.Role
	role, err = service.Role().GetRoleById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetRoleFailed, err)
		return
	}
	// 获取所有菜单
	var allMenus []*entity.Menu
	allMenus, err = service.Menu().GetAllMenus(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetRoleFailed, err)
		return
	}
	menuList := service.Menu().GetMenuTree(allMenus, 0)
	// 获取角色ID关联的菜单ID列表
	var menuIds []uint64
	menuIds, err = service.Role().GetMenuIdsByRoleId(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetRoleFailed, err)
		return
	}

	res = &v1.RoleInfoRes{
		Info:     role,
		MenuList: menuList,
		MenuIds:  menuIds,
	}
	return
}

// Edit 编辑角色
func (c *cRole) Edit(ctx context.Context, req *v1.RoleEditReq) (res *v1.RoleEditRes, err error) {
	err = service.Role().EditRole(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditRoleFailed, err)
		return
	}

	return
}

// SetStatus 设置角色状态
func (c *cRole) SetStatus(ctx context.Context, req *v1.RoleSetStatusReq) (res *v1.RoleSetStatusRes, err error) {
	err = service.Role().SetRoleStatus(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.SetRoleStatusFailed, err)
		return
	}

	return
}

// SetDataScope 设置数据权限
func (c *cRole) SetDataScope(ctx context.Context, req *v1.RoleSetDataScopeReq) (res *v1.RoleSetDataScopeRes, err error) {
	err = service.Role().SetRoleDataScope(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.SetRoleDataScopeFailed, err)
		return
	}

	return
}

// Delete 删除角色
func (c *cRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	err = service.Role().DeleteRoleByIds(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteRoleFailed, err)
		return
	}

	return
}

// EnableRoles 获取全部可用的角色
func (c *cRole) EnableRoles(ctx context.Context, req *v1.RoleEnableRolesReq) (res *v1.RoleEnableRolesRes, err error) {
	var roles []*entity.Role
	roles, err = service.Role().GetEnableRoles(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetEnableRolesFailed, err)
		return
	}

	res = &v1.RoleEnableRolesRes{List: roles}
	return
}
