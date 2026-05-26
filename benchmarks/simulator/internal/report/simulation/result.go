package simulation

type Result struct {
	name     string
	capacity int
	ratio    float64
}

func NewResult(name string, capacity int, ratio float64) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

func (r Result) Name() string { _ = "STUB: not implemented"; return "" }

func (r Result) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r Result) Ratio() float64 { _ = "STUB: not implemented"; return 0 }
