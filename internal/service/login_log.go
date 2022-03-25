package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/user_agent"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/utility/logger"
	"lczx/utility/utils"
)

type sLoginLog struct {
	Pool *grpool.Pool
}

var insLoginLog = sLoginLog{
	Pool: grpool.New(100),
}

// LoginLog 系统登录日志管理服务
func LoginLog() *sLoginLog {
	return &insLoginLog
}

// Invoke 异步保存日志
func (s *sLoginLog) Invoke(req *ghttp.Request, data *entity.LoginLog) {
	ctx := req.GetCtx()
	err := s.Pool.Add(ctx, func(ctx context.Context) {
		// 写入日志数据
		ip := utils.GetClientIp(req)
		data.Ip = ip
		data.Location = utils.GetCityByIp(ctx, ip)
		userAgent := req.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		explorer, _ := ua.Browser()
		data.Browser = explorer
		data.Os = ua.OS()
		data.Time = gtime.Now()
		if data.Module == "" {
			data.Module = "系统后台"
		}
		_, err := dao.LoginLog.Ctx(ctx).Insert(data)
		if err != nil {
			logger.Error(ctx, "SysLoginLog Insert Error: ", err.Error())
		}
	})
	if err != nil {
		logger.Error(ctx, "SysLoginLog Pool Add Error: ", err.Error())
	}
}
