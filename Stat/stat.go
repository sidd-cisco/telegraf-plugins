package stat

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/shirou/gopsutil/v3/host"
)

// type InfoStat struct {
// 	Uptime               uint64 `json:"uptime"`
// 	BootTime             uint64 `json:"bootTime"`
// 	Procs                uint64 `json:"procs"`           // number of processes
// 	OS                   string `json:"os"`              // ex: freebsd, linux
// 	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
// 	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
// 	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
// 	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
// 	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
// 	VirtualizationSystem string `json:"virtualizationSystem"`
// 	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
// 	HostID               string `json:"hostid"`             // ex: uuid
// }

type InfoStat struct {
	uptime   bool
	bootTime bool
	procs    bool
}

func (*InfoStat) SampleConfig() string {
	return "Data about host"
}

func (s *InfoStat) Gather(acc telegraf.Accumulator) error {
	data, _ := host.Info()

	fields := make(map[string]interface{})
	if s.uptime {
		fields["up_time"] = data.Uptime
	}
	if s.procs {
		fields["pros_nums"] = data.Procs
	}
	if s.bootTime {
		fields["boot_time"] = data.BootTime
	}

	tags := make(map[string]string)

	acc.AddFields("processes_data", fields, tags)

	return nil
}

func init() {
	inputs.Add("processes_mes", func() telegraf.Input { return &InfoStat{procs: true, bootTime: true, uptime: true} })
}
