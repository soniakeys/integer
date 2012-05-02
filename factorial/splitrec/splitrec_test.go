package splitrec_test

import (
	"math/big"
	"testing"

	"github.com/soniakeys/integer/factorial/splitrec"
)

var tcs = []struct {
	n uint
	s string
}{
	{0, "1"},
	{1, "1"},
	{2, "2"},
	{3, "6"},
	{4, "24"},
	{5, "120"},
	{100, "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"},
}

func TestSplitRecursive(t *testing.T) {
	var f big.Int
	for _, tc := range tcs {
		if fs := splitrec.Factorial(&f, tc.n).String(); fs != tc.s {
			t.Errorf("%d! incorrect.  expected %s, got %s", tc.n, tc.s, fs)
		}
	}
}

func Benchmark1e2(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		splitrec.Factorial(&f, 1e2)
	}
}

func Benchmark1e3(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		splitrec.Factorial(&f, 1e3)
	}
}

func Benchmark1e4(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		splitrec.Factorial(&f, 1e4)
	}
}

func Benchmark1e5(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		splitrec.Factorial(&f, 1e5)
	}
}
