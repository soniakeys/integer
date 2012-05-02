package binomial_test

import (
	"math/big"
	"testing"

	"github.com/soniakeys/integer/binomial"
	"github.com/soniakeys/integer/prime/sieve"
)

func TestBinomial(t *testing.T) {
	var b big.Int
	for _, tc := range tcs {
		if a := binomial.Binomial(&b, tc.n, tc.k).String(); a != tc.s {
			t.Error("Binomial(%d, %d) expected %s, got %s", tc.n, tc.k, tc.s, a)
		}
	}
}

func TestBinomialS(t *testing.T) {
	p := sieve.New(uint64(tcs[len(tcs)-1].n))
	var b big.Int
	for _, tc := range tcs {
		if a := binomial.BinomialS(&b, p, tc.n, tc.k).String(); a != tc.s {
			t.Error("Binomial(%d, %d) expected %s, got %s", tc.n, tc.k, tc.s, a)
		}
	}
}

var tcs = []struct {
	n, k uint
	s    string
}{
	{0, 0, "1"},

	{1, 0, "1"},
	{1, 1, "1"},

	{2, 0, "1"},
	{2, 1, "2"},
	{2, 2, "1"},

	{3, 0, "1"},
	{3, 1, "3"},
	{3, 2, "3"},
	{3, 3, "1"},

	{123, 108, "7012067478708989884"},
	{400, 199, "102440298642203415893508338627265658464886492916495172372984138881515753604628814525708055242762679553795073231350024000"},
	{500, 50, "2314422827984300469017756871661048812545657819062792522329327913362690"},
	{1000, 998, "499500"},
}

func BenchmarkBinomial1e2(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.Binomial(&c, 1e2, uint(1e2)/3)
	}
}

func BenchmarkBinomial1e3(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.Binomial(&c, 1e3, uint(1e3)/3)
	}
}

func BenchmarkBinomial1e4(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.Binomial(&c, 1e4, uint(1e4)/3)
	}
}

func BenchmarkBinomial1e5(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.Binomial(&c, 1e5, uint(1e5)/3)
	}
}

func BenchmarkBinomial1e6(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.Binomial(&c, 1e6, uint(1e6)/3)
	}
}

func BenchmarkBinomial2e6(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.Binomial(&c, 2e6, uint(2e6)/3)
	}
}

var p = sieve.New(2e6)

func BenchmarkBinomialS1e2(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.BinomialS(&c, p, 1e2, uint(1e2)/3)
	}
}

func BenchmarkBinomialS1e3(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.BinomialS(&c, p, 1e3, uint(1e3)/3)
	}
}

func BenchmarkBinomialS1e4(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.BinomialS(&c, p, 1e4, uint(1e4)/3)
	}
}

func BenchmarkBinomialS1e5(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.BinomialS(&c, p, 1e5, uint(1e5)/3)
	}
}

func BenchmarkBinomialS1e6(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.BinomialS(&c, p, 1e6, uint(1e6)/3)
	}
}

func BenchmarkBinomialS2e6(b *testing.B) {
	var c big.Int
	for i := 0; i < b.N; i++ {
		binomial.BinomialS(&c, p, 2e6, uint(2e6)/3)
	}
}
