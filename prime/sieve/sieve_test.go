package sieve_test

import (
	"testing"

	"github.com/soniakeys/integer/prime/sieve"
)

// a few tests on the zero object
func TestZero(t *testing.T) {
	// test Limit() on zero object
	var s sieve.Sieve
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

var tcs = []struct {
	n, nPrimes uint64
}{
	{0, 0},
	{1, 0},
	{2, 1},
	{3, 2},
	{4, 2},
	{5, 3},
	{6, 3},
	{7, 4},
	{8, 4},
	{100000000, 5761455},
}

// a little different that the generic test a directory up. creates sieves
// with different limits, starts iterating at 1, verifies number of primes.
func TestSieve(t *testing.T) {
	for _, tc := range tcs {
		var count uint64
		sieve.New(tc.n).Iterate(1, tc.n, func(uint64) (terminate bool) {
			count++
			return
		})
		if count != tc.nPrimes {
			t.Errorf("Wrong number of primes <= %d.  expected %d, found %d",
				tc.n, tc.nPrimes, count)
		}
	}
}

func Benchmark1e4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e4).Iterate(1, 1e4, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e5).Iterate(1, 1e5, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e6).Iterate(1, 1e6, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e7).Iterate(1, 1e7, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e8).Iterate(1, 1e8, func(uint64) (terminate bool) {
			return
		})
	}
}
