package service

import (
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
)

type sCrontab struct{}

var (
	insCrontab = sCrontab{}
)

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
