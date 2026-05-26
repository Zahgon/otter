package client

import (
	"github.com/jellydator/ttlcache/v3"
)

type TTLCache[K comparable, V any] struct {
	client *ttlcache.Cache[K, V]
}

func (c *TTLCache[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

// ttlcache.WithTTL[K, V](time.Hour),

func (c *TTLCache[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *TTLCache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *TTLCache[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *TTLCache[K, V]) Close() { _ = "STUB: not implemented"; return }
