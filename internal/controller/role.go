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
	var menus []*entity.Menu
	menus, err = service.Menu().GetAllMenus(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetMenuTreeFailed, err)
		return
	}
	menuList := service.Menu().GetMenuTree(menus, 0)
	// 获取角色ID关联的菜单ID列表
	var menuIds []uint64
	menuIds, err = service.Role().GetMenuIdsByRoleId(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetMenuIdsByRoleFailed, err)
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
