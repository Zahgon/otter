package generator

type File struct {
	base
}

type FilePath struct {
	TraceType string
	Path      string
}

func NewFile(paths []FilePath, limit *uint) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
