package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sDept struct{}

var insDept = sDept{}

// Dept 部门管理服务
func Dept() *sDept {
	return &insDept
}

// IsDeptNameAvailable 部门名称是否可用
func (s *sDept) IsDeptNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// DeptExistsById 通过部门ID判断部门信息是否存在
func (s *sDept) DeptExistsById(ctx context.Context, id uint) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Id: id}).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// GetDeptList 获取部门列表
func (s *sDept) GetDeptList(ctx context.Context) (deptList []*entity.Dept, err error) {
	err = dao.Dept.Ctx(ctx).Scan(&deptList)
	return
}

// AddDept 新增部门
func (s *sDept) AddDept(ctx context.Context, name string) (id int64, err error) {
	var available bool
	available, err = s.IsDeptNameAvailable(ctx, name)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`部门名称[%s]已存在`, name)
		return
	}
	model := dao.Dept.Ctx(ctx)
	err = model.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		id, err = model.Data(do.Dept{Name: name}).InsertAndGetId()
		return err
	})
	return
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, id uint) (err error) {
	var deptExists bool
	deptExists, err = s.DeptExistsById(ctx, id)
	if err != nil {
		return
	}
	if !deptExists {
		err = gerror.Newf(`部门ID[%d]不存在`, id)
		return
	}
	_, err = dao.Dept.Ctx(ctx).Delete(do.Dept{Id: id})
	return
}

// UpdateDept 更新部门
func (s *sDept) UpdateDept(ctx context.Context, req *v1.DeptUpdateReq) (err error) {
	var deptExists bool
	deptExists, err = s.DeptExistsById(ctx, req.Id)
	if err != nil {
		return
	}
	if !deptExists {
		err = gerror.Newf(`部门ID[%d]不存在`, req.Id)
		return
	}
	_, err = dao.Dept.Ctx(ctx).Data(do.Dept{Name: req.Name}).Where(do.Dept{Id: req.Id}).Update()
	return
}
