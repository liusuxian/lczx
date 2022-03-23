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
	return deptList, err
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

// IsDeptNameAvailable 部门名称是否可用
func (s *sDept) IsDeptNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
