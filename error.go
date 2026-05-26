package otter

const (
	// ErrNotFound should be returned from a Loader.Load/Loader.Reload to indicate that an entry is
	// missing at the underlying data source. This helps the cache to determine
	// if an entry should be deleted.
	//
	// NOTE: this only applies to Cache.Get/Cache.Refresh/Loader.Load/Loader.Reload. For Cache.BulkGet/Cache.BulkRefresh,
	// this works implicitly if you return a map without the key.
	ErrNotFound strError = "otter: the entry was not found in the data source"
)

// strError allows declaring errors as constants.
type strError string

func (err strError) Error() string {
	_ = "STUB: not implemented"

	// A panicError is an arbitrary value recovered from a panic
	// with the stack trace during the execution of given function.
	return ""
}

type panicError struct {
	value any
	stack []byte
}

// Error implements error interface.
func (p *panicError) Error() string { _ = "STUB: not implemented"; return "" }

func (p *panicError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newPanicError(v any) error { _ = "STUB: not implemented"; return nil }

// The first line of the stack trace is of the form "goroutine N [status]:"
// but by the time the panic reaches cache the goroutine may no longer exist
// and its status will have changed. Trim out the misleading line.
