package service

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	v1 "lczx/api/v1"
	"lczx/utility/utils"
	"os"
	"runtime"
	"strconv"
	"time"
)

type sServerMonitor struct {
	startTime *gtime.Time
}

var (
	insServerMonitor = sServerMonitor{
		startTime: gtime.Now(),
	}
)

// ServerMonitor 服务监控服务
func ServerMonitor() *sServerMonitor {
	return &insServerMonitor
}

// ServerInfo 服务监控信息
func (s *sServerMonitor) ServerInfo() (serverInfo *v1.ServerMonitorInfo) {
	var err error
	// cpu使用率
	var cpuUsed float64 = 0
	var cpuInfo []float64
	cpuInfo, err = cpu.Percent(time.Second, false)
	if err == nil {
		cpuUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuInfo[0]), 64)
	}
	// cpu负载5、当前空闲率
	var cpuAvg5 float64 = 0
	var cpuAvg15 float64 = 0
	var loadInfo *load.AvgStat
	loadInfo, err = load.Avg()
	if err == nil {
		cpuAvg5, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
		cpuAvg15, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
	}
	// 总内存、已用内存、剩余内存、内存使用率
	var memTotal uint64 = 0
	var memUsed uint64 = 0
	var memFree uint64 = 0
	var memUsage float64 = 0
	var vMem *mem.VirtualMemoryStat
	vMem, err = mem.VirtualMemory()
	if err == nil {
		memTotal = vMem.Total
		memUsed = vMem.Used
		memFree = memTotal - memUsed
		memUsage, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", vMem.UsedPercent), 64)
	}
	// go分配的总内存、go使用的内存、go剩余的内存、go内存使用率
	var goTotal uint64 = 0
	var goUsed uint64 = 0
	var goFree uint64 = 0
	var goUsage float64 = 0
	var gomem runtime.MemStats
	runtime.ReadMemStats(&gomem)
	goTotal = gomem.TotalAlloc
	goUsed = gomem.Sys
	goFree = gomem.Frees
	goUsage = gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(goUsed)/gconv.Float64(memTotal)*100))
	// 服务器IP、硬件服务器名称、操作系统、系统架构
	var serverName string
	var osName string
	var osArch string
	serverIp, _ := utils.GetLocalIp()
	var sysInfo *host.InfoStat
	sysInfo, err = host.Info()
	if err == nil {
		serverName = sysInfo.Hostname
		osName = sysInfo.OS
		osArch = sysInfo.KernelArch
	}
	// 项目路径
	serverDir, _ := os.Getwd()
	// 运行时长
	serverRunTime := gtime.Now().Timestamp() - s.startTime.Timestamp()
	// 服务器磁盘信息
	diskList := make([]disk.UsageStat, 0)
	var diskInfo []disk.PartitionStat
	diskInfo, err = disk.Partitions(true) // 所有分区
	if err == nil {
		for _, p := range diskInfo {
			var diskDetail *disk.UsageStat
			diskDetail, err = disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskList = append(diskList, *diskDetail)
			}
		}
	}

	serverInfo = &v1.ServerMonitorInfo{
		CpuNum:          runtime.NumCPU(),  // cpu核心数
		CpuUsed:         cpuUsed,           // cpu使用率
		CpuAvg5:         cpuAvg5,           // cpu负载5
		CpuAvg15:        cpuAvg15,          // 当前空闲率
		MemTotal:        memTotal,          // 总内存
		GoTotal:         goTotal,           // go分配的总内存
		MemUsed:         memUsed,           // 已用内存
		GoUsed:          goUsed,            // go使用的内存
		MemFree:         memFree,           // 剩余内存
		GoFree:          goFree,            // go剩余的内存
		MemUsage:        memUsage,          // 内存使用率
		GoUsage:         goUsage,           // go内存使用率
		ServerName:      serverName,        // 硬件服务器名称
		OsName:          osName,            // 操作系统
		ServerIp:        serverIp,          // 服务器IP
		OsArch:          osArch,            // 系统架构
		ProgramName:     "GoLang",          // 编程语言
		ProgramVersion:  runtime.Version(), // 编程语言版本
		ProgramHome:     runtime.GOROOT(),  // 编程语言安装路径
		ServerDir:       serverDir,         // 项目路径
		ServerStartTime: s.startTime,       // 启动时间
		ServerRunTime:   serverRunTime,     // 运行时长
		DiskList:        diskList,          // 服务器磁盘信息
	}
	return
}
