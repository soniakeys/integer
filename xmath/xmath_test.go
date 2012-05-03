package xmath_test

import (
	"math"
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

func TestProduct(t *testing.T) {
	m := new(big.Int)

	// test empty list
	if xmath.Product(m, nil).Int64() != 1 {
		t.Error("Product of empty list should be 1. Got", m)
	}

	// compute product with seqential algorithm
	p := big.NewInt(int64(s[0]))
	for _, term := range s[1:] {
		p.Mul(p, m.SetInt64(int64(term)))
	}

	// test
	if _ = xmath.Product(m, s); m.Cmp(p) != 0 {
		t.Error("Product fail on", len(s), "random numbers")
	}
}

func BenchmarkProductM10(b *testing.B) {
	xmath.ProductSerialThreshold -= 10
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 10
}

func BenchmarkProductM9(b *testing.B) {
	xmath.ProductSerialThreshold -= 9
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 9
}

func BenchmarkProductM8(b *testing.B) {
	xmath.ProductSerialThreshold -= 8
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 8
}

func BenchmarkProductM7(b *testing.B) {
	xmath.ProductSerialThreshold -= 7
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 7
}

func BenchmarkProductM6(b *testing.B) {
	xmath.ProductSerialThreshold -= 6
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 6
}

func BenchmarkProductM5(b *testing.B) {
	xmath.ProductSerialThreshold -= 5
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 5
}

func BenchmarkProductM4(b *testing.B) {
	xmath.ProductSerialThreshold -= 4
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 4
}

func BenchmarkProductM3(b *testing.B) {
	xmath.ProductSerialThreshold -= 3
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 3
}

func BenchmarkProductM2(b *testing.B) {
	xmath.ProductSerialThreshold -= 2
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold += 2
}

func BenchmarkProductM1(b *testing.B) {
	xmath.ProductSerialThreshold--
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold++
}

func BenchmarkProductP0(b *testing.B) {
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
}

func BenchmarkProductP1(b *testing.B) {
	xmath.ProductSerialThreshold++
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold--
}

func BenchmarkProductP2(b *testing.B) {
	xmath.ProductSerialThreshold += 2
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 2
}

func BenchmarkProductP3(b *testing.B) {
	xmath.ProductSerialThreshold += 3
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 3
}

func BenchmarkProductP4(b *testing.B) {
	xmath.ProductSerialThreshold += 4
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 4
}

func BenchmarkProductP5(b *testing.B) {
	xmath.ProductSerialThreshold += 5
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 5
}

func BenchmarkProductP6(b *testing.B) {
	xmath.ProductSerialThreshold += 6
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 6
}

func BenchmarkProductP7(b *testing.B) {
	xmath.ProductSerialThreshold += 7
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 7
}

func BenchmarkProductP8(b *testing.B) {
	xmath.ProductSerialThreshold += 8
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 8
}

func BenchmarkProductP9(b *testing.B) {
	xmath.ProductSerialThreshold += 9
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 9
}

func BenchmarkProductP10(b *testing.B) {
	xmath.ProductSerialThreshold += 10
	b.Log("ProductSerialThreshold:", xmath.ProductSerialThreshold)
	p := new(big.Int)
	for i := 0; i < b.N; i++ {
		xmath.Product(p, s)
	}
	xmath.ProductSerialThreshold -= 10
}

func TestBitCount64(t *testing.T) {
	for _, tc := range s {
		// sequential algorithm
		var sc uint
		for mask := uint64(1); mask != 0; mask <<= 1 {
			if tc&mask != 0 {
				sc++
			}
		}
		// test
		if bc := xmath.BitCount64(tc); bc != sc {
			t.Fatalf("Wrong bit count for %x. Expected %d, got %d", tc, sc, bc)
		}
	}
}

func TestLog2(t *testing.T) {
	for _, tc := range s {
		u := uint(tc)
		u >>= u & 0x1f
		if u == 0 {
			continue // not interesting
		}
		l := xmath.Log2(u)
		s := u >> l
		if s != 1 {
			t.Fatalf("Log2(%d) returned %x", u, l)
		}
	}
}

func TestFloorSqrt(t *testing.T) {
	tcs := []struct {
		n, s uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{7, 2},
		{1<<20 - 1, 1<<10 - 1},
		{1 << 20, 1 << 10},
		{1<<20 + 1, 1 << 10},
		{math.MaxUint32 - 1, math.MaxUint16},
	}
	for _, tc := range tcs {
		if s := xmath.FloorSqrt(tc.n); s != tc.s {
			t.Errorf("FloorSqrt(%d) expected to be %d.  got %d", tc.n, tc.s, s)
		}
	}
}
