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
	"io"
)

// LoadCacheFromFile loads cache data from the given filePath.
//
// See SaveCacheToFile for saving cache data to file.
func LoadCacheFromFile[K comparable, V any](c *Cache[K, V], filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck // it's ok

// LoadCacheFrom loads cache data from the given [io.Reader].
//
// See SaveCacheToFile for saving cache data to file.
func LoadCacheFrom[K comparable, V any](c *Cache[K, V], r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// SaveCacheToFile atomically saves cache data to the given filePath.
//
// SaveCacheToFile may be called concurrently with other operations on the cache.
//
// The saved data may be loaded with LoadCacheFromFile.
//
// WARNING: Beware that this operation is performed within the eviction policy's exclusive lock.
// While the operation is in progress further eviction maintenance will be halted.
func SaveCacheToFile[K comparable, V any](c *Cache[K, V], filePath string) error {
	_ = "STUB: not implemented"
	// Create dir if it doesn't exist.
	return nil
}

//nolint:errcheck // it's ok

// SaveCacheTo atomically saves cache data to the given [io.Writer].
//
// SaveCacheToFile may be called concurrently with other operations on the cache.
//
// The saved data may be loaded with LoadCacheFrom.
//
// WARNING: Beware that this operation is performed within the eviction policy's exclusive lock.
// While the operation is in progress further eviction maintenance will be halted.
func SaveCacheTo[K comparable, V any](c *Cache[K, V], w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
