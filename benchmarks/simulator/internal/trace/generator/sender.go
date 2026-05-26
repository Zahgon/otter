package generator

type sender[T any] struct {
	s      Stream[T]
	events uint
	limit  *uint
}

func newSender[T any](s Stream[T], limit *uint) *sender[T] { _ = "STUB: not implemented"; return nil }

func (s *sender[T]) Send(event T) bool { _ = "STUB: not implemented"; return false }
