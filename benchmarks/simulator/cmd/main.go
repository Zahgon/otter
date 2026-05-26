package main

import (
	"flag"
	"log"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "configs/zipf.toml", "Path to configuration file")
	flag.Parse()

	if err := run(configPath); err != nil {
		log.Fatal(err)
	}
}

func run(configPath string) error { _ = "STUB: not implemented"; return nil }
