package controller

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

// List 菜单树列表
func (c *cMenu) List(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	var list []*entity.Menu
	list, err = service.Menu().GetMenuList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetMenuListFailed, err)
		return
	}

	treeInfos := make([]*v1.MenuTreeInfo, 0, len(list))
	idsMap := gmap.New()
	for _, v := range list {
		service.Menu().FindSonIdsByParentId(list, v.Id, idsMap)
		if !idsMap.Contains(v.Id) {
			t := &v1.MenuTreeInfo{Menu: v}
			children := service.Menu().GetMenuTree(list, v.Id)
			if len(children) > 0 {
				t.Children = children
			}
			treeInfos = append(treeInfos, t)
		}
	}
	res = &v1.MenuListRes{List: treeInfos}
	return
}

// IsMenus 菜单类型为目录和菜单的菜单树列表
func (c *cMenu) IsMenus(ctx context.Context, req *v1.MenuIsMenusReq) (res *v1.MenuIsMenusRes, err error) {
	var menuList []*entity.Menu
	menuList, err = service.Menu().GetIsMenus(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetIsMenusFailed, err)
		return
	}

	list := service.Menu().GetMenuTree(menuList, 0)
	res = &v1.MenuIsMenusRes{List: list}
	return
}

// Add 新增菜单
func (c *cMenu) Add(ctx context.Context, req *v1.MenuAddReq) (res *v1.MenuAddRes, err error) {
	err = service.Menu().AddMenu(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddMenuFailed, err)
		return
	}

	return
}

// Edit 编辑菜单
func (c *cMenu) Edit(ctx context.Context, req *v1.MenuEditReq) (res *v1.MenuEditRes, err error) {
	err = service.Menu().EditMenu(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditMenuFailed, err)
		return
	}

	return
}

// Delete 删除菜单
func (c *cMenu) Delete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error) {
	err = service.Menu().DeleteMenu(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteMenuFailed, err)
		return
	}

	return
}

// Tree 菜单树信息
func (c *cMenu) Tree(ctx context.Context, req *v1.MenuTreeReq) (res *v1.MenuTreeRes, err error) {
	// 获取所有菜单
	var allMenus []*entity.Menu
	allMenus, err = service.Menu().GetAllMenus(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetMenuTreeFailed, err)
		return
	}

	tree := service.Menu().GetMenuTree(allMenus, 0)
	res = &v1.MenuTreeRes{List: tree}
	return
}
