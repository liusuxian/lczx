package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/logger"
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

// AddDept 新增部门
func (s *sDept) AddDept(ctx context.Context, name string) (id int64, err error) {
	var available bool
	available, err = s.IsDeptNameAvailable(ctx, name)
	if err != nil {
		return
	}
	if !available {
		logger.Warningf(ctx, "部门名称: %s 已存在", name)
		return
	}
	return id, dao.Dept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		id, err = dao.Dept.Ctx(ctx).Data(do.Dept{Name: name}).InsertAndGetId()
		return err
	})
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, id uint) (deleteId uint, err error) {
	var dept *entity.Dept
	dept, err = s.GetDeptById(ctx, id)
	if err != nil {
		return
	}
	if dept == nil {
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
	var dept *entity.Dept
	dept, err = s.GetDeptById(ctx, id)
	if err != nil {
		return
	}
	if dept == nil {
		return
	}

	updateId = id
	err = dao.Dept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Dept.Ctx(ctx).Data(do.Dept{Name: name}).Where(do.Dept{Id: id}).Update()
		return err
	})
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

// GetDeptById 通过部门ID获取部门信息
func (s *sDept) GetDeptById(ctx context.Context, id uint) (dept *entity.Dept, err error) {
	err = dao.Dept.Ctx(ctx).Where(do.Dept{Id: id}).Scan(&dept)
	return
}
