package controller

import (
	"context"
	v1 "lczx/api/v1"
	"lczx/internal/service"
)

var (
	ServerMonitor = cServerMonitor{}
)

type cServerMonitor struct{}

// Info 获取服务监控信息
func (c *cServerMonitor) Info(ctx context.Context, req *v1.ServerMonitorInfoReq) (res *v1.ServerMonitorInfoRes, err error) {
	serverInfo := service.ServerMonitor().ServerInfo()
	res = &v1.ServerMonitorInfoRes{Info: serverInfo}
	return
}
