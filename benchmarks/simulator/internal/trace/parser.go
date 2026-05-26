package trace

import (
	"errors"
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

var ErrUnknownTraceFormat = errors.New("unknown trace format")

type parserContract interface {
	Parse(send func(event event.AccessEvent) bool) (bool, error)
}

func NewParser(traceType string, reader io.Reader) (parserContract, error) {
	_ = "STUB: not implemented"
	return *new(parserContract), nil
}
