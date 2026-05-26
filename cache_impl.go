// Copyright (c) 2023 Alexey Mayshev and contributors. All rights reserved.
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

package otter

import (
	"context"
	"iter"
	"sync"
	"sync/atomic"
	"time"

	"github.com/maypok86/otter/v2/internal/deque/queue"
	"github.com/maypok86/otter/v2/internal/expiration"
	"github.com/maypok86/otter/v2/internal/generated/node"
	"github.com/maypok86/otter/v2/internal/hashmap"
	"github.com/maypok86/otter/v2/internal/lossy"
	"github.com/maypok86/otter/v2/internal/xmath"
	"github.com/maypok86/otter/v2/internal/xruntime"
	"github.com/maypok86/otter/v2/stats"
)

const (
	unreachableExpiresAt     = int64(xruntime.MaxDuration)
	unreachableRefreshableAt = int64(xruntime.MaxDuration)
	noTime                   = int64(0)

	minWriteBufferSize = 4
	writeBufferRetries = 100
)

const (
	// A drain is not taking place.
	idle uint32 = 0
	// A drain is required due to a pending write modification.
	required uint32 = 1
	// A drain is in progress and will transition to idle.
	processingToIdle uint32 = 2
	// A drain is in progress and will transition to required.
	processingToRequired uint32 = 3
)

var (
	maxWriteBufferSize   uint32
	maxStripedBufferSize int
)

func init() {
	parallelism := xruntime.Parallelism()
	roundedParallelism := int(xmath.RoundUpPowerOf2(parallelism))
	//nolint:gosec // there will never be an overflow
	maxWriteBufferSize = uint32(128 * roundedParallelism)
	maxStripedBufferSize = 4 * roundedParallelism
}

func zeroValue[V any]() V {
	_ = "STUB: not implemented"

	// cache is a structure performs a best-effort bounding of a hash table using eviction algorithm
	// to determine which entries to evict when the capacity is exceeded.
	return *new(V)
}

type cache[K comparable, V any] struct {
	drainStatus        atomic.Uint32
	_                  [xruntime.CacheLineSize - 4]byte
	nodeManager        *node.Manager[K, V]
	hashmap            *hashmap.Map[K, V, node.Node[K, V]]
	evictionPolicy     *policy[K, V]
	expirationPolicy   *expiration.Variable[K, V]
	stats              stats.Recorder
	statsSnapshoter    stats.Snapshoter
	logger             Logger
	clock              timeSource
	statsClock         *realSource
	readBuffer         *lossy.Striped[K, V]
	writeBuffer        *queue.MPSC[task[K, V]]
	executor           func(fn func())
	singleflight       *group[K, V]
	evictionMutex      sync.Mutex
	doneStop           chan struct{}
	stopOnce           sync.Once
	weigher            func(key K, value V) uint32
	onDeletion         func(e DeletionEvent[K, V])
	onAtomicDeletion   func(e DeletionEvent[K, V])
	expiryCalculator   ExpiryCalculator[K, V]
	refreshCalculator  RefreshCalculator[K, V]
	taskPool           sync.Pool
	hasDefaultExecutor bool
	withTime           bool
	withExpiration     bool
	withRefresh        bool
	withEviction       bool
	isWeighted         bool
	withMaintenance    bool
	withStats          bool
}

// newCache returns a new cache instance based on the settings from Options.
func newCache[K comparable, V any](o *Options[K, V]) *cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // there's no overflow

