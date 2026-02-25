package system

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfo struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	CPU      string `json:"cpu"`
	RAM      string `json:"ram"`
	Host     string `json:"host"`
	Platform string `json:"platform"`
}

func GetSystemInfo() (*SystemInfo, error) {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	h, _ := host.Info()

	cpuModel := "Unknown"
	if len(c) > 0 {
		cpuModel = c[0].ModelName
	}

	return &SystemInfo{
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		CPU:      fmt.Sprintf("%s (%d cores)", cpuModel, len(c)),
		RAM:      fmt.Sprintf("%.2f GB (Total)", float64(v.Total)/1024/1024/1024),
		Host:     h.Hostname,
		Platform: h.Platform,
	}, nil
}

func (s *SystemInfo) String() string {
	return fmt.Sprintf("OS: %s, Arch: %s, CPU: %s, RAM: %s, Host: %s, Platform: %s",
		s.OS, s.Arch, s.CPU, s.RAM, s.Host, s.Platform)
}
