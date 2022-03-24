package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
)

var (
	Dept = cDept{}
)

type cDept struct{}

// List 获取部门列表
func (c *cDept) List(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	var deptList []*entity.Dept
	deptList, err = service.Dept().GetDeptList(ctx)
	if err != nil {
		err = gerror.WrapCode(code.GetDeptListFailed, err)
		return
	}

	deptInfoList := make([]*v1.DeptInfo, len(deptList))
	for k, v := range deptList {
		deptInfoList[k] = &v1.DeptInfo{Id: v.Id, Name: v.Name}
	}
	res = &v1.DeptListRes{List: deptInfoList}
	return
}

// Add 新增部门
func (c *cDept) Add(ctx context.Context, req *v1.DeptAddReq) (res *v1.DeptAddRes, err error) {
	logger.Debug(ctx, "Add Req: ", req)
	var deptId int64
	deptId, err = service.Dept().AddDept(ctx, req.Name)
	if err != nil {
		err = gerror.WrapCode(code.AddDeptFailed, err)
		return
	}

	newDeptId := gconv.Uint(deptId)
	res = &v1.DeptAddRes{Dept: &v1.DeptInfo{Id: newDeptId, Name: req.Name}}
	return
}

// Delete 删除部门
func (c *cDept) Delete(ctx context.Context, req *v1.DeptDeleteReq) (res *v1.DeptDeleteRes, err error) {
	logger.Debug(ctx, "Delete Req: ", req)
	err = service.Dept().DeleteDept(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.DeleteDeptFailed, err)
		return
	}

	res = &v1.DeptDeleteRes{Id: req.Id}
	return
}

// Update 修改部门
func (c *cDept) Update(ctx context.Context, req *v1.DeptUpdateReq) (res *v1.DeptUpdateRes, err error) {
	logger.Debug(ctx, "Update Req: ", req)
	err = service.Dept().UpdateDept(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.UpdateDeptFailed, err)
		return
	}

	res = &v1.DeptUpdateRes{Dept: &v1.DeptInfo{Id: req.Id, Name: req.Name}}
	return
}
