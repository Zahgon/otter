// Copyright (c) 2024 Alexey Mayshev and contributors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"log"
	"os"
	"sort"
	"strings"
)

type feature struct {
	name string
}

func newFeature(name string) feature { _ = "STUB: not implemented"; return *new(feature) }

func (f feature) alias() string { _ = "STUB: not implemented"; return "" }

var (
	size       = newFeature("size")
	expiration = newFeature("expiration")
	refresh    = newFeature("refresh")
	weight     = newFeature("weight")

	declaredFeatures = []feature{
		size,
		expiration,
		refresh,
		weight,
	}

	nodeTypes      []string
	aliasToFeature map[string]feature
)

func init() {
	aliasToFeature = make(map[string]feature, len(declaredFeatures))
	for _, f := range declaredFeatures {
		aliasToFeature[f.alias()] = f
	}

	enabled := make([][]bool, len(declaredFeatures))
	for i := 0; i < len(enabled); i++ {
		enabled[i] = []bool{false, true}
	}

	// cartesian product
	total := len(enabled)
	totalCombinations := 1 << total
	combinations := make([][]bool, 0, totalCombinations)
	for i := 0; i < totalCombinations; i++ {
		combination := make([]bool, 0, total)
		for j := 0; j < total; j++ {
			if ((i >> j) & 1) == 1 {
				combination = append(combination, enabled[j][0])
			} else {
				combination = append(combination, enabled[j][1])
			}
		}
		combinations = append(combinations, combination)
	}

	featureToIdx := make(map[feature]int, len(declaredFeatures))
	for i, f := range declaredFeatures {
		featureToIdx[f] = i
	}

	nodeTypesSet := make(map[string]bool, len(combinations))
	for _, combination := range combinations {
		featureSet := make(map[feature]bool)
		for i := 0; i < len(combination); i++ {
			if combination[i] {
				featureSet[declaredFeatures[i]] = true
			}
		}
		if featureSet[size] {
			delete(featureSet, weight)
		}
		features := make([]feature, 0, len(featureSet))
		for f := range featureSet {
			features = append(features, f)
		}
		sort.Slice(features, func(i, j int) bool {
			return featureToIdx[features[i]] < featureToIdx[features[j]]
		})

		var sb strings.Builder
		sb.WriteString("b")
		for _, f := range features {
			sb.WriteString(f.alias())
		}
		nodeTypesSet[sb.String()] = true
	}

	nodeTypes = make([]string, 0, len(nodeTypesSet))
	for nodeType := range nodeTypesSet {
		nodeTypes = append(nodeTypes, nodeType)
	}
	sort.Slice(nodeTypes, func(i, j int) bool {
		return nodeTypes[i] < nodeTypes[j]
	})
}

func getFeatures(nodeType string) map[feature]bool { _ = "STUB: not implemented"; return nil }

type writer struct {
	buf    bytes.Buffer
	indent string
}

func newWriter() *writer { _ = "STUB: not implemented"; return nil }

func (w *writer) p(format string, args ...any) { _ = "STUB: not implemented"; return }

func (w *writer) in() { _ = "STUB: not implemented"; return }

func (w *writer) out() { _ = "STUB: not implemented"; return }

func (w *writer) output() []byte { _ = "STUB: not implemented"; return nil }

type generator struct {
	*writer

	structName string
	features   map[feature]bool
}

func newGenerator(nodeType string) *generator { _ = "STUB: not implemented"; return nil }

func (g *generator) isBounded() bool { _ = "STUB: not implemented"; return false }

func (g *generator) withState() bool { _ = "STUB: not implemented"; return false }

func (g *generator) printImports() { _ = "STUB: not implemented"; return }

func (g *generator) printStructComment() { _ = "STUB: not implemented"; return }

//nolint:staticcheck // used only for unicode

func (g *generator) printStruct() { _ = "STUB: not implemented"; return }

// print struct definition

func (g *generator) printConstructors() { _ = "STUB: not implemented"; return }

func (g *generator) printFunctions() { _ = "STUB: not implemented"; return }

func run(nodeType, dir string) error { _ = "STUB: not implemented"; return nil }

func printManager(dir string) error { _ = "STUB: not implemented"; return nil }

func main() {
	dir := os.Args[1]

	if err := os.RemoveAll(dir); err != nil {
		log.Fatalf("remove dir: %s\n", err.Error())
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("create dir %s: %s", dir, err.Error())
	}

	for _, nodeType := range nodeTypes {
		if err := run(nodeType, dir); err != nil {
			log.Fatal(err)
		}
	}

	if err := printManager(dir); err != nil {
		log.Fatal(err)
	}
}
