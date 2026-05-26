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

package otter

import (
	"sync"
	"sync/atomic"
	"time"
)

// Clock is a time source that
//   - Returns a time value representing the number of nanoseconds elapsed since some
//     fixed but arbitrary point in time
//   - Returns a channel that delivers “ticks” of a clock at intervals.
type Clock interface {
	// NowNano returns the number of nanoseconds elapsed since this clock's fixed point of reference.
	//
	// By default, time.Now().UnixNano() is used.
	NowNano() int64
	// Tick returns a channel that delivers “ticks” of a clock at intervals.
	//
	// The cache uses this method only for proactive expiration and calls Tick(time.Second) in a separate goroutine.
	//
	// By default, [time.Tick] is used.
	Tick(duration time.Duration) <-chan time.Time
}

type timeSource interface {
	Clock
	Init()
	Sleep(duration time.Duration)
	ProcessTick()
}

func newTimeSource(clock Clock) timeSource { _ = "STUB: not implemented"; return *new(timeSource) }

type customSource struct {
	clock         Clock
	isInitialized atomic.Bool
}

func newCustomSource(clock Clock) *customSource { _ = "STUB: not implemented"; return nil }

func (cs *customSource) Init() { _ = "STUB: not implemented"; return }

func (cs *customSource) NowNano() int64 { _ = "STUB: not implemented"; return 0 }

func (cs *customSource) Tick(duration time.Duration) <-chan time.Time {
	_ = "STUB: not implemented"
	return nil
}

func (cs *customSource) Sleep(duration time.Duration) { _ = "STUB: not implemented"; return }

func (cs *customSource) ProcessTick() { _ = "STUB: not implemented"; return }

type realSource struct {
	initMutex     sync.Mutex
	isInitialized atomic.Bool
	start         time.Time
	startNanos    atomic.Int64
}

func (c *realSource) Init() { _ = "STUB: not implemented"; return }

func (c *realSource) NowNano() int64 { _ = "STUB: not implemented"; return 0 }

func (c *realSource) Tick(duration time.Duration) <-chan time.Time {
	_ = "STUB: not implemented"
	return nil
}

func (c *realSource) Sleep(duration time.Duration) { _ = "STUB: not implemented"; return }

func (c *realSource) ProcessTick() { _ = "STUB: not implemented"; return }

type fakeSource struct {
	mutex          sync.Mutex
	now            time.Time
	initOnce       sync.Once
	sleeps         chan time.Duration
	tickWg         sync.WaitGroup
	sleepWg        sync.WaitGroup
	firstSleep     atomic.Bool
	withTick       atomic.Bool
	ticker         chan time.Time
	enableTickOnce sync.Once
	enableTick     chan time.Duration
}

func (f *fakeSource) Init() { _ = "STUB: not implemented"; return }

func (f *fakeSource) NowNano() int64 { _ = "STUB: not implemented"; return 0 }

func (f *fakeSource) Tick(d time.Duration) <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (f *fakeSource) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }

func (f *fakeSource) getNow() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *fakeSource) ProcessTick() { _ = "STUB: not implemented"; return }
