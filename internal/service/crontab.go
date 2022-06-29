package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
)

// 定时任务信息
type timeTask struct {
	funcDescName string                    // 调用方法描述名（中文）
	invokeTarget string                    // 调用目标字符串（英文）
	param        []string                  // 参数
	run          func(ctx context.Context) // 调用方法
}

// 客户端选项
type clientOption struct {
	name  string // 选项显示名
	value string // 选项值
}

// 定时任务管理信息
type sCrontab struct {
	taskList        []*timeTask                // 任务列表
	clientOptionMap map[string][]*clientOption // 客户端选项
	mu              *gmutex.Mutex
}

var (
	insCrontab = sCrontab{}
)

func init() {
	crontabCtx := gctx.New()
	insCrontab.mu = gmutex.New()
	insCrontab.clientOptionMap = map[string][]*clientOption{}
	// 注册定时任务
	checkUserOnlineTask := &timeTask{
		funcDescName: "检查在线用户",
		invokeTarget: "checkUserOnline",
		run:          Auth().CheckUserOnline,
	}
	checkWdkProjectFileUploadStatus := &timeTask{
		funcDescName: "检查项目文件上传状态",
		invokeTarget: "checkWdkProjectFileUploadStatus",
		run:          WdkProject().CheckWdkProjectFileUploadStatus,
	}
	insCrontab.AddTask(checkUserOnlineTask).AddTask(checkWdkProjectFileUploadStatus)
	// 自动执行状态正常的任务
	var crontabList []*entity.Crontab
	var err error
	if crontabList, err = insCrontab.GetStatusNormalCrontab(crontabCtx); err != nil {
		logger.Error(crontabCtx, "自动执行状态正常的任务失败：", err)
	}
	for _, crontab := range crontabList {
		if err = insCrontab.StartTask(crontabCtx, crontab, false); err != nil {
			logger.Error(crontabCtx, "启动任务失败：", err)
		}
	}
	// 处理客户端选项
	groupList := []*clientOption{
		{
			name:  "默认",
			value: "default",
		},
		{
			name:  "系统",
			value: "system",
		},
	}
	misfirePolicyList := []*clientOption{
		{
			name:  "执行一次",
			value: "0",
		},
		{
			name:  "重复执行",
			value: "1",
		},
	}
	statusList := []*clientOption{
		{
			name:  "暂停",
			value: "0",
		},
		{
			name:  "正常",
			value: "1",
		},
	}
	invokeTargetList := make([]*clientOption, 0, len(insCrontab.taskList))
	for _, v := range insCrontab.taskList {
		invokeTargetList = append(invokeTargetList, &clientOption{
			name:  v.funcDescName,
			value: v.invokeTarget,
		})
	}
	insCrontab.clientOptionMap["groupList"] = groupList
	insCrontab.clientOptionMap["misfirePolicyList"] = misfirePolicyList
	insCrontab.clientOptionMap["statusList"] = statusList
	insCrontab.clientOptionMap["invokeTargetList"] = invokeTargetList
}

// Crontab 定时任务管理服务
func Crontab() *sCrontab {
	return &insCrontab
}

