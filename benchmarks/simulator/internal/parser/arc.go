package parser

import (
	"bufio"
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type ARC struct {
	scanner *bufio.Scanner
}

func NewARC(reader io.Reader) *ARC { _ = "STUB: not implemented"; return nil }

func (a *ARC) Parse(send func(event event.AccessEvent) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
