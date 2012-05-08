// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package swing implements efficient generation of swinging factorials.
package swing

import (
	"math/big"

	"github.com/soniakeys/integer/prime/sieve"
	"github.com/soniakeys/integer/xmath"
)

// SwingingFactorial computes n≀, OEIS A056040.
//
// Result is computed in z, replacing the value of z.  z is returned.
func SwingingFactorial(z *big.Int, n uint) *big.Int {
	return New(n).SwingingFactorial(z, n)
}

// Swing type is useful for generating multiple swinging factorials
// after generating underlying prime sieve just once.
type Swing struct {
	Sieve *sieve.Sieve
	// factors holds intermediate results.  It grows as needed and
	// is maintained as a member to avoid repeated reallocation.
	factors []uint64
}

// New is a constructor that generates the underlying prime sieve.
// (If you have a sieve already, you can just assign the Sieve member of
// a zero value Swing object.)
func New(n uint) *Swing {
	return &Swing{Sieve: sieve.New(uint64(n))}
}

// SwingingFactorial member computes n≀ on a Swing object.
func (ps *Swing) SwingingFactorial(z *big.Int, n uint) *big.Int {
	if uint64(n) > ps.Sieve.Lim {
		return nil
	}
	return z.Lsh(ps.OddSwing(z, n), xmath.BitCount32(uint32(n>>1)))
}

func (ps *Swing) OddSwing(z *big.Int, k uint) *big.Int {
	if k < uint(len(SmallOddSwing)) {
		return z.SetInt64(SmallOddSwing[k])
	}
	rootK := xmath.FloorSqrt(k)
	ps.factors = ps.factors[:0] // reset length, reusing existing capacity
	ps.Sieve.Iterate(3, uint64(rootK), func(p uint64) (terminate bool) {
		q := uint64(k) / p
		for q > 0 {
			if q&1 == 1 {
				ps.factors = append(ps.factors, p)
			}
			q /= p
		}
		return
	})
	ps.Sieve.Iterate(uint64(rootK+1), uint64(k/3), func(p uint64) (term bool) {
		if (uint64(k) / p & 1) == 1 {
			ps.factors = append(ps.factors, p)
		}
		return
	})
	ps.Sieve.Iterate(uint64(k/2+1), uint64(k), func(p uint64) (term bool) {
		ps.factors = append(ps.factors, p)
		return
	})
	return xmath.Product(z, ps.factors)
}

var SmallOddSwing = []int64{1, 1, 1, 3, 3, 15, 5,
	35, 35, 315, 63, 693, 231, 3003, 429, 6435, 6435,
	109395, 12155, 230945, 46189, 969969, 88179, 2028117, 676039,
	16900975, 1300075, 35102025, 5014575, 145422675, 9694845,
	300540195, 300540195, 9917826435, 583401555, 20419054425,
	2268783825, 83945001525, 4418157975, 172308161025,
	34461632205, 1412926920405, 67282234305, 2893136075115,
	263012370465, 11835556670925, 514589420475, 24185702762325,
	8061900920775, 395033145117975, 15801325804719,
	805867616040669, 61989816618513, 3285460280781189,
	121683714103007, 6692604275665385, 956086325095055,
	54496920530418135, 1879204156221315, 110873045217057585,
	7391536347803839, 450883717216034179, 14544636039226909,
	916312070471295267, 916312070471295267,
}

var SmallOddFactorial = []int64{1, 1, 1, 3, 3,
	15, 45, 315, 315, 2835, 14175, 155925, 467775,
	6081075, 42567525, 638512875, 638512875, 10854718875,
	97692469875, 1856156927625, 9280784638125, 194896477400625,
	2143861251406875, 49308808782358125, 147926426347074375,
	3698160658676859375,
}
