package service

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sDept struct{}

var (
	insDept = sDept{}
)

// Dept 部门管理服务
func Dept() *sDept {
	return &insDept
}

// GetDeptList 获取部门列表
func (s *sDept) GetDeptList(ctx context.Context, req *v1.DeptListReq, fieldNames ...string) (list []*entity.Dept, err error) {
	model := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	if req.Name != "" {
		model = model.WhereIn(columns.Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
	}
	err = model.Scan(&list)
	return
}

// AddDept 新增部门
func (s *sDept) AddDept(ctx context.Context, req *v1.DeptAddReq) (err error) {
	// 检查父部门是否存在
	if req.ParentId != 0 {
		var deptExists bool
		deptExists, err = s.DeptExistsById(ctx, req.ParentId)
		if err != nil {
			return
		}
		if !deptExists {
			err = gerror.Newf(`父部门ID[%d]不存在`, req.ParentId)
			return
		}
	}
	// 检查部门名称是否可用
	var available bool
	available, err = s.IsDeptNameAvailable(ctx, req.Name, req.ParentId)
	if err != nil {
		return
	}
	if !available {
		return
	}
	// 写入部门数据
	user := Context().Get(ctx).User
	_, err = dao.Dept.Ctx(ctx).Data(do.Dept{
		ParentId:  req.ParentId,
		Name:      req.Name,
		Status:    req.Status,
		CreatedBy: user.Id,
	}).Insert()
	return
}

// GetDeptById 通过部门ID获取部门信息
func (s *sDept) GetDeptById(ctx context.Context, id uint64, fieldNames ...string) (dept *entity.Dept, err error) {
	model := dao.Dept.Ctx(ctx).Where(do.Dept{Id: id})
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
	}
	err = model.Scan(&dept)
	return
}

// EditDept 编辑部门
func (s *sDept) EditDept(ctx context.Context, req *v1.DeptEditReq) (err error) {
	// 检查父部门是否存在
	if req.ParentId != 0 {
		var deptExists bool
		deptExists, err = s.DeptExistsById(ctx, req.ParentId)
		if err != nil {
			return
		}
		if !deptExists {
			err = gerror.Newf(`父部门ID[%d]不存在`, req.ParentId)
			return
		}
	}
	// 检查部门名称是否可用
	var available bool
	available, err = s.IsDeptNameAvailable(ctx, req.Name, req.ParentId)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`部门名称[%s]已存在`, req.Name)
		return
	}
	// 获取部门列表
	var list []*entity.Dept
	err = dao.Dept.Ctx(ctx).Scan(&list)
	if err != nil {
		return err
	}
	// 获取部门ID下所有的子部门ID
	children := s.FindSonByParentId(list, req.Id)
	var idArray garray.Array
	for _, v := range children {
		idArray.Append(v.Id)
	}
	if idArray.Contains(req.ParentId) {
		err = gerror.Newf(`父部门ID[%d]是部门ID[%d]的子部门ID`, req.ParentId, req.Id)
		return
	}
	// 更新部门数据
	user := Context().Get(ctx).User
	_, err = dao.Dept.Ctx(ctx).Data(do.Dept{
		ParentId:  req.ParentId,
		Name:      req.Name,
		Status:    req.Status,
		UpdatedBy: user.Id,
	}).Where(do.Dept{Id: req.Id}).Update()
	return
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, id uint64) (err error) {
	// 获取部门列表
	var list []*entity.Dept
	err = dao.Dept.Ctx(ctx).Scan(&list)
	if err != nil {
		return err
	}
	// 获取所有的子部门信息
	children := s.FindSonByParentId(list, id)
	ids := make([]uint64, 0, len(list))
	for _, v := range children {
		ids = append(ids, v.Id)
	}
	ids = append(ids, id)
	// 删除部门数据
	_, err = dao.Dept.Ctx(ctx).WhereIn(dao.Dept.Columns().Id, ids).Delete()
	return
}

// GetDeptTree 获取部门树信息
func (s *sDept) GetDeptTree(deptList []*entity.Dept, parentId uint64) (tree []*v1.DeptTreeInfo) {
	tree = make([]*v1.DeptTreeInfo, 0, len(deptList))
	for _, v := range deptList {
		if v.ParentId == parentId {
			t := &v1.DeptTreeInfo{Dept: v}
			children := s.GetDeptTree(deptList, v.Id)
			if len(children) > 0 {
				t.Children = children
			}
			tree = append(tree, t)
		}
	}
	return
}

// IsDeptNameAvailable 部门名称是否可用
func (s *sDept) IsDeptNameAvailable(ctx context.Context, name string, parentId uint64) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Name: name, ParentId: parentId}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// DeptExistsById 通过部门ID判断部门信息是否存在
func (s *sDept) DeptExistsById(ctx context.Context, id uint64) (bool, error) {
	count, err := dao.Dept.Ctx(ctx).Where(do.Dept{Id: id}).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// FindSonByParentId 通过父部门ID获取所有的子部门信息
func (s *sDept) FindSonByParentId(deptList []*entity.Dept, parentId uint64) (children []*entity.Dept) {
	children = make([]*entity.Dept, 0, len(deptList))
	for _, v := range deptList {
		if v.ParentId == parentId {
			children = append(children, v)
			fChildren := s.FindSonByParentId(deptList, v.Id)
			children = append(children, fChildren...)
		}
	}
	return
}
