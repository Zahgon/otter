// Copyright (c) 2024 Alexey Mayshev and contributors. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
//
// Copyright notice. Initial version of the following code was based on
// the following file from the Go Programming Language core repo:
// https://cs.opensource.google/go/go/+/refs/tags/go1.21.5:src/container/list/list_test.go
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// That can be found at https://cs.opensource.google/go/go/+/refs/tags/go1.21.5:LICENSE

package otter

import (
	"context"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/maypok86/otter/v2/internal/hashmap"
)

type call[K comparable, V any] struct {
	key        K
	value      V
	err        error
	wg         sync.WaitGroup
	isRefresh  bool
	isNotFound bool
	isFake     bool
}

func newCall[K comparable, V any](key K, isRefresh bool) *call[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *call[K, V]) Key() K { _ = "STUB: not implemented"; return *new(K) }

func (c *call[K, V]) Value() V { _ = "STUB: not implemented"; return *new(V) }

func (c *call[K, V]) AsPointer() unsafe.Pointer {
	_ = "STUB: not implemented"
	//nolint:gosec // it's ok
	return *new(unsafe.Pointer)
}

func (c *call[K, V]) cancel() { _ = "STUB: not implemented"; return }

func (c *call[K, V]) wait() { _ = "STUB: not implemented"; return }

type mapCallManager[K comparable, V any] struct{}

func (m *mapCallManager[K, V]) FromPointer(ptr unsafe.Pointer) *call[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *mapCallManager[K, V]) IsNil(c *call[K, V]) bool { _ = "STUB: not implemented"; return false }

type group[K comparable, V any] struct {
	calls         *hashmap.Map[K, V, *call[K, V]]
	initMutex     sync.Mutex
	isInitialized atomic.Bool
}

func (g *group[K, V]) init() {
	if !g.isInitialized.Load() {
		g.initMutex.Lock()
		if !g.isInitialized.Load() {
			g.calls = hashmap.New[K, V, *call[K, V]](&mapCallManager[K, V]{})
			g.isInitialized.Store(true)
		}
		g.initMutex.Unlock()
	}
}

func (g *group[K, V]) getCall(key K) *call[K, V] { _ = "STUB: not implemented"; return nil }

func (g *group[K, V]) startCall(key K, isRefresh bool) (c *call[K, V], shouldLoad bool) {
	_ = "STUB: not implemented"
	// fast path
	return nil, false
}

// double check

func (g *group[K, V]) doCall(
	ctx context.Context,
	c *call[K, V],
	load func(ctx context.Context, key K) (V, error),
	afterFinish func(c *call[K, V]),
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *group[K, V]) doBulkCall(
	ctx context.Context,
	callsInBulk map[K]*call[K, V],
	bulkLoad func(ctx context.Context, keys []K) (map[K]V, error),
	afterFinish func(c *call[K, V]),
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *group[K, V]) deleteCall(c *call[K, V]) (deleted bool) {
	_ = "STUB: not implemented"
	// fast path
	return false
}

// double check

// delete

func (g *group[K, V]) delete(key K) { _ = "STUB: not implemented"; return }
