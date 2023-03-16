package stat

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type InfoStat struct {
	uptime         bool
	bootTime       bool
	procs          bool
	Free           bool
	usedMemory     bool
	usedPercentage bool
}

func (*InfoStat) SampleConfig() string {
	return "Data about host"
}

func (s *InfoStat) Gather(acc telegraf.Accumulator) error {
	host_data, _ := host.Info()
	memory_data, _ := mem.VirtualMemory()

	fields := make(map[string]interface{})

	if s.uptime {
		fields["up_time"] = host_data.Uptime
	}
	if s.procs {
		fields["pros_nums"] = host_data.Procs
	}
	if s.bootTime {
		fields["boot_time"] = host_data.BootTime
	}
	if s.Free {
		fields["free_memory"] = memory_data.Free
	}
	if s.usedMemory {
		fields["total_memory"] = memory_data.Total
	}
	if s.usedPercentage {
		fields["used_percentage"] = memory_data.UsedPercent
	}

	tags := make(map[string]string)

	acc.AddFields("processes_data", fields, tags)

	return nil
}

func init() {
	inputs.Add("processes_mes", func() telegraf.Input {
		return &InfoStat{
			uptime:         true,
			bootTime:       true,
			procs:          true,
			Free:           true,
			usedMemory:     true,
			usedPercentage: true,
		}
	})
}