func (c *cache[K, V]) newNode(key K, value V, old node.Node[K, V]) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache[K, V]) nodeToEntry(n node.Node[K, V], nanos int64) Entry[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// has checks if there is an item with the given key in the cache.
func (c *cache[K, V]) has(key K) bool { _ = "STUB: not implemented"; return false }

// GetIfPresent returns the value associated with the key in this cache.
func (c *cache[K, V]) GetIfPresent(key K) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// getNode returns the node associated with the key in this cache.
func (c *cache[K, V]) getNode(key K, nowNano int64) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// getNodeQuietly returns the node associated with the key in this cache.
//
// Unlike getNode, this function does not produce any side effects
// such as updating statistics or the eviction policy.
func (c *cache[K, V]) getNodeQuietly(key K, nowNano int64) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache[K, V]) afterRead(got node.Node[K, V], nowNano int64, recordHit, calcExpiresAt bool) {
	_ = "STUB: not implemented"
	return
}

// Set associates the value with the key in this cache.
//
// If the specified key is not already associated with a value, then it returns new value and true.
//
// If the specified key is already associated with a value, then it returns existing value and false.
func (c *cache[K, V]) Set(key K, value V) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// SetIfAbsent if the specified key is not already associated with a value associates it with the given value.
//
// If the specified key is not already associated with a value, then it returns new value and true.
//
// If the specified key is already associated with a value, then it returns existing value and false.
func (c *cache[K, V]) SetIfAbsent(key K, value V) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (c *cache[K, V]) calcExpiresAtAfterRead(n node.Node[K, V], nowNano int64) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) setExpiresAfterRead(n node.Node[K, V], nowNano int64, expiresAfter time.Duration) {
	_ = "STUB: not implemented"
	return
}

