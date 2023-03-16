package Cord

import (
	_ "embed"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Cord struct {
	x float64
	y float64
}

func (*Cord) SampleConfig() string {
	return "Cord"
}

func (s *Cord) Gather(acc telegraf.Accumulator) error {

	fields := make(map[string]interface{})
	fields["x"] = s.x
	fields["y"] = s.y

	tags := make(map[string]string)

	s.x += 1.0
	s.y += 2.0
	acc.AddFields("cord", fields, tags)

	return nil
}

func init() {
	inputs.Add("cord", func() telegraf.Input { return &Cord{x: 0.0} })
}
