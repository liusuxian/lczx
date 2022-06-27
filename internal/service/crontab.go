package service

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gmutex"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/utility/logger"
)

// 定时任务信息
type timeTask struct {
	name     string
	funcName string
	param    []string
	run      func(ctx context.Context)
}

// 定时任务管理信息
type sCrontab struct {
	taskList []*timeTask
	mu       *gmutex.Mutex
}

var (
	insCrontab = sCrontab{}
)

func init() {
	crontabCtx := gctx.New()
	insCrontab.mu = gmutex.New()
	// 注册定时任务
	checkUserOnlineTask := &timeTask{
		name:     "检查在线用户",
		funcName: "checkUserOnline",
		run:      Auth().CheckUserOnline,
	}
	checkWdkProjectFileUploadStatus := &timeTask{
		name:     "检查项目文件上传状态",
		funcName: "checkWdkProjectFileUploadStatus",
		run:      WdkProject().CheckWdkProjectFileUploadStatus,
	}
	insCrontab.AddTask(checkUserOnlineTask).AddTask(checkWdkProjectFileUploadStatus)
	// 自动执行状态正常的任务
	var crontabList []*entity.Crontab
	var err error
	crontabList, err = insCrontab.GetStatusNormalCrontab(crontabCtx)
	if err != nil {
		logger.Error(crontabCtx, "自动执行状态正常的任务失败：", err)
	}
	for _, crontab := range crontabList {
		err = insCrontab.StartTask(crontab)
		if err != nil {
			logger.Error(crontabCtx, "启动任务失败：", err)
		}
	}
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
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	return
}

// AddCrontab 添加任务
func (s *sCrontab) AddCrontab(ctx context.Context, req *v1.CrontabAddReq) (err error) {
	user := Context().Get(ctx).User
	_, err = dao.Crontab.Ctx(ctx).Data(do.Crontab{
		Name:           req.Name,
		Group:          req.Group,
		Params:         req.Params,
		InvokeTarget:   req.InvokeTarget,
		CronExpression: req.CronExpression,
		MisfirePolicy:  req.MisfirePolicy,
		Concurrent:     req.Concurrent,
		Status:         req.Status,
		CreateBy:       user.Id,
		Remark:         req.Remark,
	}).FieldsEx(dao.Crontab.Columns().Id).Insert()
	return
}

// GetCrontabById 通过任务ID获取任务信息
func (s *sCrontab) GetCrontabById(ctx context.Context, id uint64) (crontab *entity.Crontab, err error) {
	err = dao.Crontab.Ctx(ctx).Where(do.Crontab{Id: id}).Scan(&crontab)
	return
}

// EditCrontab 修改任务
func (s *sCrontab) EditCrontab(ctx context.Context, req *v1.CrontabEditReq) (err error) {
	user := Context().Get(ctx).User
	_, err = dao.Crontab.Ctx(ctx).Data(do.Crontab{
		Name:           req.Name,
		Group:          req.Group,
		Params:         req.Params,
		InvokeTarget:   req.InvokeTarget,
		CronExpression: req.CronExpression,
		MisfirePolicy:  req.MisfirePolicy,
		Concurrent:     req.Concurrent,
		Status:         req.Status,
		UpdateBy:       user.Id,
		Remark:         req.Remark,
	}).Where(do.Crontab{Id: req.Id}).Update()
	return
}

// GetStatusNormalCrontab 获取状态正常的任务
func (s *sCrontab) GetStatusNormalCrontab(ctx context.Context) (list []*entity.Crontab, err error) {
	err = dao.Crontab.Ctx(ctx).Where(do.Crontab{Status: 1}).Scan(&list)
	return
}

// GetRegisteredTask 获取已注册的任务
func (s *sCrontab) GetRegisteredTask() (list []*timeTask) {
	return s.taskList
}

// AddTask 添加任务
func (s *sCrontab) AddTask(task *timeTask) *sCrontab {
	if task.name == "" || task.funcName == "" || task.run == nil {
		return s
	}
	s.taskList = append(s.taskList, task)
	return s
}

// GetTaskNameAndFuncName 获取任务名称和任务方法名称
func (s *sCrontab) GetTaskNameAndFuncName(timeTask *timeTask) (name, funcName string) {
	return timeTask.name, timeTask.funcName
}

// GetTaskByName 通过方法名获取对应task信息
func (s *sCrontab) GetTaskByName(funcName string) *timeTask {
	var result *timeTask
	for _, item := range s.taskList {
		if item.funcName == funcName {
			result = item
			break
		}
	}
	return result
}

// EditParams 修改参数
func (s *sCrontab) EditParams(funcName string, params []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range s.taskList {
		if item.funcName == funcName {
			item.param = params
			break
		}
	}
}

// StartTask 启动任务
func (s *sCrontab) StartTask(crontab *entity.Crontab) (err error) {
	var task *timeTask
	task = s.GetTaskByName(crontab.InvokeTarget)
	if task == nil {
		return gerror.Newf("没有绑定对应的方法", crontab.InvokeTarget)
	}

	return
}
