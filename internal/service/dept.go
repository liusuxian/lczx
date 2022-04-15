package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
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
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if len(fieldNames) != 0 {
		model = model.FieldsEx(fieldNames)
	}
	err = model.OrderAsc(columns.Id).Scan(&list)
	return
}

// AddDept 新增部门
func (s *sDept) AddDept(ctx context.Context, req *v1.DeptAddReq) (err error) {
	// 检查父部门是否存在
	parentDept := &entity.Dept{}
	if req.ParentId != 0 {
		parentDept, err = s.GetDeptById(ctx, req.ParentId)
		if err != nil {
			return
		}
		if parentDept == nil {
			err = gerror.Newf(`父部门ID[%d]不存在`, req.ParentId)
			return
		}
	}
	if parentDept.Name == req.Name {
		err = gerror.Newf(`父部门名称[%s]不能与子部门名称[%s]相同`, parentDept.Name, req.Name)
		return
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
	// 写入部门数据
	user := Context().Get(ctx).User
	_, err = dao.Dept.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.DeptKey,
		Force:    false,
	}).Data(do.Dept{
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
	parentDept := &entity.Dept{}
	if req.ParentId != 0 {
		parentDept, err = s.GetDeptById(ctx, req.ParentId)
		if err != nil {
			return
		}
		if parentDept == nil {
			err = gerror.Newf(`父部门ID[%d]不存在`, req.ParentId)
			return
		}
	}
	if parentDept.Name == req.Name {
		err = gerror.Newf(`父部门名称[%s]不能与子部门名称[%s]相同`, parentDept.Name, req.Name)
		return
	}
	// 检查部门信息是否存在
	var dept *entity.Dept
	dept, err = s.GetDeptById(ctx, req.Id)
	if err != nil {
		return
	}
	if dept == nil {
		err = gerror.Newf(`部门ID[%d]不存在`, req.Id)
		return
	}
	// 检查部门名称是否可用
	if dept.Name != req.Name {
		var available bool
		available, err = s.IsDeptNameAvailable(ctx, req.Name, req.ParentId)
		if err != nil {
			return
		}
		if !available {
			err = gerror.Newf(`部门名称[%s]已存在`, req.Name)
			return
		}
	}
	// 获取所有部门
	var list []*entity.Dept
	list, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}
	// 获取部门ID下所有的子部门ID
	idsMap := gmap.New(true)
	s.FindSonIdsByParentId(list, req.Id, idsMap)
	if idsMap.Contains(req.ParentId) {
		err = gerror.Newf(`父部门ID[%d]是部门ID[%d]的子部门ID`, req.ParentId, req.Id)
		return
	}
	// 更新部门数据
	user := Context().Get(ctx).User
	_, err = dao.Dept.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.DeptKey,
		Force:    false,
	}).Data(do.Dept{
		ParentId:  req.ParentId,
		Name:      req.Name,
		Status:    req.Status,
		UpdatedBy: user.Id,
	}).Where(do.Dept{Id: req.Id}).Update()
	return
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, ids []uint64) (err error) {
	// 获取所有部门
	var list []*entity.Dept
	list, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}
	// 获取所有的子部门ID
	idsMap := gmap.New(true)
	for _, id := range ids {
		idsMap.Set(id, true)
		s.FindSonIdsByParentId(list, id, idsMap)
	}
	delIds := idsMap.Keys()
	// 删除部门数据
	_, err = dao.Dept.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     consts.DeptKey,
		Force:    false,
	}).WhereIn(dao.Dept.Columns().Id, delIds).Delete()
	return
}

// GetStatusEnableDepts 获取部门状态为正常的部门列表
func (s *sDept) GetStatusEnableDepts(ctx context.Context) (depts []*entity.Dept, err error) {
	// 获取所有部门
	var list []*entity.Dept
	list, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}

	depts = make([]*entity.Dept, 0, len(list))
	for _, v := range list {
		if v.Status == consts.DeptStatusEnable {
			depts = append(depts, v)
		}
	}
	return
}

// GetAllDepts 获取所有部门
func (s *sDept) GetAllDepts(ctx context.Context) (depts []*entity.Dept, err error) {
	// 从缓存获取
	var deptsCacheValue *gvar.Var
	deptsCacheValue, err = g.Redis().Do(ctx, "GET", consts.DeptKey)
	if err != nil {
		return
	}
	if deptsCacheValue != nil {
		deptsCacheMap := deptsCacheValue.Map()["Result"]
		if deptsCacheMap != nil {
			err = gconv.Structs(deptsCacheMap, &depts)
			if err != nil {
				return
			}
			if depts != nil {
				return
			}
		}
	}
	// 从数据库获取
	err = dao.Dept.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     consts.DeptKey,
		Force:    false,
	}).OrderAsc(dao.Dept.Columns().Id).Scan(&depts)
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

// FindSonIdsByParentId 通过父部门ID获取所有的子部门ID
func (s *sDept) FindSonIdsByParentId(deptList []*entity.Dept, parentId uint64, idsMap *gmap.Map) {
	for _, v := range deptList {
		if v.ParentId == parentId {
			idsMap.Set(v.Id, true)
			s.FindSonIdsByParentId(deptList, v.Id, idsMap)
		}
	}
}
