package client

import (
	"github.com/viccon/sturdyc"
)

type Sturdyc[V any] struct {
	client *sturdyc.Client[V]
}

func (c *Sturdyc[V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *Sturdyc[V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Sturdyc[V]) Get(key string) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Sturdyc[V]) Set(key string, value V) { _ = "STUB: not implemented"; return }

func (c *Sturdyc[V]) Close() { _ = "STUB: not implemented"; return }
