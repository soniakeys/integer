package swing_test

import (
	"math/big"
	"testing"

	"github.com/soniakeys/integer/swing"
)

var tcs = []struct {
	n uint
	s string
}{
	{0, "1"},
	{1, "1"},
	{2, "2"},
	{3, "6"},
	{4, "6"},
	{5, "30"},
	{6, "20"},
	{7, "140"},
	{8, "70"},
	{9, "630"},
	{10, "252"},
	{31, "4808643120"},
	{32, "601080390"},
	{76, "6892620648693261354600"},
	{98, "25477612258980856902730428600"},
	{155, "227209095567131504514305357864168421386392983600"},
	{210, "90492479540310008180848641429024900729436748166512078795479440"},
	{399, "10295250013541443297297588032040198675721092538107764823484905957592333237265195859833659551897649295156404859750677412000"},
	{400, "102952500135414432972975880320401986757210925381077648234849059575923332372651958598336595518976492951564048597506774120"},
}

func TestFunction(t *testing.T) {
	var f big.Int
	for _, tc := range tcs {
		if sf := swing.SwingingFactorial(&f, tc.n).String(); sf != tc.s {
			t.Error("wrong swinging factorial for %d. Expected %s, got %s:",
				tc.n, tc.s, sf)
		}
	}
}

func TestMethod(t *testing.T) {
	s := swing.New(tcs[len(tcs)-1].n)
	var f big.Int
	for _, tc := range tcs {
		if sf := s.SwingingFactorial(&f, tc.n).String(); sf != tc.s {
			t.Error("wrong swinging factorial for %d. Expected %s, got %s:",
				tc.n, tc.s, sf)
		}
	}
}

func TestSmallOdd(t *testing.T) {
	s0 := swing.SmallOddSwing
	swing.SmallOddSwing = nil
	s := swing.New(uint(len(s0)))
	var f big.Int
	for n := 4; n < len(s0); n++ {
		if sc := s.OddSwing(&f, uint(n)).Int64(); sc != s0[n] {
			t.Errorf("SmallOddSwing(%d) expected %d, got %d", n, s0[n], sc)
		}
	}
	swing.SmallOddSwing = s0
}

func BenchmarkFunction1e2(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		swing.SwingingFactorial(&f, 1e2)
	}
}

func BenchmarkFunction1e3(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		swing.SwingingFactorial(&f, 1e3)
	}
}

func BenchmarkFunction1e4(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		swing.SwingingFactorial(&f, 1e4)
	}
}

func BenchmarkFunction1e5(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		swing.SwingingFactorial(&f, 1e5)
	}
}

func BenchmarkFunction1e6(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		swing.SwingingFactorial(&f, 1e6)
	}
}

var s = swing.New(1e6)

func BenchmarkMethod1e2(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		s.SwingingFactorial(&f, 1e2)
	}
}

func BenchmarkMethod1e3(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		s.SwingingFactorial(&f, 1e3)
	}
}

func BenchmarkMethod1e4(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		s.SwingingFactorial(&f, 1e4)
	}
}

func BenchmarkMethod1e5(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		s.SwingingFactorial(&f, 1e5)
	}
}

func BenchmarkMethod1e6(b *testing.B) {
	var f big.Int
	for i := 0; i < b.N; i++ {
		s.SwingingFactorial(&f, 1e6)
	}
}
