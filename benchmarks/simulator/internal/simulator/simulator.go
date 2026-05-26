package simulator

import (
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/config"
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/policy/product"
	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/report/simulation"
)

func getPolicies() map[string]product.Policy[uint64, uint64] { _ = "STUB: not implemented"; return nil }

type Simulator struct {
	cfg     config.Config
	results chan simulation.Result
}

func New(cfg config.Config) (Simulator, error) {
	_ = "STUB: not implemented"
	return *new(Simulator), nil
}

func (s Simulator) Simulate() error { _ = "STUB: not implemented"; return nil }

func (s Simulator) simulatePolicy(p policyContract, unsignedCapacity uint) error {
	_ = "STUB: not implemented"
	//nolint:gosec // there will never be an overflow
	return nil
}
