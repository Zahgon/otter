package main

import (
	"log"
	"os"
	"path/filepath"
)

type cacheInfo struct {
	cacheName string
	opsPerSec int
}

func main() {
	path := os.Args[1]
	dir := filepath.Dir(path)

	if err := run(path, dir); err != nil {
		log.Fatal(err)
	}
}

func run(path, dir string) error { _ = "STUB: not implemented"; return nil }

// for png render
