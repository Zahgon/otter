package client

import (
	"github.com/karlseguin/ccache/v3"
)

type Ccache[V any] struct {
	client *ccache.Cache[V]
}

func (c *Ccache[V]) Init(capacity int) { _ = "STUB: not implemented"; return }

//nolint:gosec // there will never be an overflow

func (c *Ccache[V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Ccache[V]) Get(key string) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Ccache[V]) Set(key string, value V) { _ = "STUB: not implemented"; return }

func (c *Ccache[V]) Close() { _ = "STUB: not implemented"; return }
