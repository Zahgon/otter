// Copyright (c) 2024 Alexey Mayshev and contributors. All rights reserved.
// Copyright (c) 2021 Andrey Pechkurov. All rights reserved.
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
// Copyright notice. This code is a fork of xsync.MapOf from this file with some changes:
// https://github.com/puzpuzpuz/xsync/blob/main/mapof_test.go
//
// Use of this source code is governed by a MIT license that can be found
// at https://github.com/puzpuzpuz/xsync/blob/main/LICENSE

package hashmap

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/maypok86/otter/v2/internal/xruntime"
)

type mapResizeHint int

const (
	mapGrowHint   mapResizeHint = 0
	mapShrinkHint mapResizeHint = 1
	mapClearHint  mapResizeHint = 2
)

const (
	// number of Map nodes per bucket; 5 nodes lead to size of 64B
	// (one cache line) on 64-bit machines.
	nodesPerMapBucket        = 5
	defaultMeta       uint64 = 0x8080808080808080
	metaMask          uint64 = 0xffffffffff
	defaultMetaMasked        = defaultMeta & metaMask
	emptyMetaSlot     uint8  = 0x80

	// threshold fraction of table occupation to start a table shrinking
	// when deleting the last entry in a bucket chain.
	mapShrinkFraction = 128
	// map load factor to trigger a table resize during insertion;
	// a map holds up to mapLoadFactor*nodesPerMapBucket*mapTableLen
	// key-value pairs (this is a soft limit).
	mapLoadFactor = 0.75
	// minimal table size, i.e. number of buckets; thus, minimal map
	// capacity can be calculated as nodesPerMapBucket*defaultMinMapTableLen.
	defaultMinMapTableLen = 32
	// minimum counter stripes to use.
	minMapCounterLen = 8
	// maximum counter stripes to use; stands for around 4KB of memory.
	maxMapCounterLen = 32
	// minimum buckets per goroutine during parallel resize.
	minBucketsPerGoroutine = 64
)

// Map is like a Go map[K]V but is safe for concurrent
// use by multiple goroutines without additional locking or
// coordination. It follows the interface of sync.Map with
// a number of valuable extensions like Compute or Size.
//
// A Map must not be copied after first use.
//
// Map uses a modified version of Cache-Line Hash Table (CLHT)
// data structure: https://github.com/LPD-EPFL/CLHT
//
// CLHT is built around idea to organize the hash table in
// cache-line-sized buckets, so that on all modern CPUs update
// operations complete with at most one cache-line transfer.
// Also, Get operations involve no write to memory, as well as no
// mutexes or any other sort of locks. Due to this design, in all
// considered scenarios Map outperforms sync.Map.
//
// Map also borrows ideas from Java's j.u.c.ConcurrentHashMap
// (immutable K/V pair structs instead of atomic snapshots)
// and C++'s absl::flat_hash_map (meta memory and SWAR-based
// lookups).
type Map[K comparable, V any, N mapNode[K, V]] struct {
	totalGrowths atomic.Int64
	totalShrinks atomic.Int64
	resizing     atomic.Bool                 // resize in progress flag
	resizeMu     sync.Mutex                  // only used along with resizeCond
	resizeCond   sync.Cond                   // used to wake up resize waiters (concurrent modifications)
	table        atomic.Pointer[mapTable[K]] // *mapTable
	nodeManager  mapNodeManager[K, V, N]
	minTableLen  int
}

type counterStripe struct {
	c int64
	//lint:ignore U1000 prevents false sharing
	pad [xruntime.CacheLineSize - 8]byte
}

type mapTable[K comparable] struct {
	buckets []bucketPadded
	// striped counter for number of table nodes;
	// used to determine if a table shrinking is needed
	// occupies min(buckets_memory/1024, 64KB) of memory
	size   []counterStripe
	hasher xruntime.Hasher[K]
}

// bucketPadded is a CL-sized map bucket holding up to
// nodesPerMapBucket nodes.
type bucketPadded struct {
	//lint:ignore U1000 ensure each bucket takes two cache lines on both 32 and 64-bit archs
	pad [64 - unsafe.Sizeof(bucket{})]byte
	bucket
}

type bucket struct {
	meta  atomic.Uint64
	nodes [nodesPerMapBucket]unsafe.Pointer // node.Node
	next  atomic.Pointer[bucketPadded]
	mu    sync.Mutex
}

