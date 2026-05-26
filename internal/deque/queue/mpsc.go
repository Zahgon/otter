// Copyright (c) 2025 Alexey Mayshev and contributors. All rights reserved.
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

package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/maypok86/otter/v2/internal/xruntime"
)

type buffer struct {
	data []unsafe.Pointer
}

func newBuffer(capacity uint64) *buffer { _ = "STUB: not implemented"; return nil }

// MPSC is an MPSC array queue which starts at initialCapacity and grows to maxCapacity in
// linked chunks of the initial size. The queue grows only when the current buffer is full and
// elements are not copied on resize, instead a link to the new buffer is stored in the old buffer
// for the consumer to follow.
type MPSC[T any] struct {
	producerIndex    atomic.Uint64
	_                [xruntime.CacheLineSize - 8]byte
	consumerBuffer   atomic.Pointer[buffer]
	consumerIndex    atomic.Uint64
	consumerMask     atomic.Uint64
	_                [xruntime.CacheLineSize - 8*3]byte
	producerBuffer   atomic.Pointer[buffer]
	producerLimit    atomic.Uint64
	producerMask     atomic.Uint64
	_                [xruntime.CacheLineSize - 8*2]byte
	jump             unsafe.Pointer
	maxQueueCapacity uint64
}

func NewMPSC[T any](initialCapacity, maxCapacity uint32) *MPSC[T] {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // it's ok

func (m *MPSC[T]) getNextBufferSize(buffer *buffer) uint64 { _ = "STUB: not implemented"; return 0 }

func (m *MPSC[T]) getCurrentBufferCapacity(mask uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (m *MPSC[T]) availableInQueue(pIndex, cIndex uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (m *MPSC[T]) capacity() int {
	_ = "STUB: not implemented"
	//nolint:gosec // there's no overflow
	return 0
}

func (m *MPSC[T]) TryPush(t *T) bool { _ = "STUB: not implemented"; return false }

// lower bit is indicative of resize, if we see it we spin until it's cleared

// pIndex is even (lower bit is 0) -> actual index is (pIndex >> 1)

// mask/buffer may get changed by resizing -> only use for array access after successful CAS.

// a successful CAS ties the ordering, lv(pIndex)-[mask/buffer]->cas(pIndex)

// assumption behind this optimization is that queue is almost always empty or near empty

//nolint:gosec // it's ok

// We do not inline resize into this method because we do not resize on fill.
func (m *MPSC[T]) pushSlowPath(mask, pIndex, producerLimit uint64) uint8 {
	_ = "STUB: not implemented"
	// 0 - goto pIndex CAS
	return 0
}

// retry from top

// full and cannot grow

// -> return false

// -> resize

// failed resize attempt, retry from top

func (m *MPSC[T]) TryPop() *T { _ = "STUB: not implemented"; return nil }

func (m *MPSC[T]) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// NOTE: because indices are on even numbers we cannot use the size util.

// It is possible for a thread to be interrupted or reschedule between the read of the producer
// and consumer indices, therefore protection is required to ensure size is within valid range.
// In the event of concurrent polls/offers to this method the size is OVER estimated as we read
// consumer index BEFORE the producer index.

func (m *MPSC[T]) IsEmpty() bool {
	_ = "STUB: not implemented"
	// Order matters!
	// Loading consumer before producer allows for producer increments after consumer index is read.
	// This ensures this method is conservative in its estimate. Note that as this is an MPMC there
	// is nothing we can do to make this an exact method.
	return false
}

func (m *MPSC[T]) getNextBuffer(b *buffer, mask uint64) *buffer {
	_ = "STUB: not implemented"
	return nil
}

func (m *MPSC[T]) newBufferTryPush(b *buffer, index uint64) *T {
	_ = "STUB: not implemented"
	return nil
}

func (m *MPSC[T]) newBufferAndOffset(b *buffer, index uint64) uint64 {
	_ = "STUB: not implemented"
	return 0

	//nolint:gosec // there's no overflow
}

func (m *MPSC[T]) resize(oldMask uint64, oldBuffer *buffer, pIndex uint64, t *T) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // it's ok
// element in new array
//nolint:gosec // it's ok
// buffer linked

// Invalidate racing CASs
// We never set the limit beyond the bounds of a buffer

// make resize visible to the other producers

// INDEX visible before ELEMENT, consistent with consumer expectation

// make resize visible to consumer

func nextArrayOffset(mask uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// This method assumes index is actually (index << 1) because lower bit is used for resize. This
// is compensated for by reducing the element shift. The computation is constant folded, so
// there's no cost.
func modifiedCalcElementOffset(index, mask uint64) uint64 { _ = "STUB: not implemented"; return 0 }
