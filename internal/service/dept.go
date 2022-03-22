package service

import (
	"context"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
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
