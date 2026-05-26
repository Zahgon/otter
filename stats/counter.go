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

package stats

import (
	"sync/atomic"
	"time"

	"github.com/maypok86/otter/v2/internal/xruntime"
	"github.com/maypok86/otter/v2/internal/xsync"
)

// Counter is a goroutine-safe [Recorder] implementation for use by otter.Cache.
type Counter struct {
	hits           *xsync.Adder
	misses         *xsync.Adder
	_              [xruntime.CacheLineSize - 16]byte
	evictions      atomic.Uint64
	evictionWeight atomic.Uint64
	_              [xruntime.CacheLineSize - 16]byte
	loadSuccesses  atomic.Uint64
	loadFailures   atomic.Uint64
	totalLoadTime  atomic.Uint64
}

// NewCounter constructs a [Counter] instance with all counts initialized to zero.
func NewCounter() *Counter { _ = "STUB: not implemented"; return nil }

// Snapshot returns a snapshot of this recorder's values. Note that this may be an inconsistent view, as it
// may be interleaved with update operations.
//
// NOTE: the values of the metrics are undefined in case of overflow. If you require specific handling, we recommend
// implementing your own [Recorder].
func (c *Counter) Snapshot() Stats { _ = "STUB: not implemented"; return *new(Stats) }

// RecordHits records cache hits. This should be called when a cache request returns a cached value.
func (c *Counter) RecordHits(count int) {
	_ = "STUB: not implemented"
	//nolint:gosec // there is no overflow
	return
}

// RecordMisses records cache misses. This should be called when a cache request returns a value that was not
// found in the cache.
func (c *Counter) RecordMisses(count int) {
	_ = "STUB: not implemented"
	//nolint:gosec // there is no overflow
	return
}

// RecordEviction records the eviction of an entry from the cache. This should only been called when an entry is
// evicted due to the cache's eviction strategy, and not as a result of manual deletions.
func (c *Counter) RecordEviction(weight uint32) { _ = "STUB: not implemented"; return }

// RecordLoadSuccess records the successful load of a new entry. This method should be called when a cache request
// causes an entry to be loaded and the loading completes successfully (either no error or otter.ErrNotFound).
func (c *Counter) RecordLoadSuccess(loadTime time.Duration) { _ = "STUB: not implemented"; return }

//nolint:gosec // there is no overflow

// RecordLoadFailure records the failed load of a new entry. This method should be called when a cache request
// causes an entry to be loaded, but the loading function returns an error that is not otter.ErrNotFound.
func (c *Counter) RecordLoadFailure(loadTime time.Duration) { _ = "STUB: not implemented"; return }

//nolint:gosec // there is no overflow
