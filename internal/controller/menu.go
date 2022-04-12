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
	Menu = cMenu{}
)

type cMenu struct{}

// List 菜单列表
func (c *cMenu) List(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	var list []*entity.Menu
	list, err = service.Menu().GetMenuList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetMenuListFailed, err)
		return
	}
	treeInfos := make([]*v1.MenuTreeInfo, 0, len(list))
	if !(req.Name == "" && req.Status == "") {
		for _, v := range list {
			treeInfos = append(treeInfos, &v1.MenuTreeInfo{Menu: v})
		}
	} else {
		treeInfos = service.Menu().GetMenuTree(list, 0)
	}

	res = &v1.MenuListRes{List: treeInfos}
	return
}

// IsMenus 获取菜单类型为目录和菜单的菜单列表
func (c *cMenu) IsMenus(ctx context.Context, req *v1.MenuIsMenusReq) (res *v1.MenuIsMenusRes, err error) {
	var list []*entity.Menu
	list, err = service.Menu().GetIsMenus(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetIsMenusFailed, err)
		return
	}

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
