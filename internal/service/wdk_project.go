package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sWdkProject struct{}

var (
	insWdkProject = sWdkProject{}
)

// WdkProject 文档库项目管理服务
func WdkProject() *sWdkProject {
	return &insWdkProject
}

// GetWdkProjectList 获取文档库项目列表
func (s *sWdkProject) GetWdkProjectList(ctx context.Context, req *v1.WdkProjectListReq) (total int, list []*entity.WdkProject, err error) {
	model := dao.WdkProject.Ctx(ctx)
	columns := dao.WdkProject.Columns()
	order := "id DESC"
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Type != "" {
		model = model.Where(columns.Type, gconv.Uint(req.Type))
	}
	if req.Origin != "" {
		model = model.Where(columns.Origin, gconv.Uint(req.Origin))
	}
	if req.Step != "" {
		model = model.Where(columns.Step, gconv.Uint(req.Step))
	}
	if req.FileUploadStatus != "" {
		model = model.Where(columns.FileUploadStatus, gconv.Uint(req.FileUploadStatus))
	}
	if req.BusinessType != "" {
		model = model.Where(columns.BusinessType, gconv.Uint(req.BusinessType))
	}
	if req.DeepCulture != "" {
		model = model.Where(columns.DeepCulture, gconv.Uint(req.DeepCulture))
	}
	if req.Status != "" {
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if req.EntrustCompany != "" {
		model = model.WhereLike(columns.EntrustCompany, "%"+req.EntrustCompany+"%")
	}
	if req.SignCompany != "" {
		model = model.Where(columns.SignCompany, gconv.Uint(req.SignCompany))
	}
	if req.PrincipalName != "" {
		model = model.WhereLike(columns.PrincipalName, "%"+req.PrincipalName+"%")
	}
	deptIdsMap := gmap.New()
	if req.DeptId != "" {
		// 获取部门状态为正常的部门列表
		var depts []*entity.Dept
		depts, err = Dept().GetStatusEnableDepts(ctx)
		if err != nil {
			return
		}
		deptId := gconv.Uint64(req.DeptId)
		deptIdsMap.Set(deptId, true)
		Dept().FindSonIdsByParentId(depts, deptId, deptIdsMap)
	}
	if !deptIdsMap.IsEmpty() {
		model = model.WhereIn(columns.DeptId, deptIdsMap.Keys())
	}
	if req.Region != "" {
		model = model.WhereLike(columns.Region, "%"+req.Region+"%")
	}
	if req.StartTime.String() != "" {
		model = model.WhereGTE(columns.StartTime, req.StartTime)
	}
	if req.EndTime.String() != "" {
		model = model.WhereLTE(columns.EndTime, req.EndTime)
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	return
}

// AddWdkProject 新增文档库项目
func (s *sWdkProject) AddWdkProject(ctx context.Context, req *v1.WdkProjectAddReq) (err error) {
	// 检查项目时间
	if req.StartTime.Timestamp() >= req.EndTime.Timestamp() {
		err = gerror.Newf(`项目开始时间[%s]大于等于项目结束时间[%s]`, req.StartTime.String(), req.EndTime.String())
		return
	}
	// 检查文档库项目名称是否可用
	var available bool
	available, err = s.IsWdkProjectNameAvailable(ctx, req.Name)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`文档库项目名称[%s]已存在`, req.Name)
		return
	}
	// 检查负责人是否存在
	var principalUser *entity.User
	principalUser, err = User().GetUserById(ctx, req.PrincipalUid)
	if err != nil {
		return
	}
	if principalUser == nil {
		err = gerror.Newf(`负责人用户ID[%d]不存在`, req.PrincipalUid)
		return
	}
	// 检查项目负责人所属部门和项目所属部门是否一致
	if req.DeptId != principalUser.DeptId {
		err = gerror.Newf(`项目负责人所属部门ID[%d]与项目所属部门ID[%d]不一致`, principalUser.DeptId, req.DeptId)
		return
	}
	// 写入文档库项目数据
	user := Context().Get(ctx).User
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{
		Name:             req.Name,
		Type:             req.Type,
		Origin:           req.Origin,
		Step:             0,
		FileUploadStatus: 0,
		BusinessType:     req.BusinessType,
		DeepCulture:      req.DeepCulture,
		Status:           0,
		EntrustCompany:   req.EntrustCompany,
		SignCompany:      req.SignCompany,
		PrincipalUid:     req.PrincipalUid,
		PrincipalName:    principalUser.Realname,
		DeptId:           req.DeptId,
		Region:           req.Region,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		CreateBy:         user.Id,
		Remark:           req.Remark,
	}).Insert()
	return
}

