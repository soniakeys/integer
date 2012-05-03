package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"testing"

	"github.com/soniakeys/integer/xmath"
)

// random numbers for testing
var s terms

type terms []uint64

func (t terms) Len() int           { return len(t) }
func (t terms) Less(i, j int) bool { return t[i] < t[j] }
func (t terms) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func init() {
	// generate random terms
	s = make(terms, 500)
	for i := range s {
		s[i] = uint64(rand.Int63())
	}
	// sort them, becuase sequences in practice are typically increasing.
	sort.Sort(s)
}

func main() {
	tMin := xmath.ProductSerialThreshold - 10
	if tMin < 2 {
		tMin = 2
	}
	tMax := xmath.ProductSerialThreshold + 10
	p := new(big.Int)
	var st int
	// close on st
	f := func(b *testing.B) {
		t0 := xmath.ProductSerialThreshold
		xmath.ProductSerialThreshold = st
		for i := 0; i < b.N; i++ {
			xmath.Product(p, s)
		}
		xmath.ProductSerialThreshold = t0
	}
	for st = tMin; st <= tMax; st++ {
		fmt.Println("Threshold:", st, testing.Benchmark(f))
	}
}
