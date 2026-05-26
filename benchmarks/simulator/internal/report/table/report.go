package table

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/report/simulation"
)

func formatInt(n int64) string { _ = "STUB: not implemented"; return "" }

// First character is the - sign (not a digit)

type Table struct {
	table [][]simulation.Result
}

func NewTable(table [][]simulation.Result) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) Report() error { _ = "STUB: not implemented"; return nil }
