package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

var keys []string

func round(num float64) int { _ = "STUB: not implemented"; return 0 }

func toFixed(num float64, precision int) float64 { _ = "STUB: not implemented"; return 0 }

func toMB(bytes uint64) float64 { _ = "STUB: not implemented"; return 0 }

func main() {
	name := os.Args[1]
	stringCapacity := os.Args[2]
	capacity, err := strconv.Atoi(stringCapacity)
	if err != nil {
		log.Fatal(err)
	}

	keys = make([]string, 0, capacity)
	for i := 0; i < capacity; i++ {
		keys = append(keys, strconv.Itoa(i))
	}

	constructor, ok := map[string]func(int){
		"otter":      newOtter,
		"theine":     newTheine,
		"ristretto":  newRistretto,
		"ccache":     newCcache,
		"gcache":     newGcache,
		"ttlcache":   newTTLCache,
		"golang-lru": newHashicorp,
		"sturdyc":    newSturdyc,
	}[name]
	if !ok {
		log.Fatalf("not found cache %s\n", name)
	}

	var o runtime.MemStats
	runtime.ReadMemStats(&o)

	constructor(capacity)

	// runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%s\t%d\t%v MB\t%v MB\n",
		name,
		capacity,
		toFixed(toMB(m.Alloc-o.Alloc), 2),
		toFixed(toMB(m.TotalAlloc-o.TotalAlloc), 2),
	)
}

func newOtter(capacity int) { _ = "STUB: not implemented"; return }

func newRistretto(capacity int) { _ = "STUB: not implemented"; return }

func newTheine(capacity int) { _ = "STUB: not implemented"; return }

func newCcache(capacity int) { _ = "STUB: not implemented"; return }

func newGcache(capacity int) { _ = "STUB: not implemented"; return }

func newTTLCache(capacity int) { _ = "STUB: not implemented"; return }

func newHashicorp(capacity int) { _ = "STUB: not implemented"; return }

func newSturdyc(capacity int) { _ = "STUB: not implemented"; return }
