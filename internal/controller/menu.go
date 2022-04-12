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
	list, err = service.Menu().GetMenuList(ctx, req, "createAt", "updateAt", "deletedAt")
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
