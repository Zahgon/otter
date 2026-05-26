package parser

import (
	"io"

	"github.com/maypok86/otter/v2/benchmarks/simulator/internal/event"
)

/*
struct {
    uint32_t timestamp;
    uint64_t obj_id;
    uint32_t obj_size;
    int64_t next_access_vtime;  // -1 if no next access
}
*/

type OracleGeneral struct {
	reader io.Reader
	buffer []byte
}

func NewOracleGeneral(reader io.Reader) *OracleGeneral { _ = "STUB: not implemented"; return nil }

func (og *OracleGeneral) Parse(send func(event event.AccessEvent) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