// GetEntry returns the cache entry associated with the key in this cache.
func (c *cache[K, V]) GetEntry(key K) (Entry[K, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetEntryQuietly returns the cache entry associated with the key in this cache.
//
// Unlike GetEntry, this function does not produce any side effects
// such as updating statistics or the eviction policy.
func (c *cache[K, V]) GetEntryQuietly(key K) (Entry[K, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetExpiresAfter specifies that the entry should be automatically removed from the cache once the duration has
// elapsed. The expiration policy determines when the entry's age is reset.
func (c *cache[K, V]) SetExpiresAfter(key K, expiresAfter time.Duration) {
	_ = "STUB: not implemented"
	return
}

// SetRefreshableAfter specifies that each entry should be eligible for reloading once a fixed duration has elapsed.
// The refresh policy determines when the entry's age is reset.
func (c *cache[K, V]) SetRefreshableAfter(key K, refreshableAfter time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) calcExpiresAtAfterWrite(n, old node.Node[K, V], nowNano int64) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) set(key K, value V, onlyIfAbsent bool) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// no op

// set

func (c *cache[K, V]) atomicSet(key K, value V, old node.Node[K, V], cl *call[K, V], nowNano int64) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

//nolint:unparam // it's ok
func (c *cache[K, V]) atomicDelete(key K, old node.Node[K, V], cl *call[K, V], nowNano int64) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Compute either sets the computed new value for the key,
// invalidates the value for the key, or does nothing, based on
// the returned [ComputeOp]. When the op returned by remappingFunc
// is [WriteOp], the value is updated to the new value. If
// it is [InvalidateOp], the entry is removed from the cache
// altogether. And finally, if the op is [CancelOp] then the
// entry is left as-is. In other words, if it did not already
// exist, it is not created, and if it did exist, it is not
// updated. This is useful to synchronously execute some
// operation on the value without incurring the cost of
// updating the cache every time.
//
// The ok result indicates whether the entry is present in the cache after the compute operation.
// The actualValue result contains the value of the cache
// if a corresponding entry is present, or the zero value otherwise.
// You can think of these results as equivalent to regular key-value lookups in a map.
//
// This call locks a hash table bucket while the compute function
// is executed. It means that modifications on other entries in
// the bucket will be blocked until the remappingFunc executes. Consider
// this when the function includes long-running operations.
func (c *cache[K, V]) Compute(
	key K,
	remappingFunc func(oldValue V, found bool) (newValue V, op ComputeOp),
) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// ComputeIfAbsent returns the existing value for the key if
// present. Otherwise, it tries to compute the value using the
// provided function. If mappingFunc returns true as the cancel value, the computation is cancelled and the zero value
// for type V is returned.
//
// The ok result indicates whether the entry is present in the cache after the compute operation.
// The actualValue result contains the value of the cache
// if a corresponding entry is present, or the zero value
// otherwise. You can think of these results as equivalent to regular key-value lookups in a map.
//
// This call locks a hash table bucket while the compute function
// is executed. It means that modifications on other entries in
// the bucket will be blocked until the valueFn executes. Consider
// this when the function includes long-running operations.
func (c *cache[K, V]) ComputeIfAbsent(
	key K,
	mappingFunc func() (newValue V, cancel bool),
) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// ComputeIfPresent returns the zero value for type V if the key is not found.
// Otherwise, it tries to compute the value using the provided function.
//
// ComputeIfPresent either sets the computed new value for the key,
// invalidates the value for the key, or does nothing, based on
// the returned [ComputeOp]. When the op returned by remappingFunc
// is [WriteOp], the value is updated to the new value. If
// it is [InvalidateOp], the entry is removed from the cache
// altogether. And finally, if the op is [CancelOp] then the
// entry is left as-is. In other words, if it did not already
// exist, it is not created, and if it did exist, it is not
// updated. This is useful to synchronously execute some
// operation on the value without incurring the cost of
// updating the cache every time.
//
// The ok result indicates whether the entry is present in the cache after the compute operation.
// The actualValue result contains the value of the cache
// if a corresponding entry is present, or the zero value
// otherwise. You can think of these results as equivalent to regular key-value lookups in a map.
//
// This call locks a hash table bucket while the compute function
// is executed. It means that modifications on other entries in
// the bucket will be blocked until the valueFn executes. Consider
// this when the function includes long-running operations.
func (c *cache[K, V]) ComputeIfPresent(
	key K,
	remappingFunc func(oldValue V) (newValue V, op ComputeOp),
) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (c *cache[K, V]) doCompute(
	key K,
	remappingFunc func(oldValue V, found bool) (newValue V, op ComputeOp),
	nowNano int64,
	recordStats bool,
) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (c *cache[K, V]) afterWrite(n, old node.Node[K, V], nowNano int64) {
	_ = "STUB: not implemented"
	return
}

// insert

// update

type refreshableKey[K comparable, V any] struct {
	key K
	old node.Node[K, V]
}

func (c *cache[K, V]) refreshKey(
	ctx context.Context,
	rk refreshableKey[K, V],
	loader Loader[K, V],
	isManual bool,
) <-chan RefreshResult[K, V] {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck // there is no need to check error

// Get returns the value associated with key in this cache, obtaining that value from loader if necessary.
// The method improves upon the conventional "if cached, return; otherwise create, cache and return" pattern.
//
// Get can return an ErrNotFound error if the Loader returns it.
// This means that the entry was not found in the data source.
//
// If another call to Get is currently loading the value for key,
// simply waits for that goroutine to finish and returns its loaded value. Note that
// multiple goroutines can concurrently load values for distinct keys.
//
// No observable state associated with this cache is modified until loading completes.
//
// WARNING: Loader.Load must not attempt to update any mappings of this cache directly.
//
// WARNING: For any given key, every loader used with it should compute the same value.
// Otherwise, a call that passes one loader may return the result of another call
// with a differently behaving loader. For example, a call that requests a short timeout
// for an RPC may wait for a similar call that requests a long timeout, or a call by an
// unprivileged user may return a resource accessible only to a privileged user making a similar call.
func (c *cache[K, V]) Get(ctx context.Context, key K, loader Loader[K, V]) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

//nolint:errcheck // there is no need to check error

func (c *cache[K, V]) calcRefreshableAt(n, old node.Node[K, V], cl *call[K, V], nowNano int64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gocritic // it's ok

func (c *cache[K, V]) afterDeleteCall(cl *call[K, V]) { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) bulkRefreshKeys(
	ctx context.Context,
	rks []refreshableKey[K, V],
	bulkLoader BulkLoader[K, V],
	isManual bool,
) <-chan []RefreshResult[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// BulkGet returns the value associated with key in this cache, obtaining that value from loader if necessary.
// The method improves upon the conventional "if cached, return; otherwise create, cache and return" pattern.
//
// If another call to Get (BulkGet) is currently loading the value for key,
// simply waits for that goroutine to finish and returns its loaded value. Note that
// multiple goroutines can concurrently load values for distinct keys.
//
// No observable state associated with this cache is modified until loading completes.
//
// WARNING: BulkLoader.BulkLoad must not attempt to update any mappings of this cache directly.
//
// WARNING: For any given key, every bulkLoader used with it should compute the same value.
// Otherwise, a call that passes one bulkLoader may return the result of another call
// with a differently behaving bulkLoader. For example, a call that requests a short timeout
// for an RPC may wait for a similar call that requests a long timeout, or a call by an
// unprivileged user may return a resource accessible only to a privileged user making a similar call.
func (c *cache[K, V]) BulkGet(ctx context.Context, keys []K, bulkLoader BulkLoader[K, V]) (map[K]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:prealloc // it's ok

func (c *cache[K, V]) wrapLoad(fn func() error) error { _ = "STUB: not implemented"; return nil }

// Refresh loads a new value for the key, asynchronously. While the new value is loading the
// previous value (if any) will continue to be returned by any Get unless it is evicted.
// If the new value is loaded successfully, it will replace the previous value in the cache;
// If refreshing returned an error, the previous value will remain,
// and the error will be logged using Logger (if it's not ErrNotFound) and swallowed. If another goroutine is currently
// loading the value for key, then this method does not perform an additional load.
//
// cache will call Loader.Reload if the cache currently contains a value for the key,
// and Loader.Load otherwise.
//
// WARNING: Loader.Load and Loader.Reload must not attempt to update any mappings of this cache directly.
//
// WARNING: For any given key, every loader used with it should compute the same value.
// Otherwise, a call that passes one loader may return the result of another call
// with a differently behaving loader. For example, a call that requests a short timeout
// for an RPC may wait for a similar call that requests a long timeout, or a call by an
// unprivileged user may return a resource accessible only to a privileged user making a similar call.
func (c *cache[K, V]) Refresh(ctx context.Context, key K, loader Loader[K, V]) <-chan RefreshResult[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// BulkRefresh loads a new value for each key, asynchronously. While the new value is loading the
// previous value (if any) will continue to be returned by any Get unless it is evicted.
// If the new value is loaded successfully, it will replace the previous value in the cache;
// If refreshing returned an error, the previous value will remain,
// and the error will be logged using Logger and swallowed. If another goroutine is currently
// loading the value for key, then this method does not perform an additional load.
//
// cache will call BulkLoader.BulkReload for existing keys, and BulkLoader.BulkLoad otherwise.
//
// WARNING: BulkLoader.BulkLoad and BulkLoader.BulkReload must not attempt to update any mappings of this cache directly.
//
// WARNING: For any given key, every bulkLoader used with it should compute the same value.
// Otherwise, a call that passes one bulkLoader may return the result of another call
// with a differently behaving loader. For example, a call that requests a short timeout
// for an RPC may wait for a similar call that requests a long timeout, or a call by an
// unprivileged user may return a resource accessible only to a privileged user making a similar call.
func (c *cache[K, V]) BulkRefresh(ctx context.Context, keys []K, bulkLoader BulkLoader[K, V]) <-chan []RefreshResult[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Invalidate discards any cached value for the key.
//
// Returns previous value if any. The invalidated result reports whether the key was
// present.
func (c *cache[K, V]) Invalidate(key K) (value V, invalidated bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (c *cache[K, V]) deleteNodeFromMap(n node.Node[K, V], nowNano int64, cause DeletionCause) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache[K, V]) deleteNode(n node.Node[K, V], nowNano int64) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) afterDelete(deleted node.Node[K, V], nowNano int64, alreadyLocked bool) {
	_ = "STUB: not implemented"
	return
}

// delete

func (c *cache[K, V]) notifyDeletion(key K, value V, cause DeletionCause) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) notifyAtomicDeletion(key K, value V, cause DeletionCause) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) periodicCleanUp() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) evictNode(n node.Node[K, V], nowNanos int64) {
	_ = "STUB: not implemented"
	return
}

func (c *cache[K, V]) nodes() iter.Seq[node.Node[K, V]] { _ = "STUB: not implemented"; return nil }

func (c *cache[K, V]) entries() iter.Seq[Entry[K, V]] { _ = "STUB: not implemented"; return nil }

// All returns an iterator over all entries in the cache.
//
// Iterator is at least weakly consistent: he is safe for concurrent use,
// but if the cache is modified (including by eviction) after the iterator is
// created, it is undefined which of the changes (if any) will be reflected in that iterator.
func (c *cache[K, V]) All() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// Keys returns an iterator over all keys in the cache.
// The iteration order is not specified and is not guaranteed to be the same from one call to the next.
//
// Iterator is at least weakly consistent: he is safe for concurrent use,
// but if the cache is modified (including by eviction) after the iterator is
// created, it is undefined which of the changes (if any) will be reflected in that iterator.
func (c *cache[K, V]) Keys() iter.Seq[K] { _ = "STUB: not implemented"; return nil }

// Values returns an iterator over all values in the cache.
// The iteration order is not specified and is not guaranteed to be the same from one call to the next.
//
// Iterator is at least weakly consistent: he is safe for concurrent use,
// but if the cache is modified (including by eviction) after the iterator is
// created, it is undefined which of the changes (if any) will be reflected in that iterator.
func (c *cache[K, V]) Values() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// InvalidateAll discards all entries in the cache. The behavior of this operation is undefined for an entry
// that is being loaded (or reloaded) and is otherwise not present.
func (c *cache[K, V]) InvalidateAll() { _ = "STUB: not implemented"; return }

// Discard all entries, falling back to one-by-one to avoid excessive lock hold times

// CleanUp performs any pending maintenance operations needed by the cache. Exactly which activities are
// performed -- if any -- is implementation-dependent.
func (c *cache[K, V]) CleanUp() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) shouldDrainBuffers(delayable bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *cache[K, V]) skipReadBuffer() bool { _ = "STUB: not implemented"; return false }

// without read buffer

func (c *cache[K, V]) afterWriteTask(t *task[K, V]) { _ = "STUB: not implemented"; return }

// In scenarios where the writing goroutines cannot make progress then they attempt to provide
// assistance by performing the eviction work directly. This can resolve cases where the
// maintenance task is scheduled but not running.

func (c *cache[K, V]) scheduleAfterWrite() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) scheduleDrainBuffers() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) drainBuffers(token *atomic.Uint32) { _ = "STUB: not implemented"; return }

// already locked

// executor is sync

// executor is async

func (c *cache[K, V]) performCleanUp(t *task[K, V]) { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) rescheduleCleanUpIfIncomplete() { _ = "STUB: not implemented"; return }

// An immediate scheduling cannot be performed on a custom executor because it may use a
// caller-runs policy. This could cause the caller's penalty to exceed the amortized threshold,
// e.g. repeated concurrent writes could result in a retry loop.

func (c *cache[K, V]) maintenance(t *task[K, V]) { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) drainReadBuffer() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) drainWriteBuffer() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) runTask(t *task[K, V]) { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) onAccess(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) expireNodes() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) evictNodes() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) climb() { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) getTask(n, old node.Node[K, V], writeReason reason, cause DeletionCause) *task[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache[K, V]) putTask(t *task[K, V]) { _ = "STUB: not implemented"; return }

