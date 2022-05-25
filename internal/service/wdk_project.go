package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
func (s *sWdkProject) GetWdkProjectList(ctx context.Context, req *v1.WdkProjectListReq) (total int, list []*v1.WdkProjectInfo, err error) {
	// 处理业态
	projectIdsMap := gmap.New()
	if len(req.BusinessForms) != 0 {
		var wdkProjectBusinessforms []*entity.WdkProjectBusinessforms
		err = dao.WdkProjectBusinessforms.Ctx(ctx).WhereIn(dao.WdkProjectBusinessforms.Columns().BusinessForms, req.BusinessForms).
			Scan(&wdkProjectBusinessforms)
		if err != nil {
			return
		}
		for _, v := range wdkProjectBusinessforms {
			projectIdsMap.Set(v.ProjectId, true)
		}
		if projectIdsMap.IsEmpty() {
			return
		}
	}
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
	if !projectIdsMap.IsEmpty() {
		model = model.WhereIn(columns.Id, projectIdsMap.Keys())
	}
	if req.ContractStatus != "" {
		model = model.Where(columns.ContractStatus, gconv.Uint(req.ContractStatus))
	}
	if req.ContractSum != "" {
		model = model.WhereLike(columns.ContractSum, "%"+req.ContractSum+"%")
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
		startTime := gstr.Split(req.StartTime.String(), " ")[0]
		model = model.WhereGTE(columns.StartTime, startTime)
	}
	if req.EndTime.String() != "" {
		endTime := gstr.Split(req.EndTime.String(), " ")[0]
		model = model.WhereLTE(columns.EndTime, endTime)
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
	err = model.Page(req.CurPage, req.PageSize).Order(order).ScanList(&list, "ProjectInfo")
	if err != nil {
		return
	}
	err = dao.WdkProjectBusinessforms.Ctx(ctx).Where(dao.WdkProjectBusinessforms.Columns().ProjectId, gdb.ListItemValuesUnique(list, "ProjectInfo", "Id")).
		ScanList(&list, "Businessforms", "ProjectInfo", "ProjectId:Id")
	return
}

// AddWdkProject 新增文档库项目
func (s *sWdkProject) AddWdkProject(ctx context.Context, req *v1.WdkProjectAddReq) (err error) {
	// 检查项目时间
	if req.StartTime.Timestamp() >= req.EndTime.Timestamp() {
		err = gerror.Newf(`项目开始时间[%s]大于等于项目结束时间[%s]`, req.StartTime.String(), req.EndTime.String())
		return
	}
	err = dao.WdkProject.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查文档库项目名称是否可用
		var terr error
		var available bool
		available, terr = s.IsWdkProjectNameAvailable(ctx, req.Name)
		if terr != nil {
			return terr
		}
		if !available {
			return gerror.Newf(`文档库项目名称[%s]已存在或已被使用过`, req.Name)
		}
		// 检查负责人是否存在
		var principalUser *entity.User
		principalUser, terr = User().GetUserById(ctx, req.PrincipalUid)
		if terr != nil {
			return terr
		}
		if principalUser == nil {
			return gerror.Newf(`负责人用户ID[%d]不存在`, req.PrincipalUid)
		}
		// 通过部门ID判断部门信息是否存在
		var dept *entity.Dept
		dept, terr = Dept().SelectDeptById(ctx, req.DeptId)
		if terr != nil {
			return terr
		}
		if dept == nil {
			return gerror.Newf(`所属部门ID[%d]不存在`, req.DeptId)
		}
		// 保存文档库项目数据
		terr = s.saveWdkProject(ctx, req, principalUser, dept)
		return terr
	})
	return
}

// GetWdkProjectById 通过文档库项目ID获取文档库项目信息
func (s *sWdkProject) GetWdkProjectById(ctx context.Context, id uint64) (wdkProject *v1.WdkProjectInfo, err error) {
	wdkProject = &v1.WdkProjectInfo{}
	err = dao.WdkProject.Ctx(ctx).Where(do.WdkProject{Id: id}).Scan(&wdkProject.ProjectInfo)
	if err != nil {
		return
	}
	err = dao.WdkProjectBusinessforms.Ctx(ctx).Where(do.WdkProjectBusinessforms{ProjectId: id}).Scan(&wdkProject.Businessforms)
	return
}

// EditWdkProject 编辑文档库项目
func (s *sWdkProject) EditWdkProject(ctx context.Context, req *v1.WdkProjectEditReq) (err error) {
	// 检查项目时间
	if req.StartTime.Timestamp() >= req.EndTime.Timestamp() {
		err = gerror.Newf(`项目开始时间[%s]大于等于项目结束时间[%s]`, req.StartTime.String(), req.EndTime.String())
		return
	}
	err = dao.WdkProject.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查文档库项目信息是否存在
		var terr error
		var wdkProject *v1.WdkProjectInfo
		wdkProject, terr = s.GetWdkProjectById(ctx, req.Id)
		if terr != nil {
			return terr
		}
		if wdkProject == nil || wdkProject.ProjectInfo == nil {
			return gerror.Newf(`文档库项目ID[%d]不存在`, req.Id)
		}
		// 检查文档库项目名称是否可用
		if wdkProject.ProjectInfo.Name != req.Name {
			var available bool
			available, terr = s.IsWdkProjectNameAvailable(ctx, req.Name)
			if terr != nil {
				return terr
			}
			if !available {
				return gerror.Newf(`文档库项目名称[%s]已存在或已被使用过`, req.Name)
			}
		}
		// 检查负责人是否存在
		var principalUser *entity.User
		principalUser, terr = User().GetUserById(ctx, req.PrincipalUid)
		if terr != nil {
			return terr
		}
		if principalUser == nil {
			return gerror.Newf(`负责人用户ID[%d]不存在`, req.PrincipalUid)
		}
		// 通过部门ID判断部门信息是否存在
		var dept *entity.Dept
		dept, terr = Dept().SelectDeptById(ctx, req.DeptId)
		if terr != nil {
			return terr
		}
		if dept == nil {
			return gerror.Newf(`所属部门ID[%d]不存在`, req.DeptId)
		}
		// 更新文档库项目数据
		terr = s.updateWdkProject(ctx, req, principalUser, dept, wdkProject)
		return terr
	})
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

