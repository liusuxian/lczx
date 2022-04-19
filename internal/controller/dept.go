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
	Dept = cDept{}
)

type cDept struct{}

// List 获取部门列表
func (c *cDept) List(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	var list []*entity.Dept
	list, err = service.Dept().GetDeptList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetDeptListFailed, err)
		return
	}

	treeInfos := make([]*v1.DeptTreeInfo, 0, len(list))
	idsMap := gmap.New(true)
	for _, v := range list {
		service.Dept().FindSonIdsByParentId(list, v.Id, idsMap)
		if !idsMap.Contains(v.Id) {
			t := &v1.DeptTreeInfo{Dept: v}
			children := service.Dept().GetDeptTree(list, v.Id)
			if len(children) > 0 {
				t.Children = children
			}
			treeInfos = append(treeInfos, t)
		}
	}
	res = &v1.DeptListRes{List: treeInfos}
	return
}

// Add 新增部门
func (c *cDept) Add(ctx context.Context, req *v1.DeptAddReq) (res *v1.DeptAddRes, err error) {
	err = service.Dept().AddDept(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddDeptFailed, err)
		return
	}

	return
}

// Info 获取部门信息
func (c *cDept) Info(ctx context.Context, req *v1.DeptInfoReq) (res *v1.DeptInfoRes, err error) {
	var dept *entity.Dept
	dept, err = service.Dept().GetDeptById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetDeptFailed, err)
		return
	}

	res = &v1.DeptInfoRes{Info: dept}
	return
}

// Edit 编辑部门
func (c *cDept) Edit(ctx context.Context, req *v1.DeptEditReq) (res *v1.DeptEditRes, err error) {
	err = service.Dept().EditDept(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditDeptFailed, err)
		return
	}

	return
}

// Delete 删除部门
func (c *cDept) Delete(ctx context.Context, req *v1.DeptDeleteReq) (res *v1.DeptDeleteRes, err error) {
	err = service.Dept().DeleteDept(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteDeptFailed, err)
		return
	}

	return
}

// Tree 部门树信息
func (c *cDept) Tree(ctx context.Context, req *v1.DeptTreeReq) (res *v1.DeptTreeRes, err error) {
	// 获取部门状态为正常的部门列表
	var list []*entity.Dept
	list, err = service.Dept().GetStatusEnableDepts(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetDeptTreeFailed, err)
		return
	}

	tree := service.Dept().GetDeptTree(list, 0)
	res = &v1.DeptTreeRes{List: tree}
	return
}

// RoleDeptTree 角色部门树信息
func (c *cDept) RoleDeptTree(ctx context.Context, req *v1.DeptRoleDeptTreeReq) (res *v1.DeptRoleDeptTreeRes, err error) {
	// 获取部门状态为正常的部门列表
	var list []*entity.Dept
	list, err = service.Dept().GetStatusEnableDepts(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetRoleDeptTreeFailed, err)
		return
	}
	treeInfos := service.Dept().GetDeptTree(list, 0)
	// 获取关联的角色数据权限
	var deptIds []uint64
	deptIds, err = service.Dept().GetDeptIdsByRoleId(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetRoleDeptTreeFailed, err)
		return
	}

	res = &v1.DeptRoleDeptTreeRes{
		List:    treeInfos,
		DeptIds: deptIds,
	}
	return
}
