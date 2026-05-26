package product

import "github.com/hashicorp/golang-lru/arc/v2"

type ARC[K comparable, V any] struct {
	client *arc.ARCCache[K, V]
}

func (c *ARC[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *ARC[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *ARC[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *ARC[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *ARC[K, V]) Close() { _ = "STUB: not implemented"; return }
