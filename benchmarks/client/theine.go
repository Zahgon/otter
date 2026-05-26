package client

import (
	"github.com/Yiling-J/theine-go"
)

type Theine[K comparable, V any] struct {
	client *theine.Cache[K, V]
}

func (c *Theine[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *Theine[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Theine[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Theine[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Theine[K, V]) Close() { _ = "STUB: not implemented"; return }
