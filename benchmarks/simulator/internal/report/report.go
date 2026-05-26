package report

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/report/simulation"
)

type reporter interface {
	Report() error
}

type Reporter struct {
	reporters []reporter
}

func NewReporter(name string, t [][]simulation.Result) *Reporter {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reporter) Report() error { _ = "STUB: not implemented"; return nil }
