package product

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type LRU[K comparable, V any] struct {
	client *lru.Cache[K, V]
}

func (c *LRU[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *LRU[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *LRU[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *LRU[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *LRU[K, V]) Close() { _ = "STUB: not implemented"; return }
