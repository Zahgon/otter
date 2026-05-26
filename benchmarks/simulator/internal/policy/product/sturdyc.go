package product

import (
	"github.com/viccon/sturdyc"
)

type Sturdyc struct {
	client *sturdyc.Client[uint64]
}

func (c *Sturdyc) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *Sturdyc) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Sturdyc) Get(key uint64) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (c *Sturdyc) Set(key uint64, value uint64) { _ = "STUB: not implemented"; return }

func (c *Sturdyc) Close() { _ = "STUB: not implemented"; return }