// SetMaximum specifies the maximum total size of this cache. This value may be interpreted as the weighted
// or unweighted threshold size based on how this cache was constructed. If the cache currently
// exceeds the new maximum size this operation eagerly evict entries until the cache shrinks to
// the appropriate size.
func (c *cache[K, V]) SetMaximum(maximum uint64) { _ = "STUB: not implemented"; return }

// GetMaximum returns the maximum total weighted or unweighted size of this cache, depending on how the
// cache was constructed. If this cache does not use a (weighted) size bound, then the method will return math.MaxUint64.
func (c *cache[K, V]) GetMaximum() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *cache[K, V]) StopAllGoroutines() bool { _ = "STUB: not implemented"; return false }

// EstimatedSize returns the approximate number of entries in this cache. The value returned is an estimate; the
// actual count may differ if there are concurrent insertions or deletions, or if some entries are
// pending deletion due to expiration. In the case of stale entries
// this inaccuracy can be mitigated by performing a CleanUp first.
func (c *cache[K, V]) EstimatedSize() int { _ = "STUB: not implemented"; return 0 }

// IsWeighted returns whether the cache is bounded by a maximum size or maximum weight.
func (c *cache[K, V]) IsWeighted() bool { _ = "STUB: not implemented"; return false }

// IsRecordingStats returns whether the cache statistics are being accumulated.
func (c *cache[K, V]) IsRecordingStats() bool {
	_ = "STUB: not implemented"

	// Stats returns a current snapshot of this cache's cumulative statistics.
	// All statistics are initialized to zero and are monotonically increasing over the lifetime of the cache.
	// Due to the performance penalty of maintaining statistics,
	// some implementations may not record the usage history immediately or at all.
	//
	// NOTE: If your [stats.Recorder] implementation doesn't also implement [stats.Snapshoter],
	// this method will always return a zero-value snapshot.
	return false
}

