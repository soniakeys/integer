package sieve_test

import (
	"testing"

	"github.com/soniakeys/integer/prime/sieve"
)

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
		sieve.New(1e4)
	}
}

func Benchmark1e5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e5)
	}
}

func Benchmark1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e6)
	}
}

func Benchmark1e7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e7)
	}
}

func Benchmark1e8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.New(1e8)
	}
}