// GetCrontabList 获取定时任务列表
func (s *sCrontab) GetCrontabList(ctx context.Context, req *v1.CrontabListReq) (total int, list []*entity.Crontab, err error) {
	model := dao.Crontab.Ctx(ctx)
	columns := dao.Crontab.Columns()
	order := "id ASC"
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Group != "" {
		model = model.Where(columns.Group, req.Group)
	}
	if req.Status != "" {
		model = model.Where(columns.Status, req.Status)
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	if total, err = model.Count(); err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	return
}

// AddCrontab 添加任务
func (s *sCrontab) AddCrontab(ctx context.Context, req *v1.CrontabAddReq) (err error) {
	// 检查任务名称是否可用
	var available bool
	if available, err = s.IsNameAvailable(ctx, req.Name); err != nil {
		return
	}
	if !available {
		return gerror.Newf(`任务名称[%s]已存在`, req.Name)
	}
	// 检查调用方法是否可用
	if available, err = s.IsInvokeTargetAvailable(ctx, req.InvokeTarget); err != nil {
		return
	}
	if !available {
		return gerror.Newf(`调用方法[%s]已存在`, req.InvokeTarget)
	}
	// 新增数据
	user := Context().Get(ctx).User
	var id int64
	id, err = dao.Crontab.Ctx(ctx).Data(do.Crontab{
		Name:           req.Name,
		Group:          req.Group,
		Params:         req.Params,
		InvokeTarget:   req.InvokeTarget,
		CronExpression: req.CronExpression,
		MisfirePolicy:  req.MisfirePolicy,
		Status:         req.Status,
		CreateBy:       user.Id,
		Remark:         req.Remark,
	}).FieldsEx(dao.Crontab.Columns().Id).InsertAndGetId()
	if err != nil {
		return
	}
	// 正常状态启动任务
	if req.Status == 1 {
		crontab := &entity.Crontab{
			Id:             gconv.Uint64(id),
			Name:           req.Name,
			Group:          req.Group,
			Params:         req.Params,
			InvokeTarget:   req.InvokeTarget,
			CronExpression: req.CronExpression,
			MisfirePolicy:  req.MisfirePolicy,
			Status:         req.Status,
		}
		err = s.StartTask(ctx, crontab, false)
	}
	return
}

// GetCrontabById 通过任务ID获取任务信息
func (s *sCrontab) GetCrontabById(ctx context.Context, id uint64) (crontab *entity.Crontab, err error) {
	err = dao.Crontab.Ctx(ctx).Where(do.Crontab{Id: id}).Scan(&crontab)
	return
}

// EditCrontab 修改任务
func (s *sCrontab) EditCrontab(ctx context.Context, req *v1.CrontabEditReq) (err error) {
	// 检查任务信息是否存在
	var crontab *entity.Crontab
	if crontab, err = s.GetCrontabById(ctx, req.Id); err != nil {
		return
	}
	if crontab == nil {
		return gerror.Newf(`任务ID[%d]不存在`, req.Id)
	}
	// 检查任务名称是否可用
	var available bool
	if crontab.Name != req.Name {
		if available, err = s.IsNameAvailable(ctx, req.Name); err != nil {
			return
		}
		if !available {
			return gerror.Newf(`任务名称[%s]已存在`, req.Name)
		}
	}
	// 检查调用方法是否可用
	if crontab.InvokeTarget != req.InvokeTarget {
		if available, err = s.IsInvokeTargetAvailable(ctx, req.InvokeTarget); err != nil {
			return
		}
		if !available {
			return gerror.Newf(`调用方法[%s]已存在`, req.InvokeTarget)
		}
	}
	// 更新数据
	user := Context().Get(ctx).User
	_, err = dao.Crontab.Ctx(ctx).Data(do.Crontab{
		Name:           req.Name,
		Group:          req.Group,
		Params:         req.Params,
		InvokeTarget:   req.InvokeTarget,
		CronExpression: req.CronExpression,
		MisfirePolicy:  req.MisfirePolicy,
		Status:         req.Status,
		UpdateBy:       user.Id,
		Remark:         req.Remark,
	}).Where(do.Crontab{Id: req.Id}).Update()
	if err != nil {
		return
	}
	// 正常状态启动任务
	newCrontab := &entity.Crontab{
		Id:             req.Id,
		Name:           req.Name,
		Group:          req.Group,
		Params:         req.Params,
		InvokeTarget:   req.InvokeTarget,
		CronExpression: req.CronExpression,
		MisfirePolicy:  req.MisfirePolicy,
		Status:         req.Status,
	}
	if req.Status == 1 {
		if err = s.StopTask(ctx, crontab, false); err != nil {
			return
		}
		err = s.StartTask(ctx, newCrontab, false)
	} else {
		err = s.StopTask(ctx, crontab, false)
	}
	return
}

// DeleteCrontab 删除任务
func (s *sCrontab) DeleteCrontab(ctx context.Context, ids []uint64) (err error) {
	idSet := gset.NewFrom(ids)
	var crontabList []*entity.Crontab
	if crontabList, err = insCrontab.GetStatusNormalCrontab(ctx); err != nil {
		return
	}
	for _, crontab := range crontabList {
		if idSet.Contains(crontab.Id) {
			err = gerror.Newf("状态为正常的任务[%s]不能删除", crontab.Name)
			return
		}
	}
	_, err = dao.Crontab.Ctx(ctx).WhereIn(dao.Crontab.Columns().Id, idSet.Slice()).Delete()
	return
}

// GetStatusNormalCrontab 获取状态为正常的任务
func (s *sCrontab) GetStatusNormalCrontab(ctx context.Context) (list []*entity.Crontab, err error) {
	err = dao.Crontab.Ctx(ctx).Where(do.Crontab{Status: 1}).Scan(&list)
	return
}

// SetCrontabStatusNormal 设置任务状态为正常
func (s *sCrontab) SetCrontabStatusNormal(ctx context.Context, id uint64) (err error) {
	_, err = dao.Crontab.Ctx(ctx).Data(do.Crontab{Status: 1}).Where(do.Crontab{Id: id}).Unscoped().Update()
	return
}

// SetCrontabStatusPause 设置任务状态为暂停
func (s *sCrontab) SetCrontabStatusPause(ctx context.Context, id uint64) (err error) {
	_, err = dao.Crontab.Ctx(ctx).Data(do.Crontab{Status: 0}).Where(do.Crontab{Id: id}).Unscoped().Update()
	return
}

// IsNameAvailable 任务名称是否可用
func (s *sCrontab) IsNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.Crontab.Ctx(ctx).Where(do.Crontab{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsInvokeTargetAvailable 调用方法是否可用
func (s *sCrontab) IsInvokeTargetAvailable(ctx context.Context, invokeTarget string) (bool, error) {
	count, err := dao.Crontab.Ctx(ctx).Where(do.Crontab{InvokeTarget: invokeTarget}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetClientOptionMap 获取客户端选项Map
func (s *sCrontab) GetClientOptionMap() map[string][]*clientOption {
	return s.clientOptionMap
}

// GetClientOption 获取客户端选项
func (s *sCrontab) GetClientOption(op *clientOption) (name, value string) {
	return op.name, op.value
}

// AddTask 添加任务
func (s *sCrontab) AddTask(task *timeTask) *sCrontab {
	if task.funcDescName == "" || task.invokeTarget == "" || task.run == nil {
		return s
	}
	s.taskList = append(s.taskList, task)
	return s
}

// GetTaskByInvokeTarget 通过调用目标字符串获取对应task信息
func (s *sCrontab) GetTaskByInvokeTarget(invokeTarget string) *timeTask {
	var result *timeTask
	for _, item := range s.taskList {
		if item.invokeTarget == invokeTarget {
			result = item
			break
		}
	}
	return result
}

// EditParams 修改参数
func (s *sCrontab) EditParams(invokeTarget string, params []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range s.taskList {
		if item.invokeTarget == invokeTarget {
			item.param = params
			break
		}
	}
}

// StartTask 启动任务
func (s *sCrontab) StartTask(ctx context.Context, crontab *entity.Crontab, upStatus bool) (err error) {
	// 获取调用目标是否注册对应的方法
	var task *timeTask
	task = s.GetTaskByInvokeTarget(crontab.InvokeTarget)
	if task == nil {
		return gerror.Newf("调用目标[%s]没有注册对应的方法", crontab.InvokeTarget)
	}
	// 传参
	params := gstr.Split(crontab.Params, "|")
	s.EditParams(task.invokeTarget, params)
	// 搜索任务
	cron := gcron.Search(crontab.InvokeTarget)
	if cron == nil {
		var newCron *gcron.Entry
		if crontab.MisfirePolicy == 1 {
			newCron, err = gcron.AddSingleton(ctx, crontab.CronExpression, task.run, crontab.InvokeTarget)
		} else {
			newCron, err = gcron.AddOnce(ctx, crontab.CronExpression, task.run, crontab.InvokeTarget)
		}
		if err != nil {
			return
		}
		if newCron == nil {
			return gerror.Newf("启动任务[%s]失败", crontab.Name)
		}
	}
	gcron.Start(crontab.InvokeTarget)
	if upStatus {
		err = s.SetCrontabStatusNormal(ctx, crontab.Id)
	}
	return
}

// StopTask 停止任务
func (s *sCrontab) StopTask(ctx context.Context, crontab *entity.Crontab, upStatus bool) (err error) {
	// 获取调用目标是否注册对应的方法
	var task *timeTask
	task = s.GetTaskByInvokeTarget(crontab.InvokeTarget)
	if task == nil {
		return gerror.Newf("调用目标[%s]没有注册对应的方法", crontab.InvokeTarget)
	}
	// 搜索任务
	cron := gcron.Search(crontab.InvokeTarget)
	if cron != nil {
		gcron.Remove(crontab.InvokeTarget)
	}
	if upStatus {
		err = s.SetCrontabStatusPause(ctx, crontab.Id)
	}
	return
}

// RunTask 执行任务
func (s *sCrontab) RunTask(ctx context.Context, crontab *entity.Crontab) (err error) {
	// 获取调用目标是否注册对应的方法
	var task *timeTask
	task = s.GetTaskByInvokeTarget(crontab.InvokeTarget)
	if task == nil {
		return gerror.Newf("调用目标[%s]没有注册对应的方法", crontab.InvokeTarget)
	}
	// 传参
	params := gstr.Split(crontab.Params, "|")
	s.EditParams(task.invokeTarget, params)
	var newCron *gcron.Entry
	if newCron, err = gcron.AddOnce(ctx, "@every 1s", task.run); err != nil {
		return
	}
	if newCron == nil {
		return gerror.Newf("启动任务[%s]失败", crontab.Name)
	}
	return
}
