package parser

import (
	"bufio"
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type LIRS struct {
	scanner *bufio.Scanner
}

func NewLIRS(reader io.Reader) *LIRS { _ = "STUB: not implemented"; return nil }

func (l *LIRS) Parse(send func(event event.AccessEvent) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
