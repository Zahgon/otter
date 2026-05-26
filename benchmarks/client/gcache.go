package client

import (
	"github.com/bluele/gcache"
)

type Gcache[K comparable, V any] struct {
	client gcache.Cache
}

func (c *Gcache[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *Gcache[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Gcache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Gcache[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Gcache[K, V]) Close() { _ = "STUB: not implemented"; return }
