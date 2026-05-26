package client

import (
	"github.com/dgraph-io/ristretto"
)

type Ristretto[K ristretto.Key, V any] struct {
	client *ristretto.Cache[K, V]
}

func (c *Ristretto[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *Ristretto[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Ristretto[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Ristretto[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Ristretto[K, V]) Close() { _ = "STUB: not implemented"; return }
