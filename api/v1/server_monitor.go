package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shirou/gopsutil/disk"
)

// ServerMonitorInfoReq 服务监控信息请求参数
type ServerMonitorInfoReq struct {
	g.Meta `path:"/info" tags:"ServerMonitorInfo" method:"get" summary:"You first monitor/server_monitor/info api"`
}

// ServerMonitorInfoRes 服务监控信息返回参数
type ServerMonitorInfoRes struct {
	Info *ServerMonitorInfo `json:"info" dc:"服务监控信息"` // 服务监控信息
}

type ServerMonitorInfo struct {
	CpuNum          int              `json:"cpuNum" dc:"cpu核心数"`         // cpu核心数
	CpuUsed         float64          `json:"cpuUsed" dc:"cpu使用率"`        // cpu使用率
	CpuAvg5         float64          `json:"cpuAvg5" dc:"cpu负载5"`        // cpu负载5
	CpuAvg15        float64          `json:"cpuAvg15" dc:"当前空闲率"`        // 当前空闲率
	MemTotal        uint64           `json:"memTotal" dc:"总内存"`          // 总内存
	GoTotal         uint64           `json:"goTotal" dc:"go分配的总内存"`      // go分配的总内存
	MemUsed         uint64           `json:"memUsed" dc:"已用内存"`          // 已用内存
	GoUsed          uint64           `json:"goUsed" dc:"go使用的内存"`        // go使用的内存
	MemFree         uint64           `json:"memFree" dc:"剩余内存"`          // 剩余内存
	GoFree          uint64           `json:"goFree" dc:"go剩余的内存"`        // go剩余的内存
	MemUsage        float64          `json:"memUsage" dc:"内存使用率"`        // 内存使用率
	GoUsage         float64          `json:"goUsage" dc:"go内存使用率"`       // go内存使用率
	ServerName      string           `json:"serverName" dc:"硬件服务器名称"`    // 硬件服务器名称
	OsName          string           `json:"osName" dc:"操作系统"`           // 操作系统
	ServerIp        string           `json:"serverIp" dc:"服务器IP"`        // 服务器IP
	OsArch          string           `json:"osArch" dc:"系统架构"`           // 系统架构
	ProgramName     string           `json:"programName" dc:"编程语言"`      // 编程语言
	ProgramVersion  string           `json:"programVersion" dc:"编程语言版本"` // 编程语言版本
	ProgramHome     string           `json:"programHome" dc:"编程语言安装路径"`  // 编程语言安装路径
	ServerDir       string           `json:"serverDir" dc:"项目路径"`        // 项目路径
	ServerStartTime *gtime.Time      `json:"serverStartTime" dc:"启动时间"`  // 启动时间
	ServerRunTime   int64            `json:"serverRunTime" dc:"运行时长"`    // 运行时长
	DiskList        []disk.UsageStat `json:"diskList" dc:"服务器磁盘信息"`      // 服务器磁盘信息
}
