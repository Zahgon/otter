package product

import fifo "github.com/scalalang2/golang-fifo/v2"

type S3FIFO[K comparable, V any] struct {
	client fifo.Cache[K, V]
}

func (c *S3FIFO[K, V]) Init(capacity int) { _ = "STUB: not implemented"; return }

func (c *S3FIFO[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *S3FIFO[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (c *S3FIFO[K, V]) Name() string { _ = "STUB: not implemented"; return "" }

func (c *S3FIFO[K, V]) Close() { _ = "STUB: not implemented"; return }
