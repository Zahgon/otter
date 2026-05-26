package product

import "github.com/maypok86/otter/v2"

type Otter[K comparable, V any] struct {
	client *otter.Cache[K, V]
}

func (c *Otter[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *Otter[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Otter[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Otter[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Otter[K, V]) Close() { _ = "STUB: not implemented"; return }
