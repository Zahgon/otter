package chart

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/report/simulation"
)

type Chart struct {
	name  string
	table [][]simulation.Result
}

func NewChart(name string, table [][]simulation.Result) *Chart {
	_ = "STUB: not implemented"
	return nil
}

func (c *Chart) Report() error { _ = "STUB: not implemented"; return nil }

// for png render
