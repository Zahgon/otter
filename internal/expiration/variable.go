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

package expiration

import (
	"math/bits"
	"time"

	"github.com/maypok86/otter/v2/internal/generated/node"
	"github.com/maypok86/otter/v2/internal/xmath"
)

var (
	buckets = []uint64{64, 64, 32, 4, 1}
	spans   = []uint64{
		xmath.RoundUpPowerOf264(uint64((1 * time.Second).Nanoseconds())),             // 1.07s
		xmath.RoundUpPowerOf264(uint64((1 * time.Minute).Nanoseconds())),             // 1.14m
		xmath.RoundUpPowerOf264(uint64((1 * time.Hour).Nanoseconds())),               // 1.22h
		xmath.RoundUpPowerOf264(uint64((24 * time.Hour).Nanoseconds())),              // 1.63d
		buckets[3] * xmath.RoundUpPowerOf264(uint64((24 * time.Hour).Nanoseconds())), // 6.5d
		buckets[3] * xmath.RoundUpPowerOf264(uint64((24 * time.Hour).Nanoseconds())), // 6.5d
	}
	shift = []uint64{
		uint64(bits.TrailingZeros64(spans[0])),
		uint64(bits.TrailingZeros64(spans[1])),
		uint64(bits.TrailingZeros64(spans[2])),
		uint64(bits.TrailingZeros64(spans[3])),
		uint64(bits.TrailingZeros64(spans[4])),
	}
)

type Variable[K comparable, V any] struct {
	wheel [][]node.Node[K, V]
	time  uint64
}

func NewVariable[K comparable, V any](nodeManager *node.Manager[K, V]) *Variable[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// findBucket determines the bucket that the timer event should be added to.
func (v *Variable[K, V]) findBucket(expiration uint64) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Add schedules a timer event for the node.
func (v *Variable[K, V]) Add(n node.Node[K, V]) {
	_ = "STUB: not implemented"
	//nolint:gosec // there is no overflow
	return
}

// Delete removes a timer event for this entry if present.
func (v *Variable[K, V]) Delete(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (v *Variable[K, V]) DeleteExpired(nowNanos int64, expireNode func(n node.Node[K, V], nowNanos int64)) {
	_ = "STUB: not implemented"
	return
}

func (v *Variable[K, V]) deleteExpiredFromBucket(
	index int,
	prevTicks, delta uint64,
	expireNode func(n node.Node[K, V], nowNanos int64),
) {
	_ = "STUB: not implemented"
	return
}

// link adds the entry at the tail of the bucket's list.
func link[K comparable, V any](root, n node.Node[K, V]) { _ = "STUB: not implemented"; return }

// unlink removes the entry from its bucket, if scheduled.
func unlink[K comparable, V any](n node.Node[K, V]) { _ = "STUB: not implemented"; return }
