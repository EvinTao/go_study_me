package utils

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// cpu info
func GetCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

func formatF(f1 float64) string {
	return fmt.Sprintf("%f.2", f1)
}

func formatMB(by uint64) string {
	return fmt.Sprintf("%d MB", by/1024/1024)
}

func GetCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("CpuLoad: %s %s %s\n", formatF(info.Load1), formatF(info.Load5), formatF(info.Load15))
}

// mem info
func GetMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%s free:%d\n", formatF(memInfo.UsedPercent), memInfo.Free)
}

// host info
func GetHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// GetDiskInfo disk info
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

func GetNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		if v.Name == "en0" {
			fmt.Printf("net :%v send:%v recv:%v\n", index, formatMB(v.BytesSent), formatMB(v.BytesRecv))
			break
		}
	}
	time.Sleep(time.Second)
}
