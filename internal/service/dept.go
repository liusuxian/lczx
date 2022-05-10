package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/consts"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/utils"
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
func (s *sDept) GetDeptList(ctx context.Context, req *v1.DeptListReq) (list []*entity.Dept, err error) {
	model := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	err = model.OrderAsc(columns.Id).Scan(&list)
	return
}

// AddDept 新增部门
func (s *sDept) AddDept(ctx context.Context, req *v1.DeptAddReq) (err error) {
	// 获取所有部门
	var allDepts []*entity.Dept
	allDepts, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}
	// 检查父部门是否存在
	parentDept := &entity.Dept{}
	if req.ParentId != 0 {
		parentDept = s.GetDeptById(allDepts, req.ParentId)
		if parentDept == nil {
			err = gerror.Newf(`父部门ID[%d]不存在`, req.ParentId)
			return
		}
		if parentDept.Status == 0 {
			err = gerror.Newf(`父部门ID[%d]已停用`, req.ParentId)
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
	// 保存部门数据
	err = s.saveDept(ctx, req)
	return
}

// EditDept 编辑部门
func (s *sDept) EditDept(ctx context.Context, req *v1.DeptEditReq) (err error) {
	// 获取所有部门
	var allDepts []*entity.Dept
	allDepts, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}
	// 检查父部门是否存在
	parentDept := &entity.Dept{}
	if req.ParentId != 0 {
		parentDept = s.GetDeptById(allDepts, req.ParentId)
		if parentDept == nil {
			err = gerror.Newf(`父部门ID[%d]不存在`, req.ParentId)
			return
		}
		if parentDept.Status == 0 {
			err = gerror.Newf(`父部门ID[%d]已停用`, req.ParentId)
			return
		}
	}
	if parentDept.Name == req.Name {
		err = gerror.Newf(`父部门名称[%s]不能与子部门名称[%s]相同`, parentDept.Name, req.Name)
		return
	}
	// 检查部门信息是否存在
	dept := s.GetDeptById(allDepts, req.Id)
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
	// 获取部门ID下所有的子部门ID
	idsMap := gmap.New()
	s.FindSonIdsByParentId(allDepts, dept.Id, idsMap)
	if idsMap.Contains(req.ParentId) {
		err = gerror.Newf(`父部门ID[%d]是部门ID[%d]的子部门ID`, req.ParentId, req.Id)
		return
	}
	// 更新部门数据
	err = s.updateDept(ctx, req, idsMap.Keys())
	return
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, ids []uint64) (err error) {
	// 获取所有部门
	var allDepts []*entity.Dept
	allDepts, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}
	// 获取所有的子部门ID
	idsMap := gmap.New()
	for _, id := range ids {
		idsMap.Set(id, true)
		s.FindSonIdsByParentId(allDepts, id, idsMap)
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
	var allDepts []*entity.Dept
	allDepts, err = s.GetAllDepts(ctx)
	if err != nil {
		return
	}

	depts = make([]*entity.Dept, 0, len(allDepts))
	for _, v := range allDepts {
		if v.Status == 1 {
			depts = append(depts, v)
		}
	}
	return
}

// GetAllDepts 获取所有部门
func (s *sDept) GetAllDepts(ctx context.Context) (depts []*entity.Dept, err error) {
	// 从缓存获取
	deptsCacheVal := Cache().GetCache(ctx, consts.DeptKey)
	if deptsCacheVal != nil {
		err = gconv.Structs(deptsCacheVal, &depts)
		if err != nil {
			return
		}
		if depts != nil {
			return
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

// GetDeptIdsByRoleId 获取角色ID关联的部门ID列表
func (s *sDept) GetDeptIdsByRoleId(ctx context.Context, id uint64) (deptIds []uint64, err error) {
	var array []*gvar.Var
	array, err = dao.RoleDept.Ctx(ctx).Fields(dao.RoleDept.Columns().DeptId).Where(do.RoleDept{RoleId: id}).Array()
	deptIds = make([]uint64, 0, len(array))
	for _, deptVar := range array {
		deptIds = append(deptIds, deptVar.Uint64())
	}
	return
}

// GetDeptAllNameById 通过部门ID获取部门名称全称
func (s *sDept) GetDeptAllNameById(deptList []*entity.Dept, id uint64) (deptAllName string) {
	deptNameList := make([]string, 0, len(deptList))
	s.GetDeptNameListById(deptList, id, &deptNameList)
	utils.Reverse(deptNameList)
	deptAllName = gstr.Join(deptNameList, "->")
	return
}

// GetDeptNameListById 通过部门ID获取部门名称列表
func (s *sDept) GetDeptNameListById(deptList []*entity.Dept, id uint64, deptNameList *[]string) {
	for _, v := range deptList {
		if v.Id == id {
			*deptNameList = append(*deptNameList, v.Name)
			s.GetDeptNameListById(deptList, v.ParentId, deptNameList)
		}
	}
}

// CopyDept 拷贝部门信息
func (s *sDept) CopyDept(srcDept *entity.Dept) (dstDept *entity.Dept) {
	dstDept = new(entity.Dept)
	dstDept.Id = srcDept.Id
	dstDept.ParentId = srcDept.ParentId
	dstDept.Name = srcDept.Name
	dstDept.Status = srcDept.Status
	dstDept.CreatedBy = srcDept.CreatedBy
	dstDept.UpdatedBy = srcDept.UpdatedBy
	dstDept.CreateAt = srcDept.CreateAt
	dstDept.UpdateAt = srcDept.UpdateAt
	dstDept.DeletedAt = srcDept.DeletedAt
	return
}

// GetDeptById 通过部门ID获取部门信息
func (s *sDept) GetDeptById(deptList []*entity.Dept, id uint64) (dept *entity.Dept) {
	for _, d := range deptList {
		if d.Id == id {
			dept = d
			return
		}
	}
	return
}

// SelectDeptById 通过部门ID查询部门信息
func (s *sDept) SelectDeptById(ctx context.Context, id uint64) (dept *entity.Dept, err error) {
	err = dao.Dept.Ctx(ctx).Where(do.Dept{Id: id}).Scan(&dept)
	return
}

// saveDept 保存部门数据
func (s *sDept) saveDept(ctx context.Context, req *v1.DeptAddReq) (err error) {
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
	}).FieldsEx(dao.Dept.Columns().Id).Insert()
	return
}

// updateDept 更新部门数据
func (s *sDept) updateDept(ctx context.Context, req *v1.DeptEditReq, ids []any) (err error) {
	user := Context().Get(ctx).User
	err = dao.Dept.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var terr error
		_, terr = dao.Dept.Ctx(ctx).Data(do.Dept{
			ParentId:  req.ParentId,
			Name:      req.Name,
			Status:    req.Status,
			UpdatedBy: user.Id,
		}).Where(do.Dept{Id: req.Id}).Update()
		if terr != nil {
			return terr
		}

		if req.Status == 0 {
			_, terr = dao.Dept.Ctx(ctx).Data(do.Dept{
				Status:    req.Status,
				UpdatedBy: user.Id,
			}).WhereIn(dao.Dept.Columns().Id, ids).Update()
			if terr != nil {
				return terr
			}
		}
		return nil
	})
	if err != nil {
		return
	}
	// 清除部门缓存
	err = Cache().ClearCache(ctx, consts.DeptKey)
	return
}
