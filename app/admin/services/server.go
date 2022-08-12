package services

import (
	"fmt"
	"runtime"
	"seed-admin/app/admin/response"
	"seed-admin/common"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type ServerService struct{}

func (*ServerService) GetSystemInfo() (*response.ServerInfo, error) {
	cpu, err := GetCPUInfo()
	if err != nil {
		return nil, err
	}
	host, err := GetHostInfo()
	if err != nil {
		return nil, err
	}
	ram, err := GetRAMInfo()
	if err != nil {
		return nil, err
	}
	disk, err := GetDiskInfo()
	if err != nil {
		return nil, err
	}
	runtime, err := GetRuntimeInfo()
	if err != nil {
		return nil, err
	}
	res := &response.ServerInfo{
		CPU:     *cpu,
		RAM:     *ram,
		Disk:    *disk,
		Host:    *host,
		Runtime: *runtime,
	}
	return res, nil
}

// 获取CPU信息
func GetCPUInfo() (*response.CPU, error) {
	info, err := cpu.Info()
	if err != nil {
		common.LOG.Error("获取CPU信息出错:" + err.Error())
		return nil, err
	}
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		common.LOG.Error("获取CPU百分比出错:" + err.Error())
		return nil, err
	}
	res := &response.CPU{
		Name:    info[0].ModelName,
		Cores:   info[0].Cores,
		Percent: int(cpuPercent[0]),
	}
	return res, nil
}

// 获取主机信息
func GetHostInfo() (*response.Host, error) {
	info, err := host.Info()
	if err != nil {
		common.LOG.Error("获取内存信息出错:" + err.Error())
		return nil, err
	}
	res := &response.Host{
		OS:      info.OS,
		Kernel:  info.KernelArch,
		Runtime: resolveTime(info.Uptime),
	}
	return res, nil
}

// 获取内存信息
func GetRAMInfo() (*response.RAM, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		common.LOG.Error("获取内存信息出错:" + err.Error())
		return nil, err
	}
	res := &response.RAM{
		Total:     BtoGb(info.Total),
		Available: BtoGb(info.Available),
		Used:      BtoGb(info.Used),
		Percent:   int(info.UsedPercent),
	}
	return res, nil
}

// 获取硬盘信息
func GetDiskInfo() (*response.Disk, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		common.LOG.Error("获取硬盘盘符出错:" + err.Error())
		return nil, err
	}
	var total uint64
	var free uint64
	var used uint64
	var percent float64
	for _, part := range parts {
		info, err := disk.Usage(part.Mountpoint)
		if err != nil {
			common.LOG.Error("获取硬盘信息出错:" + err.Error())
			return nil, err
		}
		total += info.Total
		free += info.Free
		used += info.Used
		percent += info.UsedPercent
	}
	res := &response.Disk{
		Total:     BtoGb(total),
		Available: BtoGb(free),
		Used:      BtoGb(used),
		Percent:   int(percent),
	}
	return res, nil
}

// 获取运行信息
func GetRuntimeInfo() (*response.Runtime, error) {
	res := &response.Runtime{
		Version:   runtime.Version(),
		Language:  "Go",
		StartTime: common.StartTime.Format("2006-01-02 15:04:05"),
		Runtime:   resolveTime(uint64(time.Since(common.StartTime).Seconds())),
	}
	return res, nil
}

// B转Gb
func BtoGb(b uint64) float64 {
	str := fmt.Sprintf("%.2f", float64(b)/1024/1024/1024)
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		common.LOG.Error("B转GB出错:" + err.Error())
		return 0
	}
	return res
}

// 秒转换时间
func resolveTime(seconds uint64) string {
	days := seconds / (24 * 3600)
	hours := seconds % (24 * 3600) / 3600
	minutes := seconds % 3600 / 60
	return fmt.Sprintf("%v天%v小时%v分钟", days, hours, minutes)
}