// SetWdkProjectFileUploadStatus 设置文档库项目文件上传状态为是
func (s *sWdkProject) SetWdkProjectFileUploadStatus(ctx context.Context, id uint64) (err error) {
	var fileCount int
	fileCount, err = WdkFile().GetWdkFileCountByProjectId(ctx, id)
	if err != nil {
		return err
	}
	var reportCount int
	reportCount, err = WdkReport().GetWdkReportCountByProjectId(ctx, id)
	if err != nil {
		return err
	}
	var reportCfgCount int
	reportCfgCount, err = WdkReportCfg().GetWdkReportCfgCount(ctx)
	if err != nil {
		return err
	}
	if (fileCount + reportCount) >= reportCfgCount+7 {
		_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{FileUploadStatus: 1}).Where(do.WdkProject{Id: id}).Update()
		if err != nil {
			return err
		}
	}
	return
}

// SetWdkProjectStep 设置文档库项目阶段
func (s *sWdkProject) SetWdkProjectStep(ctx context.Context, id uint64, fType uint) (err error) {
	// 文件类型 0:合同扫描文件 1:年度服务计划书 2:总结报告 3:项目移交 4:复盘报告 5:文件签收单 6:满意度调查表 7:服务记录 8:咨询报告
	// 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	step := 0
	switch fType {
	case 0:
		step = 1
	case 1:
		step = 2
	case 2, 3:
		step = 4
	case 4:
		step = 5
	case 5, 6, 7, 8:
		step = 3
	}
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{Step: step}).Where(do.WdkProject{Id: id}).
		WhereLT(dao.WdkProject.Columns().Step, step).Update()
	return
}

// saveWdkProject 保存文档库项目数据
func (s *sWdkProject) saveWdkProject(ctx context.Context, req *v1.WdkProjectAddReq, principalUser *entity.User, dept *entity.Dept) (err error) {
	user := Context().Get(ctx).User
	var projectId int64
	projectId, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{
		Name:             req.Name,
		Type:             req.Type,
		Origin:           req.Origin,
		Step:             0,
		FileUploadStatus: 0,
		BusinessType:     req.BusinessType,
		ContractStatus:   req.ContractStatus,
		ContractSum:      req.ContractSum,
		DeepCulture:      req.DeepCulture,
		Status:           req.Status,
		EntrustCompany:   req.EntrustCompany,
		SignCompany:      req.SignCompany,
		PrincipalUid:     principalUser.Id,
		PrincipalName:    principalUser.Realname,
		DeptId:           dept.Id,
		DeptName:         dept.Name,
		Region:           req.Region,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		CreateBy:         user.Id,
		CreateName:       user.Realname,
		Remark:           req.Remark,
	}).FieldsEx(dao.WdkProject.Columns().Id).InsertAndGetId()
	if err != nil {
		return
	}
	// 保存文档库项目业态类型数据
	if len(req.BusinessForms) != 0 {
		businessFormsData := g.List{}
		for _, v := range req.BusinessForms {
			businessFormsData = append(businessFormsData, g.Map{
				"project_id":     projectId,
				"business_forms": v,
			})
		}
		_, err = dao.WdkProjectBusinessforms.Ctx(ctx).Data(businessFormsData).Batch(len(businessFormsData)).Insert()
		if err != nil {
			return
		}
	}
	return nil
}

// updateWdkProject 更新文档库项目数据
func (s *sWdkProject) updateWdkProject(ctx context.Context, req *v1.WdkProjectEditReq, principalUser *entity.User, dept *entity.Dept, wdkProject *v1.WdkProjectInfo) (err error) {
	user := Context().Get(ctx).User
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{
		Name:           req.Name,
		Type:           req.Type,
		Origin:         req.Origin,
		BusinessType:   req.BusinessType,
		ContractStatus: req.ContractStatus,
		ContractSum:    req.ContractSum,
		DeepCulture:    req.DeepCulture,
		Status:         req.Status,
		EntrustCompany: req.EntrustCompany,
		SignCompany:    req.SignCompany,
		PrincipalUid:   principalUser.Id,
		PrincipalName:  principalUser.Realname,
		DeptId:         dept.Id,
		DeptName:       dept.Name,
		Region:         req.Region,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		UpdatedBy:      user.Id,
		UpdatedName:    user.Realname,
		Remark:         req.Remark,
	}).Where(do.WdkProject{Id: req.Id}).Update()
	if err != nil {
		return
	}
	// 删除旧的文档库项目业态类型数据
	_, err = dao.WdkProjectBusinessforms.Ctx(ctx).Where(do.WdkProjectBusinessforms{ProjectId: wdkProject.ProjectInfo.Id}).Delete()
	if err != nil {
		return
	}
	// 保存文档库项目业态类型数据
	if len(req.BusinessForms) != 0 {
		businessFormsData := g.List{}
		for _, v := range req.BusinessForms {
			businessFormsData = append(businessFormsData, g.Map{
				"project_id":     wdkProject.ProjectInfo.Id,
				"business_forms": v,
			})
		}
		_, err = dao.WdkProjectBusinessforms.Ctx(ctx).Data(businessFormsData).Batch(len(businessFormsData)).Insert()
		if err != nil {
			return
		}
	}
	return nil
}
