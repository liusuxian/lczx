package wdk

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
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sWdkProject struct {
	clientOptionMap map[string][]*model.ClientOption // 客户端选项
}

func init() {
	service.RegisterWdkProject(newWdkProject())
}

// 文档库项目管理服务
func newWdkProject() *sWdkProject {
	insWdkProject := &sWdkProject{}
	insWdkProject.clientOptionMap = map[string][]*model.ClientOption{}
	typeList := []*model.ClientOption{
		{
			Name:  "蓝绿体系",
			Value: "0",
		},
		{
			Name:  "非绿",
			Value: "1",
		},
	}
	originList := []*model.ClientOption{
		{
			Name:  "绿中",
			Value: "0",
		},
		{
			Name:  "分子公司",
			Value: "1",
		},
		{
			Name:  "合伙人",
			Value: "2",
		},
		{
			Name:  "老客户",
			Value: "3",
		},
		{
			Name:  "中交",
			Value: "4",
		},
		{
			Name:  "蓝城",
			Value: "5",
		},
		{
			Name:  "自拓",
			Value: "6",
		},
		{
			Name:  "其他",
			Value: "7",
		},
	}
	stepList := []*model.ClientOption{
		{
			Name:  "未开始",
			Value: "0",
		},
		{
			Name:  "合同签约",
			Value: "1",
		},
		{
			Name:  "项目启动会",
			Value: "2",
		},
		{
			Name:  "服务中-规划设计",
			Value: "3",
		},
		{
			Name:  "服务中-项目展示区施工",
			Value: "4",
		},
		{
			Name:  "服务中-主体结构工程",
			Value: "5",
		},
		{
			Name:  "服务中-主体安装工程",
			Value: "6",
		},
		{
			Name:  "服务中-装饰装修工程",
			Value: "7",
		},
		{
			Name:  "服务中-景观市政工程",
			Value: "8",
		},
		{
			Name:  "服务中-项目交付验收",
			Value: "9",
		},
		{
			Name:  "合同结束",
			Value: "30",
		},
		{
			Name:  "复盘",
			Value: "31",
		},
	}
	uploadStatusList := []*model.ClientOption{
		{
			Name:  "异常",
			Value: "0",
		},
		{
			Name:  "正常",
			Value: "1",
		},
		{
			Name:  "已完成",
			Value: "2",
		},
	}
	businessTypeList := []*model.ClientOption{
		{
			Name:  "物业",
			Value: "0",
		},
		{
			Name:  "专项",
			Value: "1",
		},
		{
			Name:  "全过程",
			Value: "2",
		},
	}
	businessFormsList := []*model.ClientOption{
		{
			Name:  "住宅",
			Value: "0",
		},
		{
			Name:  "小高层",
			Value: "1",
		},
		{
			Name:  "高层",
			Value: "2",
		},
		{
			Name:  "超高层",
			Value: "3",
		},
		{
			Name:  "公寓",
			Value: "4",
		},
		{
			Name:  "合院",
			Value: "5",
		},
		{
			Name:  "叠墅",
			Value: "6",
		},
		{
			Name:  "排屋",
			Value: "7",
		},
		{
			Name:  "多层",
			Value: "8",
		},
		{
			Name:  "会所",
			Value: "9",
		},
		{
			Name:  "商住",
			Value: "10",
		},
		{
			Name:  "综合体",
			Value: "11",
		},
		{
			Name:  "产业园",
			Value: "12",
		},
		{
			Name:  "酒店",
			Value: "13",
		},
		{
			Name:  "酒店式公寓",
			Value: "14",
		},
		{
			Name:  "商业",
			Value: "15",
		},
		{
			Name:  "普通商业",
			Value: "16",
		},
		{
			Name:  "公共配套",
			Value: "17",
		},
		{
			Name:  "办公",
			Value: "18",
		},
		{
			Name:  "公寓式办公",
			Value: "19",
		},
		{
			Name:  "厂房",
			Value: "20",
		},
	}
	contractStatusList := []*model.ClientOption{
		{
			Name:  "新签",
			Value: "0",
		},
		{
			Name:  "续签",
			Value: "1",
		},
		{
			Name:  "未签",
			Value: "2",
		},
	}
	deepCultureList := []*model.ClientOption{
		{
			Name:  "否",
			Value: "0",
		},
		{
			Name:  "是",
			Value: "1",
		},
	}
	statusList := []*model.ClientOption{
		{
			Name:  "服务中",
			Value: "0",
		},
		{
			Name:  "暂停",
			Value: "1",
		},
		{
			Name:  "提前终止",
			Value: "2",
		},
		{
			Name:  "跟踪期",
			Value: "3",
		},
		{
			Name:  "洽谈中",
			Value: "4",
		},
		{
			Name:  "正常结束",
			Value: "5",
		},
	}
	signCompanyList := []*model.ClientOption{
		{
			Name:  "绿城房地产咨询集团有限公司",
			Value: "0",
		},
		{
			Name:  "浙江幸福绿城房地产咨询有限公司",
			Value: "1",
		},
		{
			Name:  "浙江美好绿城房地产咨询有限公司",
			Value: "2",
		},
	}
	insWdkProject.clientOptionMap["typeList"] = typeList
	insWdkProject.clientOptionMap["originList"] = originList
	insWdkProject.clientOptionMap["stepList"] = stepList
	insWdkProject.clientOptionMap["uploadStatusList"] = uploadStatusList
	insWdkProject.clientOptionMap["businessTypeList"] = businessTypeList
	insWdkProject.clientOptionMap["businessFormsList"] = businessFormsList
	insWdkProject.clientOptionMap["contractStatusList"] = contractStatusList
	insWdkProject.clientOptionMap["deepCultureList"] = deepCultureList
	insWdkProject.clientOptionMap["statusList"] = statusList
	insWdkProject.clientOptionMap["signCompanyList"] = signCompanyList

	return insWdkProject
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
	gmodel := dao.WdkProject.Ctx(ctx)
	columns := dao.WdkProject.Columns()
	order := "id DESC"
	if req.Id > 0 {
		gmodel = gmodel.Where(columns.Id, req.Id)
	}
	if req.Name != "" {
		gmodel = gmodel.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Type != "" {
		gmodel = gmodel.Where(columns.Type, req.Type)
	}
	if req.Origin != "" {
		gmodel = gmodel.Where(columns.Origin, req.Origin)
	}
	if req.Step != "" {
		gmodel = gmodel.Where(columns.Step, req.Step)
	}
	if req.FileUploadStatus != "" {
		gmodel = gmodel.Where(columns.FileUploadStatus, req.FileUploadStatus)
	}
	if req.BusinessType != "" {
		gmodel = gmodel.Where(columns.BusinessType, req.BusinessType)
	}
	if !projectIdsMap.IsEmpty() {
		gmodel = gmodel.WhereIn(columns.Id, projectIdsMap.Keys())
	}
	if req.ContractStatus != "" {
		gmodel = gmodel.Where(columns.ContractStatus, req.ContractStatus)
	}
	if req.ContractSum != "" {
		gmodel = gmodel.WhereLike(columns.ContractSum, "%"+req.ContractSum+"%")
	}
	if req.DeepCulture != "" {
		gmodel = gmodel.Where(columns.DeepCulture, req.DeepCulture)
	}
	if req.Status != "" {
		gmodel = gmodel.Where(columns.Status, req.Status)
	}
	if req.EntrustCompany != "" {
		gmodel = gmodel.WhereLike(columns.EntrustCompany, "%"+req.EntrustCompany+"%")
	}
	if req.SignCompany != "" {
		gmodel = gmodel.Where(columns.SignCompany, req.SignCompany)
	}
	if req.PrincipalName != "" {
		gmodel = gmodel.WhereLike(columns.PrincipalName, "%"+req.PrincipalName+"%")
	}
	deptIdsMap := gmap.New()
	if req.DeptId != "" {
		// 获取部门状态为正常的部门列表
		var depts []*entity.Dept
		depts, err = service.Dept().GetStatusEnableDepts(ctx)
		if err != nil {
			return
		}
		deptId := gconv.Uint64(req.DeptId)
		deptIdsMap.Set(deptId, true)
		service.Dept().FindSonIdsByParentId(depts, deptId, deptIdsMap)
	}
	if !deptIdsMap.IsEmpty() {
		gmodel = gmodel.WhereIn(columns.DeptId, deptIdsMap.Keys())
	}
	if req.Region != "" {
		gmodel = gmodel.WhereLike(columns.Region, "%"+req.Region+"%")
	}
	if req.StartTime.String() != "" {
		startTime := req.StartTime.Format("Y-m-d")
		gmodel = gmodel.WhereGTE(columns.StartTime, startTime)
	}
	if req.EndTime.String() != "" {
		endTime := req.EndTime.Format("Y-m-d")
		gmodel = gmodel.WhereLTE(columns.EndTime, endTime)
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	if total, err = gmodel.Count(); err != nil {
		return
	}
	if err = gmodel.Page(req.CurPage, req.PageSize).Order(order).ScanList(&list, "ProjectInfo"); err != nil {
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
		principalUser, terr = service.User().GetUserById(ctx, req.PrincipalUid)
		if terr != nil {
			return terr
		}
		if principalUser == nil {
			return gerror.Newf(`负责人用户ID[%d]不存在`, req.PrincipalUid)
		}
		// 通过部门ID判断部门信息是否存在
		var dept *entity.Dept
		dept, terr = service.Dept().SelectDeptById(ctx, req.DeptId)
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
		principalUser, terr = service.User().GetUserById(ctx, req.PrincipalUid)
		if terr != nil {
			return terr
		}
		if principalUser == nil {
			return gerror.Newf(`负责人用户ID[%d]不存在`, req.PrincipalUid)
		}
		// 通过部门ID判断部门信息是否存在
		var dept *entity.Dept
		dept, terr = service.Dept().SelectDeptById(ctx, req.DeptId)
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
func (s *sWdkProject) ExportWdkProject(ctx context.Context, req *v1.WdkProjectExportReq) (filePath string, err error) {
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
	gmodel := dao.WdkProject.Ctx(ctx)
	columns := dao.WdkProject.Columns()
	order := "id DESC"
	if req.Id > 0 {
		gmodel = gmodel.Where(columns.Id, req.Id)
	}
	if req.Name != "" {
		gmodel = gmodel.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Type != "" {
		gmodel = gmodel.Where(columns.Type, req.Type)
	}
	if req.Origin != "" {
		gmodel = gmodel.Where(columns.Origin, req.Origin)
	}
	if req.Step != "" {
		gmodel = gmodel.Where(columns.Step, req.Step)
	}
	if req.FileUploadStatus != "" {
		gmodel = gmodel.Where(columns.FileUploadStatus, req.FileUploadStatus)
	}
	if req.BusinessType != "" {
		gmodel = gmodel.Where(columns.BusinessType, req.BusinessType)
	}
	if !projectIdsMap.IsEmpty() {
		gmodel = gmodel.WhereIn(columns.Id, projectIdsMap.Keys())
	}
	if req.ContractStatus != "" {
		gmodel = gmodel.Where(columns.ContractStatus, req.ContractStatus)
	}
	if req.ContractSum != "" {
		gmodel = gmodel.WhereLike(columns.ContractSum, "%"+req.ContractSum+"%")
	}
	if req.DeepCulture != "" {
		gmodel = gmodel.Where(columns.DeepCulture, req.DeepCulture)
	}
	if req.Status != "" {
		gmodel = gmodel.Where(columns.Status, req.Status)
	}
	if req.EntrustCompany != "" {
		gmodel = gmodel.WhereLike(columns.EntrustCompany, "%"+req.EntrustCompany+"%")
	}
	if req.SignCompany != "" {
		gmodel = gmodel.Where(columns.SignCompany, req.SignCompany)
	}
	if req.PrincipalName != "" {
		gmodel = gmodel.WhereLike(columns.PrincipalName, "%"+req.PrincipalName+"%")
	}
	deptIdsMap := gmap.New()
	if req.DeptId != "" {
		// 获取部门状态为正常的部门列表
		var depts []*entity.Dept
		if depts, err = service.Dept().GetStatusEnableDepts(ctx); err != nil {
			return
		}
		deptId := gconv.Uint64(req.DeptId)
		deptIdsMap.Set(deptId, true)
		service.Dept().FindSonIdsByParentId(depts, deptId, deptIdsMap)
	}
	if !deptIdsMap.IsEmpty() {
		gmodel = gmodel.WhereIn(columns.DeptId, deptIdsMap.Keys())
	}
	if req.Region != "" {
		gmodel = gmodel.WhereLike(columns.Region, "%"+req.Region+"%")
	}
	if req.StartTime.String() != "" {
		startTime := req.StartTime.Format("Y-m-d")
		gmodel = gmodel.WhereGTE(columns.StartTime, startTime)
	}
	if req.EndTime.String() != "" {
		endTime := req.EndTime.Format("Y-m-d")
		gmodel = gmodel.WhereLTE(columns.EndTime, endTime)
	}
	var total int
	if total, err = gmodel.Count(); err != nil {
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
		if err = gmodel.Page(curPage, pageSize).Order(order).ScanList(&list, "ProjectInfo"); err != nil {
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
				s.getClientOption(v.ProjectInfo.Type, "typeList"),
				s.getClientOption(v.ProjectInfo.Origin, "originList"),
				s.getClientOption(v.ProjectInfo.Step, "stepList"),
				s.getClientOption(v.ProjectInfo.FileUploadStatus, "uploadStatusList"),
				s.getClientOption(v.ProjectInfo.BusinessType, "businessTypeList"),
				s.getWdkProjectBusinessforms(v.Businessforms),
				s.getClientOption(v.ProjectInfo.ContractStatus, "contractStatusList"),
				v.ProjectInfo.ContractSum,
				s.getClientOption(v.ProjectInfo.DeepCulture, "deepCultureList"),
				s.getClientOption(v.ProjectInfo.Status, "statusList"),
				v.ProjectInfo.EntrustCompany,
				s.getClientOption(v.ProjectInfo.SignCompany, "signCompanyList"),
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
	filePath, err = s.createWdkProjectExcel(excelData)
	return
}

// GetClientOptionMap 获取客户端选项Map
func (s *sWdkProject) GetClientOptionMap() map[string][]*model.ClientOption {
	return s.clientOptionMap
}

// IsWdkProjectNameAvailable 文档库项目名称是否可用
func (s *sWdkProject) IsWdkProjectNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.WdkProject.Ctx(ctx).Where(do.WdkProject{Name: name}).Unscoped().Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// SetWdkProjectFileUploadStatusFinish 设置文档库项目文件上传状态为已完成
func (s *sWdkProject) SetWdkProjectFileUploadStatusFinish(ctx context.Context, id uint64) (err error) {
	// 通过项目ID获取文档库项目上传文件类型数量
	var fileCount int
	fileCount, err = service.WdkFile().GetWdkFileCountByProjectId(ctx, id)
	if err != nil {
		return
	}
	// 通过项目ID获取文档库项目上传报告类型数量
	var reportCount int
	reportCount, err = service.WdkReport().GetWdkReportCountByProjectId(ctx, id)
	if err != nil {
		return
	}
	// 获取文档库报告类型配置数量
	var reportCfgCount int
	reportCfgCount, err = service.WdkReportCfg().GetWdkReportCfgCount(ctx)
	if err != nil {
		return
	}
	// 判断项目上传文件和项目上传报告是否已全部传完
	if (fileCount + reportCount) >= reportCfgCount+7 {
		// 已传完
		_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{FileUploadStatus: 2, FileUploadLastTime: gtime.Now()}).
			Where(do.WdkProject{Id: id}).Update()
	} else {
		// 未传完
		_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{FileUploadStatus: 1, FileUploadLastTime: gtime.Now()}).
			Where(do.WdkProject{Id: id}).WhereNot(dao.WdkProject.Columns().FileUploadStatus, 2).Update()
	}
	return
}

// SetWdkProjectFileUploadStatusAbnormal 设置文档库项目文件上传状态为异常
func (s *sWdkProject) SetWdkProjectFileUploadStatusAbnormal(ctx context.Context, id uint64) (err error) {
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{FileUploadStatus: 0}).Where(do.WdkProject{Id: id}).
		WhereNot(dao.WdkProject.Columns().FileUploadStatus, 2).Update()
	return
}

// SetWdkProjectFileUploadStatusNormal 设置文档库项目文件上传状态为正常
func (s *sWdkProject) SetWdkProjectFileUploadStatusNormal(ctx context.Context, id uint64) (err error) {
	_, err = dao.WdkProject.Ctx(ctx).Data(do.WdkProject{FileUploadStatus: 1, FileUploadLastTime: gtime.Now()}).
		Where(do.WdkProject{Id: id}).WhereNot(dao.WdkProject.Columns().FileUploadStatus, 2).Update()
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

// CheckWdkProjectFileUploadStatus 检查项目文件上传状态
func (s *sWdkProject) CheckWdkProjectFileUploadStatus(ctx context.Context) {
	var total int
	var err error
	wpModel := dao.WdkProject.Ctx(ctx)
	if total, err = wpModel.Count(); err != nil {
		logger.Error(ctx, "Get WdkProject Count Error: ", err)
		return
	}
	// 循环读取
	wpColumns := dao.WdkProject.Columns()
	curPage := 1
	pageSize := 50
	for {
		var list []*entity.WdkProject
		if err = wpModel.WhereNot(wpColumns.FileUploadStatus, 2).Page(curPage, pageSize).OrderDesc(wpColumns.Id).
			Scan(&list); err != nil {
			logger.Error(ctx, "Get WdkProject Info Error: ", err)
			return
		}
		if list == nil {
			return
		}
		for _, v := range list {
			lastTime := v.FileUploadLastTime
			now := gtime.Now()
			if lastTime == nil || !(lastTime.Year() == now.Year() && lastTime.Month() == now.Month()) {
				// 设置文档库项目文件上传状态为异常
				if err = s.SetWdkProjectFileUploadStatusAbnormal(ctx, v.Id); err != nil {
					logger.Error(ctx, "Set WdkProject FileUploadStatus Abnormal Error: ", err)
				}
			}
		}
		if curPage*pageSize >= total {
			return
		}
		curPage++
	}
}

// saveWdkProject 保存文档库项目数据
func (s *sWdkProject) saveWdkProject(ctx context.Context, req *v1.WdkProjectAddReq, principalUser *entity.User, dept *entity.Dept) (err error) {
	var user *model.ContextUser
	user, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
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
	var user *model.ContextUser
	user, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
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

// getClientOption 通过 value 获取 name
func (s *sWdkProject) getClientOption(value uint, optionKey string) (name string) {
	valueStr := gconv.String(value)
	options := s.GetClientOptionMap()[optionKey]
	for _, option := range options {
		if option.Value == valueStr {
			name = option.Name
			return
		}
	}
	return
}

// getWdkProjectBusinessforms 获取项目业态中文名称
func (s *sWdkProject) getWdkProjectBusinessforms(businessforms []*entity.WdkProjectBusinessforms) (name string) {
	nameList := make([]string, 0, len(businessforms))
	for _, v := range businessforms {
		option := s.getClientOption(v.BusinessForms, "businessFormsList")
		nameList = append(nameList, option)
	}
	name = gstr.Join(nameList, ", ")
	return
}

// createWdkProjectExcel 创建文档库项目信息Excel表
func (s *sWdkProject) createWdkProjectExcel(data [][]any) (filePath string, err error) {
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
	filePath = "cache/excel/" + fileName
	err = f.SaveAs(filePath)
	return
}
