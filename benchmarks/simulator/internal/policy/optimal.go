package policy

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type Optimal struct {
	capacity uint64
	hits     map[uint64]uint64
	access   []uint64
}

func NewOptimal(capacity int) *Optimal { _ = "STUB: not implemented"; return nil }

func (o *Optimal) Record(e event.AccessEvent) { _ = "STUB: not implemented"; return }

func (o *Optimal) Ratio() float64 { _ = "STUB: not implemented"; return 0 }

func (o *Optimal) Name() string { _ = "STUB: not implemented"; return "" }

func (o *Optimal) Close() { _ = "STUB: not implemented"; return }

type optimalItem struct {
	key  uint64
	hits uint64
}

type optimalHeap []*optimalItem

func (h optimalHeap) Len() int           { _ = "STUB: not implemented"; return 0 }
func (h optimalHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (h optimalHeap) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *optimalHeap) Push(x any) { _ = "STUB: not implemented"; return }

func (h *optimalHeap) Pop() any { _ = "STUB: not implemented"; return *new(any) }
