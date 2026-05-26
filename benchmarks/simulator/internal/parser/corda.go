package parser

import (
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type Corda struct {
	reader io.Reader
	buffer []byte
}

func NewCorda(reader io.Reader) *Corda { _ = "STUB: not implemented"; return nil }

func (c *Corda) Parse(send func(event event.AccessEvent) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
