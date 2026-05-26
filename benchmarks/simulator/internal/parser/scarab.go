package parser

import (
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type Scarab struct {
	reader io.Reader
	buffer []byte
}

func NewScarab(reader io.Reader) *Scarab { _ = "STUB: not implemented"; return nil }

func (c *Scarab) Parse(send func(event event.AccessEvent) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