// NewWithSize creates a new Map instance with capacity enough
// to hold size nodes. If size is zero or negative, the value
// is ignored.
func NewWithSize[K comparable, V any, N mapNode[K, V]](nodeManager mapNodeManager[K, V, N], size int) *Map[K, V, N] {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new Map instance.
func New[K comparable, V any, N mapNode[K, V]](nodeManager mapNodeManager[K, V, N]) *Map[K, V, N] {
	_ = "STUB: not implemented"
	return nil
}

func newMap[K comparable, V any, N mapNode[K, V]](nodeManager mapNodeManager[K, V, N], sizeHint int) *Map[K, V, N] {
	_ = "STUB: not implemented"
	return nil
}

func newMapTable[K comparable](minTableLen int) *mapTable[K] { _ = "STUB: not implemented"; return nil }

func zeroValue[V any]() V {
	_ = "STUB: not implemented"
	return *

	// Get returns the node stored in the map for a key, or nil
	// if no value is present.
	new(V)
}

func (m *Map[K, V, N]) Get(key K) N { _ = "STUB: not implemented"; return *new(N) }

//nolint:gosec // there is no overflow

// Compute either sets the computed new value for the key or deletes
// the value for the key.
//
// This call locks a hash table bucket while the compute function
// is executed. It means that modifications on other nodes in
// the bucket will be blocked until the computeFn executes. Consider
// this when the function includes long-running operations.
func (m *Map[K, V, N]) Compute(key K, computeFunc func(n N) N) N {
	_ = "STUB: not implemented"
	return *new(N)
}

//nolint:gosec // there is no overflow

// The following two checks must go in reverse to what's
// in the resize method.

// Resize is in progress. Wait, then go for another attempt.

// Someone resized the table. Go for another attempt.

// In-place update/delete.

// oldNode != nil

// Deletion.
// First we update the hash, then the node.

// Might need to shrink the table if we left bucket empty.

// Search for empty nodes (up to 5 per bucket).

// Insertion into an existing bucket.

// oldNode == nil.

// no op.

// First we update meta, then the node.

// Need to grow the table. Then go for another attempt.

// Insertion into a new bucket.

// oldNode == nil

// Create and append a bucket.

func (m *Map[K, V, N]) newerTableExists(table *mapTable[K]) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Map[K, V, N]) resizeInProgress() bool { _ = "STUB: not implemented"; return false }

func (m *Map[K, V, N]) waitForResize() { _ = "STUB: not implemented"; return }

func (m *Map[K, V, N]) resize(knownTable *mapTable[K], hint mapResizeHint) {
	_ = "STUB: not implemented"
	return
}

// Fast path for shrink attempts.

// Slow path.

// Someone else started resize. Wait for it to finish.

// Grow the table with factor of 2.

// Shrink the table with factor of 2.

// No need to shrink. Wake up all waiters and give up.

// Copy the data only if we're not clearing the map.

// Enable parallel resizing when serialResize is false and table is large enough.
// Calculate optimal goroutine count based on table size and available CPUs

//nolint:gosec // there is no overflow

//nolint:gosec // there is no overflow

// Publish the new table and wake up all waiters.

func (m *Map[K, V, N]) copyBucketWithDestLock(b *bucketPadded, destTable *mapTable[K]) (copied int) {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec // there is no overflow

func (m *Map[K, V, N]) copyBucket(b *bucketPadded, destTable *mapTable[K]) (copied int) {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gocritic // nesting is normal here

//nolint:gosec // there is no overflow

// Range calls f sequentially for each key and value present in the
// map. If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot
// of the Map's contents: no key will be visited more than once, but
// if the value for any key is stored or deleted concurrently, Range
// may reflect any mapping for that key from any point during the
// Range call.
//
// It is safe to modify the map while iterating it, including entry
// creation, modification and deletion. However, the concurrent
// modification rule apply, i.e. the changes may be not reflected
// in the subsequently iterated nodes.
func (m *Map[K, V, N]) Range(fn func(n N) bool) { _ = "STUB: not implemented"; return }

// Pre-allocate array big enough to fit nodes for most hash tables.

// Prevent concurrent modifications and copy all nodes into
// the intermediate slice.

// Call the function for all copied nodes.

// Remove the reference to avoid preventing the copied
// nodes from being GCed until this method finishes.

// Clear deletes all keys and values currently stored in the map.
func (m *Map[K, V, N]) Clear() { _ = "STUB: not implemented"; return }

// Size returns current size of the map.
func (m *Map[K, V, N]) Size() int { _ = "STUB: not implemented"; return 0 }

func appendToBucket(h2 uint8, nodePtr unsafe.Pointer, b *bucketPadded) {
	_ = "STUB: not implemented"
	return
}

func (table *mapTable[K]) addSize(bucketIdx uint64, delta int) {
	_ = "STUB: not implemented"
	//nolint:gosec // there is no overflow
	return
}

func (table *mapTable[K]) addSizePlain(bucketIdx uint64, delta int) {
	_ = "STUB: not implemented"
	//nolint:gosec // there is no overflow
	return
}

func (table *mapTable[K]) sumSize() int64 { _ = "STUB: not implemented"; return 0 }

func h1(h uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func h2(h uint64) uint8 {
	_ = "STUB: not implemented"
	//nolint:gosec // there is no overflow
	return 0
}

func broadcast(b uint8) uint64 { _ = "STUB: not implemented"; return 0 }

func firstMarkedByteIndex(w uint64) int { _ = "STUB: not implemented"; return 0 }

// SWAR byte search: may produce false positives, e.g. for 0x0100,
// so make sure to double-check bytes found by this function.
func markZeroBytes(w uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func setByte(w uint64, b uint8, idx int) uint64 { _ = "STUB: not implemented"; return 0 }
