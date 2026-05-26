package event

type AccessEvent struct {
	key uint64
}

func NewAccessEvent(key uint64) AccessEvent { _ = "STUB: not implemented"; return *new(AccessEvent) }

func (ae AccessEvent) Key() uint64 { _ = "STUB: not implemented"; return 0 }
