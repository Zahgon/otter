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

package deque

import (
	"iter"

	"github.com/maypok86/otter/v2/internal/generated/node"
)

type Linked[K comparable, V any] struct {
	head  node.Node[K, V]
	tail  node.Node[K, V]
	len   int
	isExp bool
}

func NewLinked[K comparable, V any](isExp bool) *Linked[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (d *Linked[K, V]) PushBack(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) UpdateNode(n, old node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) PushFront(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) PopFront() node.Node[K, V] { _ = "STUB: not implemented"; return nil }

/*
func (d *Linked[K, V]) PopBack() node.Node[K, V] {
	if d.IsEmpty() {
		return nil
	}

	result := d.tail
	d.Delete(result)
	return result
}
*/

func (d *Linked[K, V]) NotContains(n node.Node[K, V]) bool { _ = "STUB: not implemented"; return false }

func (d *Linked[K, V]) Contains(n node.Node[K, V]) bool { _ = "STUB: not implemented"; return false }

func (d *Linked[K, V]) MoveToBack(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) MoveToFront(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) Delete(n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (d *Linked[K, V]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (d *Linked[K, V]) Head() node.Node[K, V] { _ = "STUB: not implemented"; return nil }

func (d *Linked[K, V]) Tail() node.Node[K, V] { _ = "STUB: not implemented"; return nil }

func (d *Linked[K, V]) All() iter.Seq[node.Node[K, V]] { _ = "STUB: not implemented"; return nil }

func (d *Linked[K, V]) Backward() iter.Seq[node.Node[K, V]] { _ = "STUB: not implemented"; return nil }

func (d *Linked[K, V]) setPrev(to, n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) setNext(to, n node.Node[K, V]) { _ = "STUB: not implemented"; return }

func (d *Linked[K, V]) getNext(n node.Node[K, V]) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (d *Linked[K, V]) getPrev(n node.Node[K, V]) node.Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}
