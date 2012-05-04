package sprp_test

import (
	"testing"

	"github.com/soniakeys/integer/prime/sprp"
)

func TestLimit(t *testing.T) {
	if l := sprp.New().Limit(); l != 1<<32-1 {
		t.Errorf("Limit() returned %d. 1<<32-1 expected.", l)
	}
}

func TestBoundaryCases(t *testing.T) {
	// source: prime pages
	b12 := []uint64{5309, 5323, 5333, 5347}                        // 1-2 bases
	b23 := []uint64{316349261, 316349279, 316349329, 316349333}    // 2-3 bases	
	b32 := []uint64{1<<32 - 99, 1<<32 - 65, 1<<32 - 17, 1<<32 - 5} // near the top

	s := sprp.New()
	for _, p := range [][]uint64{b12, b23, b32} {
		i = 0
		s.Iterate(p[0], p[len(p)-1], func(prime uint64) bool {
			if prime != p[i] {
				t.Fatal("Incorrect prime.  Expected", p[i], "got", prime)
			}
			i++
			return false
		})
	}
}

func Benchmark1e4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprp.New().Iterate(1, 1e4, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprp.New().Iterate(1, 1e5, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprp.New().Iterate(1, 1e6, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprp.New().Iterate(1, 1e7, func(uint64) (terminate bool) {
			return
		})
	}
}
