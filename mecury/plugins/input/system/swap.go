package system

import (
	"fmt"

	"github.com/aiyun/openapm/mecury/agent"
)

type SwapStats struct {
	ps PS
}

func (_ *SwapStats) Description() string {
	return "Read metrics about swap memory usage"
}

func (_ *SwapStats) SampleConfig() string { return "" }

func (s *SwapStats) Gather(acc agent.Accumulator) error {
	swap, err := s.ps.SwapStat()
	if err != nil {
		return fmt.Errorf("error getting swap memory info: %s", err)
	}

	fields := map[string]interface{}{
		"total":        swap.Total,
		"used":         swap.Used,
		"free":         swap.Free,
		"used_percent": swap.UsedPercent,
		"in":           swap.Sin,
		"out":          swap.Sout,
	}
	acc.AddFields("swap", fields, nil)

	return nil
}

func init() {
	agent.AddInput("swap", &SwapStats{ps: &systemPS{}})
}