func (c *cache[K, V]) Stats() stats.Stats { _ = "STUB: not implemented"; return *new(stats.Stats) }

// WeightedSize returns the approximate accumulated weight of entries in this cache. If this cache does not
// use a weighted size bound, then the method will return 0.
func (c *cache[K, V]) WeightedSize() uint64 { _ = "STUB: not implemented"; return 0 }

// Hottest returns an iterator for ordered traversal of the cache entries. The order of
// iteration is from the entries most likely to be retained (hottest) to the entries least
// likely to be retained (coldest). This order is determined by the eviction policy's best guess
// at the start of the iteration.
//
// WARNING: Beware that this iteration is performed within the eviction policy's exclusive lock, so the
// iteration should be short and simple. While the iteration is in progress further eviction
// maintenance will be halted.
func (c *cache[K, V]) Hottest() iter.Seq[Entry[K, V]] { _ = "STUB: not implemented"; return nil }

// Coldest returns an iterator for ordered traversal of the cache entries. The order of
// iteration is from the entries least likely to be retained (coldest) to the entries most
// likely to be retained (hottest). This order is determined by the eviction policy's best guess
// at the start of the iteration.
//
// WARNING: Beware that this iteration is performed within the eviction policy's exclusive lock, so the
// iteration should be short and simple. While the iteration is in progress further eviction
// maintenance will be halted.
func (c *cache[K, V]) Coldest() iter.Seq[Entry[K, V]] { _ = "STUB: not implemented"; return nil }

func (c *cache[K, V]) evictionOrder(hottest bool) iter.Seq[Entry[K, V]] {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache[K, V]) makeRetired(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (c *cache[K, V]) makeDead(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func getCause[K comparable, V any](n node.Node[K, V], nowNano int64, cause DeletionCause) DeletionCause {
	_ = "STUB: not implemented"
	return *new(DeletionCause)
}
