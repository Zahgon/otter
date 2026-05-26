package trace

import (
	"io"
)

func wrapDecoder(r io.Reader, path string) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// without decoding

func NewReader(path string) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
