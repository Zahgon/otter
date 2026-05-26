package parser

import (
	"errors"
)

var ErrInvalidFormat = errors.New("invalid trace format")

func WrapError(err error) error { _ = "STUB: not implemented"; return nil }
