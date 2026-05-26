package product

import (
	"github.com/dgryski/go-clockpro"
)

type ClockPro struct {
	client *clockpro.Cache
}

func (c *ClockPro) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *ClockPro) Name() string { _ = "STUB: not implemented"; return "" }

func (c *ClockPro) Get(key uint64) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (c *ClockPro) Set(key, value uint64) { _ = "STUB: not implemented"; return }

func (c *ClockPro) Close() { _ = "STUB: not implemented"; return }
