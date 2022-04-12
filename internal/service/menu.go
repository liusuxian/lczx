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
func (s *sMenu) GetMenuList(ctx context.Context, req *v1.MenuListReq, fieldNames ...string) (list []*entity.AuthRule, err error) {
	model := dao.AuthRule.Ctx(ctx)
	columns := dao.AuthRule.Columns()
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

// GetMenuListTree 获取菜单树列表
func (s *sMenu) GetMenuListTree() {

}
