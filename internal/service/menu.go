package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
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
func (s *sMenu) GetMenuList(ctx context.Context, req *v1.MenuListReq) (list []*entity.Menu, err error) {
	model := dao.Menu.Ctx(ctx)
	columns := dao.Menu.Columns()
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	err = model.OrderAsc(columns.Id).Scan(&list)
	return
}

// AddMenu 新增菜单
func (s *sMenu) AddMenu(ctx context.Context, req *v1.MenuAddReq) (err error) {
	if req.ParentId == 0 && req.MenuType != consts.MenuTypeDir {
		err = gerror.Newf(`父规则ID[%d]只能添加目录类型`, req.ParentId)
		return
	}
	// 检查父规则ID是否存在
	if req.ParentId != 0 {
		var parentMenu *entity.Menu
		parentMenu, err = s.GetMenuById(ctx, req.ParentId)
		if err != nil {
			return
		}
		if parentMenu == nil {
			err = gerror.Newf(`父规则ID[%d]不存在`, req.ParentId)
			return
		}
		if parentMenu.Status == consts.MenuStatusDisabled {
			err = gerror.Newf(`父规则ID[%d]已停用`, req.ParentId)
			return
		}
		if parentMenu.MenuType == consts.MenuTypeDir && req.MenuType == consts.MenuTypeButton {
			err = gerror.Newf(`父规则菜单类型[%d]不能添加按钮类型`, parentMenu.MenuType)
			return
		}
		if parentMenu.MenuType == consts.MenuTypeMenu && req.MenuType != consts.MenuTypeButton {
			err = gerror.Newf(`父规则菜单类型[%d]只能添加按钮类型`, parentMenu.MenuType)
			return
		}
		if parentMenu.MenuType == consts.MenuTypeButton {
			err = gerror.Newf(`父规则菜单类型[%d]不能添加任何类型`, parentMenu.MenuType)
			return
		}
	}
	// 检查权限规则是否可用
	var available bool
	available, err = s.IsMenuRuleAvailable(ctx, req.Rule)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`权限规则[%s]已存在`, req.Rule)
		return
	}
	// 写入菜单数据
	_, err = dao.Menu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.MenuKey,
		Force:    false,
	}).Data(do.Menu{
		ParentId:   req.ParentId,
		Rule:       req.Rule,
		Name:       req.Name,
		Condition:  req.Condition,
		MenuType:   req.MenuType,
		Status:     req.Status,
		JumpPath:   req.JumpPath,
		IsFrame:    req.IsFrame,
		ModuleType: req.ModuleType,
		Remark:     req.Remark,
	}).Insert()
	return
}

// EditMenu 编辑菜单
func (s *sMenu) EditMenu(ctx context.Context, req *v1.MenuEditReq) (err error) {
	if req.ParentId == 0 && req.MenuType != consts.MenuTypeDir {
		err = gerror.Newf(`父规则ID[%d]只能添加目录类型`, req.ParentId)
		return
	}
	// 检查父规则ID是否存在
	if req.ParentId != 0 {
		var parentMenu *entity.Menu
		parentMenu, err = s.GetMenuById(ctx, req.ParentId)
		if err != nil {
			return
		}
		if parentMenu == nil {
			err = gerror.Newf(`父规则ID[%d]不存在`, req.ParentId)
			return
		}
		if parentMenu.Status == consts.MenuStatusDisabled {
			err = gerror.Newf(`父规则ID[%d]已停用`, req.ParentId)
			return
		}
		if parentMenu.MenuType == consts.MenuTypeDir && req.MenuType == consts.MenuTypeButton {
			err = gerror.Newf(`父规则菜单类型[%d]不能添加按钮类型`, parentMenu.MenuType)
			return
		}
		if parentMenu.MenuType == consts.MenuTypeMenu && req.MenuType != consts.MenuTypeButton {
			err = gerror.Newf(`父规则菜单类型[%d]只能添加按钮类型`, parentMenu.MenuType)
			return
		}
		if parentMenu.MenuType == consts.MenuTypeButton {
			err = gerror.Newf(`父规则菜单类型[%d]不能添加任何类型`, parentMenu.MenuType)
			return
		}
	}
	// 检查菜单信息是否存在
	var menu *entity.Menu
	menu, err = s.GetMenuById(ctx, req.Id)
	if err != nil {
		return
	}
	if menu == nil {
		err = gerror.Newf(`菜单ID[%d]不存在`, req.Id)
		return
	}
	// 检查权限规则是否可用
	if menu.Rule != req.Rule {
		var available bool
		available, err = s.IsMenuRuleAvailable(ctx, req.Rule)
		if err != nil {
			return
		}
		if !available {
			err = gerror.Newf(`权限规则[%s]已存在`, req.Rule)
			return
		}
	}
	// 获取所有菜单
	var list []*entity.Menu
	list, err = s.GetAllMenus(ctx)
	if err != nil {
		return
	}
	// 获取规则ID下所有的子规则ID
	idsMap := gmap.New(true)
	s.FindSonIdsByParentId(list, req.Id, idsMap)
	if idsMap.Contains(req.ParentId) {
		err = gerror.Newf(`父规则ID[%d]是规则ID[%d]的子规则ID`, req.ParentId, req.Id)
		return
	}
	// 更新菜单数据
	_, err = dao.Menu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.MenuKey,
		Force:    false,
	}).Data(do.Menu{
		ParentId:   req.ParentId,
		Rule:       req.Rule,
		Name:       req.Name,
		Condition:  req.Condition,
		MenuType:   req.MenuType,
		Status:     req.Status,
		JumpPath:   req.JumpPath,
		IsFrame:    req.IsFrame,
		ModuleType: req.ModuleType,
		Remark:     req.Remark,
	}).Where(do.Menu{Id: req.Id}).Update()
	return
}

