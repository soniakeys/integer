package segment_test

import (
	"testing"

	"github.com/soniakeys/integer/prime/segment"
)

// a few tests on the zero object
func TestZero(t *testing.T) {
	// test Limit() on zero object
	var s segment.Sieve
	n := s.Limit()
	if n != 0 {
		t.Error("zero object Limit() = ", n)
	}
	// test Iterate succeeds on zero object
	if !s.Iterate(0, 0, func(uint64) bool {
		return false
	}) {
		t.Error("Iterate fails on zero object")
	}
	// test Iterate fails on request > limit
	if s.Iterate(0, 1, func(uint64) bool {
		return false
	}) {
		t.Error("Iterate attempts request > limit on zero object")
	}
}

func Benchmark1e4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segment.New(1e4).Iterate(1, 1e4, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segment.New(1e5).Iterate(1, 1e5, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segment.New(1e6).Iterate(1, 1e6, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segment.New(1e7).Iterate(1, 1e7, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segment.New(1e8).Iterate(1, 1e8, func(uint64) (terminate bool) {
			return
		})
	}
}
