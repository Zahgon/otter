// Copyright (c) 2024 Alexey Mayshev and contributors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This is a port of lossy buffers from Caffeine.
// https://github.com/ben-manes/caffeine/blob/master/caffeine/src/main/java/com/github/benmanes/caffeine/cache/StripedBuffer.java

package lossy

import (
	"sync"
	"sync/atomic"

	"github.com/maypok86/otter/v2/internal/generated/node"
	"github.com/maypok86/otter/v2/internal/xruntime"
)

const (
	attempts = 3
)

// pool for P tokens.
var tokenPool sync.Pool

// a P token is used to point at the current OS thread (P)
// on which the goroutine is run; exact identity of the thread,
// as well as P migration tolerance, is not important since
// it's used to as a best effort mechanism for assigning
// concurrent operations (goroutines) to different stripes of
// the Adder.
type token struct {
	idx     uint32
	padding [xruntime.CacheLineSize - 4]byte
}

type striped[K comparable, V any] struct {
	buffers []atomic.Pointer[ring[K, V]]
	len     int
}

// Striped is a multiple-producer / single-consumer buffer that rejects new elements if it is full or
// fails spuriously due to contention. Unlike a queue and stack, a buffer does not guarantee an
// ordering of elements in either FIFO or LIFO order.
type Striped[K comparable, V any] struct {
	nodeManager *node.Manager[K, V]
	maxLen      int
	striped     atomic.Pointer[striped[K, V]]
	busy        atomic.Uint32
}

func NewStriped[K comparable, V any](maxLen int, nodeManager *node.Manager[K, V]) *Striped[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Add inserts the specified element into this buffer if it is possible to do so immediately without
// violating capacity restrictions. The addition is allowed to fail spuriously if multiple
// goroutines insert concurrently.
func (s *Striped[K, V]) Add(n node.Node[K, V]) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

//nolint:gosec // len will never overflow uint32

func (s *Striped[K, V]) expandOrRetry(n node.Node[K, V], t *token, wasUncontended bool) Status {
	_ = "STUB: not implemented"

	// True if last slot nonempty.
	return *new(Status)
}

//nolint:gosec // len will never overflow uint32

//nolint:gocritic // the switch statement looks even worse here

// Try to attach new buffer.

// Recheck under lock.
//nolint:gosec // len will never overflow uint32

// Slot is now non-empty.

// CAS already known to fail.
// Continue after rehash.

//nolint:gocritic // the switch statement looks even worse here

// At max size or stale.

// DrainTo drains the buffer, sending each element to the consumer for processing. The caller must ensure
// that a consumer has exclusive read access to the buffer.
func (s *Striped[K, V]) DrainTo(consumer func(n node.Node[K, V])) {
	_ = "STUB: not implemented"
	return
}

func (s *Striped[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

/*
func (s *Striped[K, V]) Clear() {
	bs := s.striped.Load()
	if bs == nil {
		return
	}
	for s.busy.Load() != 0 || !s.busy.CompareAndSwap(0, 1) {
		runtime.Gosched()
	}
	for i := 0; i < bs.len; i++ {
		b := bs.buffers[i].Load()
		if b != nil {
			b.clear()
		}
	}
	s.busy.Store(0)
}
*/
