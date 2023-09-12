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

var lastBytesSent uint64
var lastBytesRecv uint64
var lastTime time.Time

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
	ret := fmt.Sprintf("%.2f", f1/1024/1024)
	return ret
}

func formatMem(f1 uint64) string {
	if f1 >= 1014*1024*1024 {
		return fmt.Sprintf("%.1fG", float64(f1)/1024/1024/1024)
	} else if f1 > 1014*1024 && f1 < 1014*1024*1024 {
		return fmt.Sprintf("%dM", f1/1024/1024)
	} else {
		return fmt.Sprintf("%dK", f1/1024)
	}
}

func GetCpuLoad() string {
	info, _ := load.Avg()
	ret := fmt.Sprintf("%.2f %.2f %.2f", info.Load1, info.Load5, info.Load15)
	return ret
}

// mem info
func GetMemInfo() string {
	memInfo, _ := mem.VirtualMemory()
	return fmt.Sprintf("%s %s %s %s", formatMem(memInfo.Used), formatMem(memInfo.Buffers), formatMem(memInfo.Cached), formatMem(memInfo.Free))
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

func GetNetInfo() string {
	info, _ := net.IOCounters(true)
	for _, v := range info {
		if v.Name == "eth0" {
			bytesSent := v.BytesSent
			bytesRecv := v.BytesRecv
			if lastBytesSent != 0 && lastBytesRecv != 0 {
				second := uint64(time.Now().Sub(lastTime).Seconds())
				sentSpeed := (bytesSent - lastBytesSent) / 1024 / second
				recvSpeed := (bytesRecv - lastBytesRecv) / 1024 / second
				return fmt.Sprintf("%dk %dk", sentSpeed, recvSpeed)
			}
			lastBytesSent = bytesSent
			lastBytesRecv = bytesRecv
			lastTime = time.Now()
			break
		}
	}
	return "0k 0k"
}
