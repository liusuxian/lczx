package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
)

type sMenu struct{}

var (
	insMenu = sMenu{}
)

// Menu 菜单管理服务
func Menu() *sMenu {
	return &insMenu
}

// GetMenuList 获取菜单列表
func (s *sMenu) GetMenuList(ctx context.Context, req *v1.MenuListReq, fieldNames ...string) (list []*entity.Menu, err error) {
	model := dao.Menu.Ctx(ctx)
	columns := dao.Menu.Columns()
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
	}
	err = model.OrderAsc(columns.Id).Scan(&list)
	return
}

// GetIsMenus 获取菜单类型为目录和菜单的菜单列表
func (s *sMenu) GetIsMenus(ctx context.Context) (list []*entity.Menu, err error) {
	var menus []*entity.Menu
	menus, err = s.GetAllMenus(ctx)
	if err != nil {
		return
	}

	list = make([]*entity.Menu, 0, len(list))
	for _, v := range menus {
		if v.MenuType == consts.MenuTypeDir || v.MenuType == consts.MenuTypeMenu {
			list = append(list, v)
		}
	}
	return
}

// GetMenuTree 获取菜单树信息
func (s *sMenu) GetMenuTree(menuList []*entity.Menu, parentId uint64) (tree []*v1.MenuTreeInfo) {
	tree = make([]*v1.MenuTreeInfo, 0, len(menuList))
	for _, v := range menuList {
		if v.ParentId == parentId {
			t := &v1.MenuTreeInfo{Menu: v}
			children := s.GetMenuTree(menuList, v.Id)
			if len(children) > 0 {
				t.Children = children
			}
			tree = append(tree, t)
		}
	}
	return
}

// GetAllMenus 获取所有菜单
func (s *sMenu) GetAllMenus(ctx context.Context) (menus []*entity.Menu, err error) {
	// 从缓存获取
	var menusVar *gvar.Var
	menusVar, err = g.DB().GetCache().Get(ctx, consts.Menu)
	if err != nil {
		return
	}
	menusMap := menusVar.Map()["Result"]
	if menusMap != nil {
		err = gconv.Structs(menusMap, &menus)
		if err != nil {
			return
		}
		if menus != nil {
			return
		}
	}
	// 从数据库获取
	err = dao.Menu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     consts.Menu,
		Force:    false,
	}).OrderAsc(dao.Menu.Columns().Id).Scan(&menus)
	return
}
