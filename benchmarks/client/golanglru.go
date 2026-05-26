package client

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type GolangLRU[K comparable, V any] struct {
	client *lru.Cache[K, V]
}

func (c *GolangLRU[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *GolangLRU[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *GolangLRU[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *GolangLRU[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *GolangLRU[K, V]) Close() { _ = "STUB: not implemented"; return }
