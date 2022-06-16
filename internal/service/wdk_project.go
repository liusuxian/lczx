package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/utils"
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
		startTime := req.StartTime.Format("Y-m-d")
		model = model.WhereGTE(columns.StartTime, startTime)
	}
	if req.EndTime.String() != "" {
		endTime := req.EndTime.Format("Y-m-d")
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

// ExportWdkProject 导出文档库项目信息
func (s *sWdkProject) ExportWdkProject(ctx context.Context, req *v1.WdkProjectExportReq) (fileInfo *v1.WdkProjectExportFile, err error) {
	// 处理业态
	projectIdsMap := gmap.New()
	if len(req.BusinessForms) != 0 {
		var wdkProjectBusinessforms []*entity.WdkProjectBusinessforms
		if err = dao.WdkProjectBusinessforms.Ctx(ctx).WhereIn(dao.WdkProjectBusinessforms.Columns().BusinessForms, req.BusinessForms).
			Scan(&wdkProjectBusinessforms); err != nil {
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
		if depts, err = Dept().GetStatusEnableDepts(ctx); err != nil {
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
		startTime := req.StartTime.Format("Y-m-d")
		model = model.WhereGTE(columns.StartTime, startTime)
	}
	if req.EndTime.String() != "" {
		endTime := req.EndTime.Format("Y-m-d")
		model = model.WhereLTE(columns.EndTime, endTime)
	}
	var total int
	if total, err = model.Count(); err != nil {
		return
	}
	// 循环读取
	curPage := 1
	pageSize := 50
	excelData := make([][]any, 0, total+2)
	excelData = append(excelData, []any{"项目信息"})
	excelData = append(excelData, []any{"项目ID", "项目名称", "项目性质", "项目来源", "项目阶段", "上传状态", "业务类型", "业态", "签约状态",
		"合同金额（单位：万元）", "是否深耕", "服务状态", "委托方公司", "签订公司", "负责人", "所属部门", "地区", "开始时间", "结束时间", "备注"})
	for {
		var list []*v1.WdkProjectInfo
		if err = model.Page(curPage, pageSize).Order(order).ScanList(&list, "ProjectInfo"); err != nil {
			return
		}
		if err = dao.WdkProjectBusinessforms.Ctx(ctx).Where(dao.WdkProjectBusinessforms.Columns().ProjectId, gdb.ListItemValuesUnique(list, "ProjectInfo", "Id")).
			ScanList(&list, "Businessforms", "ProjectInfo", "ProjectId:Id"); err != nil {
			return
		}
		if list == nil {
			break
		}
		for _, v := range list {
			excelData = append(excelData, []any{
				v.ProjectInfo.Id,
				v.ProjectInfo.Name,
				s.getWdkProjectType(v.ProjectInfo.Type),
				s.getWdkProjectOrigin(v.ProjectInfo.Origin),
				s.getWdkProjectStep(v.ProjectInfo.Step),
				s.getWdkProjectFileUploadStatus(v.ProjectInfo.FileUploadStatus),
				s.getWdkProjectBusinessType(v.ProjectInfo.BusinessType),
				s.getWdkProjectBusinessforms(v.Businessforms),
				s.getWdkProjectContractStatus(v.ProjectInfo.ContractStatus),
				v.ProjectInfo.ContractSum,
				s.getWdkProjectDeepCulture(v.ProjectInfo.DeepCulture),
				s.getWdkProjectStatus(v.ProjectInfo.Status),
				v.ProjectInfo.EntrustCompany,
				s.getWdkProjectSignCompany(v.ProjectInfo.SignCompany),
				v.ProjectInfo.PrincipalName,
				v.ProjectInfo.DeptName,
				v.ProjectInfo.Region,
				v.ProjectInfo.StartTime.Format("Y-m-d"),
				v.ProjectInfo.EndTime.Format("Y-m-d"),
				v.ProjectInfo.Remark,
			})
		}
		if curPage*pageSize >= total {
			break
		}
		curPage++
	}
	// 创建文档库项目信息Excel表
	if fileInfo, err = s.createWdkProjectExcel(excelData); err != nil {
		return
	}
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

// SetWdkProjectStepByFileType 通过文件类型设置文档库项目阶段
// 文件类型 0:合同扫描文件 1:年度服务计划书 2:总结报告 3:项目移交 4:复盘报告 5:文件签收单 6:满意度调查表
// 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程
// 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘
func (s *sWdkProject) SetWdkProjectStepByFileType(ctx context.Context, id uint64, fileType uint) (err error) {
	step := 0
	switch fileType {
	case 0:
		step = 1
	case 1:
		step = 2
	case 2, 3:
		step = 30
	case 4:
		step = 31
	case 5, 6:
		return
	default:
		step = 0
	}
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{Step: step}).Where(do.WdkProject{Id: id}).
		WhereLT(dao.WdkProject.Columns().Step, step).Update()
	return
}

// SetWdkProjectStepByReportStep 通过报告阶段设置文档库项目阶段
// 项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程
// 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘
func (s *sWdkProject) SetWdkProjectStepByReportStep(ctx context.Context, id uint64, reportStep uint) (err error) {
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{Step: reportStep}).Where(do.WdkProject{Id: id}).
		WhereLT(dao.WdkProject.Columns().Step, reportStep).Update()
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

// getWdkProjectType 获取项目性质中文名称
func (s *sWdkProject) getWdkProjectType(ptype uint) (name string) {
	switch ptype {
	case 0:
		name = "蓝绿体系"
	case 1:
		name = "非绿"
	default:
		name = ""
	}
	return
}

// getWdkProjectOrigin 获取项目来源中文名称
func (s *sWdkProject) getWdkProjectOrigin(origin uint) (name string) {
	switch origin {
	case 0:
		name = "绿中"
	case 1:
		name = "分子公司"
	case 2:
		name = "合伙人"
	case 3:
		name = "老客户"
	case 4:
		name = "中交"
	case 5:
		name = "蓝城"
	case 6:
		name = "自拓"
	case 7:
		name = "其他"
	default:
		name = ""
	}
	return
}

// getWdkProjectStep 获取项目阶段中文名称
func (s *sWdkProject) getWdkProjectStep(step uint) (name string) {
	switch step {
	case 0:
		name = "未开始"
	case 1:
		name = "合同签约"
	case 2:
		name = "项目启动会"
	case 3:
		name = "服务中-规划设计"
	case 4:
		name = "服务中-项目展示区施工"
	case 5:
		name = "服务中-主体结构工程"
	case 6:
		name = "服务中-主体安装工程"
	case 7:
		name = "服务中-装饰装修工程"
	case 8:
		name = "服务中-景观市政工程"
	case 9:
		name = "服务中-项目交付验收"
	case 30:
		name = "合同结束"
	case 31:
		name = "复盘"
	default:
		name = ""
	}
	return
}

// getWdkProjectFileUploadStatus 获取项目文件上传状态中文名称
func (s *sWdkProject) getWdkProjectFileUploadStatus(fileUploadStatus uint) (name string) {
	switch fileUploadStatus {
	case 0:
		name = "异常"
	case 1:
		name = "正常"
	case 2:
		name = "已完成"
	default:
		name = ""
	}
	return
}

// getWdkProjectBusinessType 获取业务类型中文名称
func (s *sWdkProject) getWdkProjectBusinessType(businessType uint) (name string) {
	switch businessType {
	case 0:
		name = "物业"
	case 1:
		name = "专项"
	case 2:
		name = "全过程"
	default:
		name = ""
	}
	return
}

// getWdkProjectBusinessforms 获取项目业态中文名称
func (s *sWdkProject) getWdkProjectBusinessforms(businessforms []*entity.WdkProjectBusinessforms) (name string) {
	nameList := make([]string, 0, len(businessforms))
	for _, v := range businessforms {
		switch v.BusinessForms {
		case 0:
			nameList = append(nameList, "住宅")
		case 1:
			nameList = append(nameList, "小高层")
		case 2:
			nameList = append(nameList, "高层")
		case 3:
			nameList = append(nameList, "超高层")
		case 4:
			nameList = append(nameList, "公寓")
		case 5:
			nameList = append(nameList, "合院")
		case 6:
			nameList = append(nameList, "叠墅")
		case 7:
			nameList = append(nameList, "排屋")
		case 8:
			nameList = append(nameList, "多层")
		case 9:
			nameList = append(nameList, "会所")
		case 10:
			nameList = append(nameList, "商住")
		case 11:
			nameList = append(nameList, "综合体")
		case 12:
			nameList = append(nameList, "产业园")
		case 13:
			nameList = append(nameList, "酒店")
		case 14:
			nameList = append(nameList, "酒店式公寓")
		case 15:
			nameList = append(nameList, "商业")
		case 16:
			nameList = append(nameList, "普通商业")
		case 17:
			nameList = append(nameList, "公共配套")
		case 18:
			nameList = append(nameList, "办公")
		case 19:
			nameList = append(nameList, "厂房")
		default:
			nameList = append(nameList, "")
		}
	}
	name = gstr.Join(nameList, ", ")
	return
}

// getWdkProjectContractStatus 获取签约状态中文名称
func (s *sWdkProject) getWdkProjectContractStatus(contractStatus uint) (name string) {
	switch contractStatus {
	case 0:
		name = "新签"
	case 1:
		name = "续签"
	case 2:
		name = "未签"
	default:
		name = ""
	}
	return
}

// getWdkProjectDeepCulture 获取是否为深耕中文名称
func (s *sWdkProject) getWdkProjectDeepCulture(deepCulture uint) (name string) {
	switch deepCulture {
	case 0:
		name = "否"
	case 1:
		name = "是"
	default:
		name = ""
	}
	return
}

// getWdkProjectStatus 获取服务状态中文名称
func (s *sWdkProject) getWdkProjectStatus(status uint) (name string) {
	switch status {
	case 0:
		name = "服务中"
	case 1:
		name = "暂停"
	case 2:
		name = "提前终止"
	case 3:
		name = "跟踪期"
	case 4:
		name = "洽谈中"
	default:
		name = ""
	}
	return
}

// getWdkProjectSignCompany 获取我方签订公司中文名称
func (s *sWdkProject) getWdkProjectSignCompany(signCompany uint) (name string) {
	switch signCompany {
	case 0:
		name = "绿城房地产咨询集团有限公司"
	case 1:
		name = "浙江幸福绿城房地产咨询有限公司"
	case 2:
		name = "浙江美好绿城房地产咨询有限公司"
	default:
		name = ""
	}
	return
}

// createWdkProjectExcel 创建文档库项目信息Excel表
func (s *sWdkProject) createWdkProjectExcel(data [][]any) (fileInfo *v1.WdkProjectExportFile, err error) {
	f := excelize.NewFile()
	sheetName := "项目信息"
	f.SetSheetName("Sheet1", sheetName)
	// 循环写入
	for i, row := range data {
		var startCell string
		if startCell, err = excelize.JoinCellName("A", i+1); err != nil {
			return
		}
		if err = f.SetSheetRow(sheetName, startCell, &row); err != nil {
			return
		}
	}
	// 合并单元格
	if err = f.MergeCell(sheetName, "A1", "T1"); err != nil {
		return
	}
	// 获取Excel文件表头样式
	var hstyle int
	if hstyle, err = utils.GetExcelHeadStyle(f); err != nil {
		return
	}
	// 设置Excel文件表头样式
	if err = f.SetCellStyle(sheetName, "A1", "T2", hstyle); err != nil {
		return
	}
	// 获取Excel文件表体样式
	var bstyle int
	if bstyle, err = utils.GetExcelBodyStyle(f); err != nil {
		return
	}
	// 设置Excel文件表体样式
	endRow := gconv.String(len(data) + 1)
	if err = f.SetCellStyle(sheetName, "A3", "T"+endRow, bstyle); err != nil {
		return
	}
	// 设置列宽
	if err = f.SetColWidth(sheetName, "A", "T", 20); err != nil {
		return
	}
	// 合同金额求和
	maxRow := gconv.String(len(data))
	sumRowAxis := "J" + endRow
	if err = f.SetCellFormula(sheetName, sumRowAxis, "=SUM(J3:J"+maxRow+")"); err != nil {
		return
	}
	// 保存
	fileName := gtime.Datetime() + "项目信息导出表.xlsx"
	filePath := "cache/excel/" + fileName
	if err = f.SaveAs(filePath); err != nil {
		return
	}
	// 返回文件信息
	fileInfo = &v1.WdkProjectExportFile{
		FileName: fileName,
		FilePath: filePath,
	}
	return
}
