package simulator

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/config"
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/trace/generator"
)

func newGenerator(cfg config.Config) (traceGenerator, error) {
	_ = "STUB: not implemented"
	return *new(traceGenerator), nil
}

func toPaths(paths []config.FilePath) []generator.FilePath { _ = "STUB: not implemented"; return nil }
