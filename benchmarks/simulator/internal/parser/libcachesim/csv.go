package libcachesim

import (
	"bufio"
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

type CSV struct {
	scanner *bufio.Scanner
	i       int
}

func NewCSV(reader io.Reader) *CSV { _ = "STUB: not implemented"; return nil }

func (c *CSV) Parse(send func(event event.AccessEvent) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
