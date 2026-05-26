package generator

import (
	"sync"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type genFunc func(sender *sender[event.AccessEvent]) (stop bool)

type base struct {
	once     sync.Once
	stream   Stream[event.AccessEvent]
	generate genFunc
	limit    *uint
}

func newBase(generate genFunc, limit *uint) base { _ = "STUB: not implemented"; return *new(base) }

func (b *base) Generate() Stream[event.AccessEvent] { _ = "STUB: not implemented"; return nil }
