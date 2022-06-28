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

// ClientOptions 客户端选项
func (c *cCrontab) ClientOptions(ctx context.Context, req *v1.CrontabClientOptionsReq) (res *v1.CrontabClientOptionsRes, err error) {
	// 获取客户端选项Map
	clientOptionMap := service.Crontab().GetClientOptionMap()
	groupList := make([]*v1.CrontabClientOption, 0, len(clientOptionMap["groupList"]))
	for _, v := range clientOptionMap["groupList"] {
		name, value := service.Crontab().GetClientOption(v)
		groupList = append(groupList, &v1.CrontabClientOption{Name: name, Value: value})
	}
	statusList := make([]*v1.CrontabClientOption, 0, len(clientOptionMap["statusList"]))
	for _, v := range clientOptionMap["statusList"] {
		name, value := service.Crontab().GetClientOption(v)
		statusList = append(statusList, &v1.CrontabClientOption{Name: name, Value: value})
	}
	invokeTargetList := make([]*v1.CrontabClientOption, 0, len(clientOptionMap["invokeTargetList"]))
	for _, v := range clientOptionMap["invokeTargetList"] {
		name, value := service.Crontab().GetClientOption(v)
		invokeTargetList = append(invokeTargetList, &v1.CrontabClientOption{Name: name, Value: value})
	}

	res = &v1.CrontabClientOptionsRes{
		GroupList:        groupList,
		StatusList:       statusList,
		InvokeTargetList: invokeTargetList,
	}
	return
}

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

// Start 启动定时任务
func (c *cCrontab) Start(ctx context.Context, req *v1.CrontabStartReq) (res *v1.CrontabStartRes, err error) {
	var crontab *entity.Crontab
	crontab, err = service.Crontab().GetCrontabById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.StartCrontabFailed, err)
		return
	}
	err = service.Crontab().StartTask(ctx, crontab)
	if err != nil {
		err = gerror.WrapCode(code.StartCrontabFailed, err)
		return
	}

	return
}

// Stop 停止定时任务
func (c *cCrontab) Stop(ctx context.Context, req *v1.CrontabStopReq) (res *v1.CrontabStopRes, err error) {
	var crontab *entity.Crontab
	crontab, err = service.Crontab().GetCrontabById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.StopEditCrontabFailed, err)
		return
	}
	err = service.Crontab().StopTask(ctx, crontab)
	if err != nil {
		err = gerror.WrapCode(code.StopEditCrontabFailed, err)
		return
	}

	return
}

// Run 执行定时任务
func (c *cCrontab) Run(ctx context.Context, req *v1.CrontabRunReq) (res *v1.CrontabRunRes, err error) {
	var crontab *entity.Crontab
	crontab, err = service.Crontab().GetCrontabById(ctx, req.Id)
	if err != nil {
		err = gerror.WrapCode(code.RunCrontabFailed, err)
		return
	}
	err = service.Crontab().RunTask(ctx, crontab)
	if err != nil {
		err = gerror.WrapCode(code.RunCrontabFailed, err)
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
