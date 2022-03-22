package controller

import (
	"context"
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

// GetDeptList 获取部门列表
func (c *cDept) GetDeptList(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	userInfo := service.Context().Get(ctx).User
	if userInfo != nil && userInfo.Id > 0 {
		var deptList []*entity.Dept
		deptList, err = service.Dept().GetDeptList(ctx)
		if err != nil {
			logger.Error(ctx, "GetDeptList Error: ", err.Error())
			response.Json(ctx, consts.CodeGetDeptListFailed, nil)
			return
		}

		response.JsonOK(ctx, v1.DeptListRes{
			DeptList: deptList,
		})
		return
	}

	response.Json(ctx, consts.CodeGetDeptListFailed, nil)
	return
}