// DeleteMenu 删除菜单
func (s *sMenu) DeleteMenu(ctx context.Context, ids []uint64) (err error) {
	// 获取所有菜单
	var list []*entity.Menu
	list, err = s.GetAllMenus(ctx)
	if err != nil {
		return
	}
	// 获取所有的子规则ID
	idsMap := gmap.New(true)
	for _, id := range ids {
		idsMap.Set(id, true)
		s.FindSonIdsByParentId(list, id, idsMap)
	}
	delIds := idsMap.Keys()
	// 删除菜单数据
	_, err = dao.Menu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.MenuKey,
		Force:    false,
	}).WhereIn(dao.Menu.Columns().Id, delIds).Delete()
	return
}

// GetIsMenus 获取菜单类型为目录和菜单的菜单列表
func (s *sMenu) GetIsMenus(ctx context.Context) (list []*entity.Menu, err error) {
	// 获取所有菜单
	var menus []*entity.Menu
	menus, err = s.GetAllMenus(ctx)
	if err != nil {
		return
	}

	list = make([]*entity.Menu, 0, len(menus))
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
	var menusCacheValue *gvar.Var
	menusCacheValue, err = Cache().GetCache(ctx, consts.MenuKey)
	if err != nil {
		return
	}
	if menusCacheValue != nil {
		menusCacheMap := menusCacheValue.Map()["Result"]
		if menusCacheMap != nil {
			err = gconv.Structs(menusCacheMap, &menus)
			if err != nil {
				return
			}
			if menus != nil {
				return
			}
		}
	}
	// 从数据库获取
	err = dao.Menu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     consts.MenuKey,
		Force:    false,
	}).OrderAsc(dao.Menu.Columns().Id).Scan(&menus)
	return
}

// GetMenuById 通过规则ID获取菜单信息
func (s *sMenu) GetMenuById(ctx context.Context, id uint64) (menu *entity.Menu, err error) {
	err = dao.Menu.Ctx(ctx).Where(do.Dept{Id: id}).Scan(&menu)
	return
}

// IsMenuRuleAvailable 权限规则是否可用
func (s *sMenu) IsMenuRuleAvailable(ctx context.Context, rule string) (bool, error) {
	count, err := dao.Menu.Ctx(ctx).Where(do.Menu{Rule: rule}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// FindSonIdsByParentId 通过父规则ID获取所有的子规则ID
func (s *sMenu) FindSonIdsByParentId(menuList []*entity.Menu, parentId uint64, idsMap *gmap.Map) {
	for _, v := range menuList {
		if v.ParentId == parentId {
			idsMap.Set(v.Id, true)
			s.FindSonIdsByParentId(menuList, v.Id, idsMap)
		}
	}
}

// GetStatusEnableMenus 获取所有正常状态的菜单列表
func (s *sMenu) GetStatusEnableMenus(ctx context.Context) (menuList []*entity.Menu, err error) {
	// 获取所有菜单
	var menus []*entity.Menu
	menus, err = s.GetAllMenus(ctx)
	if err != nil {
		return
	}

	menuList = make([]*entity.Menu, 0, len(menus))
	for _, v := range menus {
		if v.Status == consts.MenuStatusEnable {
			menuList = append(menuList, v)
		}
	}
	return
}
