package service

type sServerMonitor struct{}

var (
	insServerMonitor = sServerMonitor{}
)

// ServerMonitor 服务监控服务
func ServerMonitor() *sServerMonitor {
	return &insServerMonitor
}
