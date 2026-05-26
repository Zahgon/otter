package config

type Config struct {
	Type       string   `toml:"type"`
	Name       string   `toml:"name"`
	Capacities []uint   `toml:"capacities"`
	Caches     []string `toml:"caches"`
	Limit      *uint    `toml:"limit"`

	Zipf *Zipf `toml:"zipf"`
	File *File `toml:"file"`
}

func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }

type Zipf struct {
	S    float64 `toml:"s"`
	V    float64 `toml:"v"`
	IMAX uint64  `toml:"imax"`
}

func (z *Zipf) validate() error { _ = "STUB: not implemented"; return nil }

type FilePath struct {
	TraceType string `toml:"trace_type"`
	Path      string `toml:"path"`
}

func (fp *FilePath) validate() error { _ = "STUB: not implemented"; return nil }

type File struct {
	Paths []FilePath `toml:"paths"`
}

func (f *File) validate() error { _ = "STUB: not implemented"; return nil }

func Load(configPath string) (Config, error) { _ = "STUB: not implemented"; return *new(Config), nil }
