package service

import (
	"github.com/gogf/gf/v2/os/gmutex"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
)

type TimeTask struct {
	FuncName string
	Param    []string
	Run      func(ctx context.Context)
}

type sCrontab struct {
	taskList []*TimeTask
	mu       *gmutex.Mutex
}

var (
	insCrontab = sCrontab{}
)

func init() {
	insCrontab.mu = gmutex.New()
	checkUserOnlineTask := &TimeTask{
		FuncName: "checkUserOnline",
		Run:      Auth().CheckUserOnline,
	}
	insCrontab.AddTask(checkUserOnlineTask)
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

// AddTask 添加任务
func (s *sCrontab) AddTask(task *TimeTask) *sCrontab {
	if task.FuncName == "" || task.Run == nil {
		return s
	}
	s.taskList = append(s.taskList, task)
	return s
}

// GetTaskByName 通过方法名获取对应task信息
func (s *sCrontab) GetTaskByName(funcName string) *TimeTask {
	var result *TimeTask
	for _, item := range s.taskList {
		if item.FuncName == funcName {
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
		if item.FuncName == funcName {
			item.Param = params
			break
		}
	}
}