// EditWdkProject 编辑文档库项目
func (s *sWdkProject) EditWdkProject(ctx context.Context, req *v1.WdkProjectEditReq) (err error) {
	// 检查项目时间
	if req.StartTime.Timestamp() >= req.EndTime.Timestamp() {
		err = gerror.Newf(`项目开始时间[%s]大于等于项目结束时间[%s]`, req.StartTime.String(), req.EndTime.String())
		return
	}
	// 检查文档库项目信息是否存在
	var wdkProject *entity.WdkProject
	wdkProject, err = s.GetWdkProjectById(ctx, req.Id)
	if err != nil {
		return
	}
	if wdkProject == nil {
		err = gerror.Newf(`文档库项目ID[%d]不存在`, req.Id)
		return
	}
	// 检查文档库项目名称是否可用
	if wdkProject.Name != req.Name {
		var available bool
		available, err = s.IsWdkProjectNameAvailable(ctx, req.Name)
		if err != nil {
			return
		}
		if !available {
			err = gerror.Newf(`文档库项目名称[%s]已存在`, req.Name)
			return
		}
	}
	// 检查负责人是否存在
	var principalUser *entity.User
	principalUser, err = User().GetUserById(ctx, req.PrincipalUid)
	if err != nil {
		return
	}
	if principalUser == nil {
		err = gerror.Newf(`负责人用户ID[%d]不存在`, req.PrincipalUid)
		return
	}
	// 检查项目负责人所属部门和项目所属部门是否一致
	if req.DeptId != principalUser.DeptId {
		err = gerror.Newf(`项目负责人所属部门ID[%d]与项目所属部门ID[%d]不一致`, principalUser.DeptId, req.DeptId)
		return
	}
	// 更新文档库项目数据
	user := Context().Get(ctx).User
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{
		Name:           req.Name,
		Type:           req.Type,
		Origin:         req.Origin,
		BusinessType:   req.BusinessType,
		DeepCulture:    req.DeepCulture,
		Status:         req.Status,
		EntrustCompany: req.EntrustCompany,
		SignCompany:    req.SignCompany,
		PrincipalUid:   req.PrincipalUid,
		PrincipalName:  principalUser.Realname,
		DeptId:         req.DeptId,
		Region:         req.Region,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		UpdatedBy:      user.Id,
		Remark:         req.Remark,
	}).Where(do.WdkProject{Id: req.Id}).Update()
	return
}

// DeleteWdkProject 删除文档库项目
func (s *sWdkProject) DeleteWdkProject(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.WdkProject.Ctx(ctx).WhereIn(dao.WdkProject.Columns().Id, ids).Delete()
	return
}

// IsWdkProjectNameAvailable 文档库项目名称是否可用
func (s *sWdkProject) IsWdkProjectNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.WdkProject.Ctx(ctx).Where(do.WdkProject{Name: name}).Unscoped().Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetWdkProjectById 通过文档库项目ID获取文档库项目信息
func (s *sWdkProject) GetWdkProjectById(ctx context.Context, id uint64) (wdkProject *entity.WdkProject, err error) {
	err = dao.WdkProject.Ctx(ctx).Where(do.WdkProject{Id: id}).Scan(&wdkProject)
	return
}
