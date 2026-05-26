package generator

type Stream[T any] struct {
	stream chan T
}

func newStream[T any](capacity int) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s Stream[T]) close() { _ = "STUB: not implemented"; return }

func (s Stream[T]) asSender() chan<- T { _ = "STUB: not implemented"; return nil }
