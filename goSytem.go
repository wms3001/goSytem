package goSytem

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shopspring/decimal"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/goTool"
	"net"
	"os"
	"runtime"
	"strings"
)

type GoMemery struct {
	Unit        string `json:"unit"`
	Total       uint64 `json:"total"`
	Free        uint64 `json:"free"`
	Active      uint64 `json:"active"`
	Available   uint64 `json:"available"`
	Used        uint64 `json:"used"`
	UsedPercent string `json:"usedPercent"`
}

type GoCpu struct {
	Name string `json:"name"`
	Core int    `json:"core"`
}

type Addr struct {
	Ip   string `json:"ip"`
	Mask string `json:"mask"`
}

type GoNet struct {
	Addrs []Addr `json:"addrs"`
}

type GoDisk struct {
	Unit        string `json:"unit"`
	Total       uint64 `json:"total"`
	Free        uint64 `json:"free"`
	Used        uint64 `json:"used"`
	UsedPercent string `json:"usedPercent"`
}

type GoSytem struct {
	Hostname string   `json:"hostname"`
	System   string   `json:"system"`
	Platform string   `json:"platform"`
	Memery   GoMemery `json:"memery"`
	Cpu      GoCpu    `json:"cpu"`
	Net      GoNet    `json:"net"`
	Disk     GoDisk   `json:"disk"`
}

func (goSytem *GoSytem) Info() *goCommon.Resp {
	goTool := goTool.GoTool{}
	goSytem.Hostname, _ = os.Hostname()
	goSytem.System = runtime.GOOS
	p, _, _, _ := host.PlatformInformation()
	goSytem.Platform = p
	memery, _ := mem.VirtualMemory()
	goSytem.Memery.Total = memery.Total / uint64(1024*1024)
	goSytem.Memery.Free = memery.Free / uint64(1024*1024)
	goSytem.Memery.Active = memery.Active / uint64(1024*1024)
	goSytem.Memery.Available = memery.Available / uint64(1024*1024)
	goSytem.Memery.Used = memery.Used / uint64(1024*1024)
	usePercent, _ := decimal.NewFromFloat(memery.UsedPercent).Round(2).Float64()
	goSytem.Memery.UsedPercent = goTool.Strval(usePercent) + "%"
	goSytem.Memery.Unit = "m"
	goSytem.Cpu.Core, _ = cpu.Counts(true)
	c, _ := cpu.Info()
	goSytem.Cpu.Name = c[0].ModelName
	addrs, _ := net.InterfaceAddrs()
	var addrList []Addr
	for _, add := range addrs {
		newAddr := Addr{}
		s := strings.Split(add.String(), "/")
		newAddr.Ip = s[0]
		newAddr.Mask = s[1]
		addrList = append(addrList, newAddr)
	}
	goSytem.Net.Addrs = addrList
	d, _ := disk.Usage("/")
	goSytem.Disk.Unit = "G"
	goSytem.Disk.Total = d.Total / uint64(1024*1024*1024)
	goSytem.Disk.Free = d.Free / uint64(1024*1024*1024)
	goSytem.Disk.Used = d.Used / uint64(1024*1024*1024)
	useDiskPercent, _ := decimal.NewFromFloat(d.UsedPercent).Round(2).Float64()
	goSytem.Disk.UsedPercent = goTool.Strval(useDiskPercent) + "%"
	jStr, _ := json.Marshal(goSytem)
	var resp = &goCommon.Resp{}
	resp.Code = 1
	resp.Message = "success"
	resp.Data = string(jStr)
	return resp
}
