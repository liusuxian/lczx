package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sDept struct{}

var insDept = sDept{}

func Dept() *sDept {
	return &insDept
}

// GetDeptList 获取部门列表
func (s *sDept) GetDeptList(ctx context.Context) (deptList []*entity.Dept, err error) {
	err = dao.Dept.Ctx(ctx).Scan(&deptList)
	return
}

// IsDeptNameAvailable 部门名称是否可用
func (s *sDept) IsDeptNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// ExistsById 通过部门ID判断部门信息是否存在
func (s *sDept) ExistsById(ctx context.Context, id uint) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Id: id}).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// AddDept 新增部门
func (s *sDept) AddDept(ctx context.Context, name string) (id int64, err error) {
	var available bool
	available, err = s.IsDeptNameAvailable(ctx, name)
	if err != nil {
		return
	}
	if !available {
		return
	}
	return id, dao.Dept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		id, err = dao.Dept.Ctx(ctx).Data(do.Dept{Name: name}).InsertAndGetId()
		return err
	})
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, id uint) (deleteId uint, err error) {
	var isExists bool
	isExists, err = s.ExistsById(ctx, id)
	if err != nil {
		return
	}
	if !isExists {
		return
	}
	deleteId = id
	err = dao.Dept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Dept.Ctx(ctx).Delete(do.Dept{Id: id})
		return err
	})
	return
}

// UpdateDept 修改部门
func (s *sDept) UpdateDept(ctx context.Context, id uint, name string) (updateId uint, err error) {
	var isExists bool
	isExists, err = s.ExistsById(ctx, id)
	if err != nil {
		return
	}
	if !isExists {
		return
	}
	updateId = id
	err = dao.Dept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Dept.Ctx(ctx).Data(do.Dept{Name: name}).Where(do.Dept{Id: id}).Update()
		return err
	})
	return
}
