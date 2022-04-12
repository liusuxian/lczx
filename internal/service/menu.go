package service

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
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
