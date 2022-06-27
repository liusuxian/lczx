package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

var (
	Crontab = cCrontab{}
)

type cCrontab struct{}

// List 定时任务列表
func (c *cCrontab) List(ctx context.Context, req *v1.CrontabListReq) (res *v1.CrontabListRes, err error) {
	var total int
	var list []*entity.Crontab
	total, list, err = service.Crontab().GetCrontabList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetCrontabListFailed, err)
		return
	}

	res = &v1.CrontabListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// Add 新增定时任务
func (c *cCrontab) Add(ctx context.Context, req *v1.CrontabAddReq) (res *v1.CrontabAddRes, err error) {
	err = service.Crontab().AddCrontab(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.AddCrontabFailed, err)
		return
	}

	return
}

// Info 获取定时任务
func (c *cCrontab) Info(ctx context.Context, req *v1.CrontabInfoReq) (res *v1.CrontabInfoRes, err error) {
	var info *entity.Crontab
	info, err = service.Crontab().GetCrontabById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.GetCrontabFailed, err)
		return
	}

	res = &v1.CrontabInfoRes{Info: info}
	return
}

// Edit 编辑定时任务
func (c *cCrontab) Edit(ctx context.Context, req *v1.CrontabEditReq) (res *v1.CrontabEditRes, err error) {
	err = service.Crontab().EditCrontab(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.EditCrontabFailed, err)
		return
	}

	return
}

// Delete 删除定时任务
func (c *cCrontab) Delete(ctx context.Context, req *v1.CrontabDeleteReq) (res *v1.CrontabDeleteRes, err error) {
	err = service.Crontab().DeleteCrontab(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteCrontabFailed, err)
		return
	}

	return
}
