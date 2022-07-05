// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/internal/model"
	"lczx/internal/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IOperLog interface {
	Invoke(req *ghttp.Request)
	SaveOperLog(ctx context.Context, data *entity.OperLog)
	GetOperLogList(ctx context.Context, req *v1.OperLogListReq) (total int, list []*entity.OperLog, err error)
	DeleteOperLogByIds(ctx context.Context, ids []uint64) (err error)
	ClearOperLog(ctx context.Context) (err error)
	GetClientOptionMap() map[string][]*model.ClientOption
}

var localOperLog IOperLog

func OperLog() IOperLog {
	if localOperLog == nil {
		panic("implement not found for interface IOperLog, forgot register?")
	}
	return localOperLog
}

func RegisterOperLog(i IOperLog) {
	localOperLog = i
}