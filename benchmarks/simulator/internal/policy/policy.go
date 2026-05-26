package policy

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/policy/product"
)

type Policy struct {
	policy product.Policy[uint64, uint64]
	hits   uint64
	misses uint64
}

func NewPolicy(c product.Policy[uint64, uint64]) *Policy { _ = "STUB: not implemented"; return nil }

func (p *Policy) Record(e event.AccessEvent) { _ = "STUB: not implemented"; return }

func (p *Policy) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Policy) Init(capacity int) { _ = "STUB: not implemented"; return }

func (p *Policy) Ratio() float64 { _ = "STUB: not implemented"; return 0 }

func (p *Policy) Close() { _ = "STUB: not implemented"; return }
