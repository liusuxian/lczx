package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/response"
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
		logger.Error(ctx, "GetDeptList Error: ", err.Error())
		response.Json(ctx, consts.CodeGetDeptListFailed, nil)
		return
	}

	deptInfoList := make([]*v1.DeptInfo, len(deptList))
	for k, v := range deptList {
		deptInfoList[k] = &v1.DeptInfo{Id: v.Id, Name: v.Name}
	}
	response.JsonOK(ctx, &v1.DeptListRes{List: deptInfoList})
	return
}

// Add 新增部门
func (c *cDept) Add(ctx context.Context, req *v1.DeptAddReq) (res *v1.DeptAddRes, err error) {
	logger.Debug(ctx, "Add Req: ", req)
	var deptId int64
	deptId, err = service.Dept().AddDept(ctx, req.Name)
	if err != nil {
		logger.Error(ctx, "Add Error: ", err.Error())
		response.Json(ctx, consts.CodeAddDeptFailed, nil)
		return
	}

	newDeptId := gconv.Uint(deptId)
	if newDeptId > 0 {
		response.JsonOK(ctx, &v1.DeptAddRes{Dept: &v1.DeptInfo{Id: newDeptId, Name: req.Name}})
	} else {
		response.Json(ctx, consts.CodeDeptIDNotExist, nil)
	}
	return
}

// Delete 删除部门
func (c *cDept) Delete(ctx context.Context, req *v1.DeptDeleteReq) (res *v1.DeptDeleteRes, err error) {
	logger.Debug(ctx, "Delete Req: ", req)
	var deleteId uint
	deleteId, err = service.Dept().DeleteDept(ctx, req.Id)
	if err != nil {
		logger.Error(ctx, "Delete Error: ", err.Error())
		response.Json(ctx, consts.CodeDeleteDeptFailed, nil)
		return
	}
	if deleteId > 0 {
		response.JsonOK(ctx, &v1.DeptDeleteRes{Id: deleteId})
	} else {
		response.Json(ctx, consts.CodeDeptIDNotExist, nil)
	}
	return
}

// Update 修改部门
func (c *cDept) Update(ctx context.Context, req *v1.DeptUpdateReq) (res *v1.DeptUpdateRes, err error) {
	logger.Debug(ctx, "Update Req: ", req)
	var updateId uint
	updateId, err = service.Dept().UpdateDept(ctx, req)
	if err != nil {
		logger.Error(ctx, "Update Error: ", err.Error())
		response.Json(ctx, consts.CodeUpdateDeptFailed, nil)
		return
	}
	if updateId > 0 {
		response.JsonOK(ctx, &v1.DeptUpdateRes{Dept: &v1.DeptInfo{Id: updateId, Name: req.Name}})
	} else {
		response.Json(ctx, consts.CodeDeptIDNotExist, nil)
	}
	return
}
