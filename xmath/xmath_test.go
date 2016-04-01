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
	// sort them, because sequences in practice are typically increasing.
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

func TestFloorSqrt32(t *testing.T) {
	tcs := []struct {
		n, s uint32
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
		if s := xmath.FloorSqrt32(tc.n); s != tc.s {
			t.Errorf("FloorSqrt(%d) expected to be %d.  got %d", tc.n, tc.s, s)
		}
	}
}

func TestFloorSqrt64(t *testing.T) {
	tcs := []struct {
		n, s uint64
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
		{math.MaxUint64 - 1, math.MaxUint32},
	}
	for _, tc := range tcs {
		if s := xmath.FloorSqrt64(tc.n); s != tc.s {
			t.Errorf("FloorSqrt64(%d) expected to be %d.  got %d", tc.n, tc.s, s)
		}
	}
}

func TestTrailingZeros(t *testing.T) {
	for _, tc := range s {
		tcu := uint(tc)
		// seqential algorithm
		var tz byte
		if tcu != 0 {
			for nz := tcu; nz&1 == 0; nz >>= 1 {
				tz++
			}
		}
		// test
		if cz := xmath.TrailingZeros(tcu); cz != tz {
			t.Errorf("TrailingZero(%x) expected %d, got %d", tcu, tz, cz)
		}
	}
}

func TestTrailingZeros32(t *testing.T) {
	for _, tc := range s {
		tc32 := uint32(tc)
		// seqential algorithm
		var tz byte
		if tc32 != 0 {
			for nz := tc32; nz&1 == 0; nz >>= 1 {
				tz++
			}
		}
		// test
		if cz := xmath.TrailingZeros32(tc32); cz != tz {
			t.Errorf("TrailingZero32(%x) expected %d, got %d", tc32, tz, cz)
		}
	}
}

func TestTrailingZeros64(t *testing.T) {
	for _, tc := range s {
		// seqential algorithm
		var tz byte
		if tc != 0 {
			for nz := tc; nz&1 == 0; nz >>= 1 {
				tz++
			}
		}
		// test
		if cz := xmath.TrailingZeros64(tc); cz != tz {
			t.Errorf("TrailingZero64(%x) expected %d, got %d", tc, tz, cz)
		}
	}
}

func TestTrailingZerosBig(t *testing.T) {
	t1 := func(v *big.Int, want int) {
		if got := xmath.TrailingZerosBig(v); got != want {
			t.Fatalf("TrailingZeroBig(0x%x) = %d, want %d", v, got, want)
		}
	}
	// boundary case 0
	t1(&big.Int{}, 0)
	// powers of 2 up to 64
	var v big.Int
	for p := 0; p < 64; p++ {
		t1(v.SetInt64(1<<uint(p)), p)
	}
	// some bigger powers of 2
	one := big.NewInt(1)
	for _, p := range []int{64, 100, 1000, 10000} {
		t1(v.Lsh(one, uint(p)), p)
	}
}

func TestTrailingOnesBig(t *testing.T) {
	t1 := func(v *big.Int, want int) {
		if got := xmath.TrailingOnesBig(v); got != want {
			t.Fatalf("TrailingOnesBig(0x%x) = %d, want %d", v, got, want)
		}
	}
	// (powers of 2 up to 64) - 1
	// thats 0, 1, 11, 111, and so on.
	var v big.Int
	for p := 0; p < 64; p++ {
		t1(v.SetInt64(1<<uint(p)-1), p)
	}
	// (some bigger powers of 2) - 1
	one := big.NewInt(1)
	for _, p := range []int{64, 100, 1000, 10000} {
		t1(v.Sub(v.Lsh(one, uint(p)), one), p)
	}
}
