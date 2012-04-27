// Adapted with permission from code by Peter Luschny
// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package swing

import (
	"math/big"

	"github.com/soniakeys/integer/prime"
	"github.com/soniakeys/integer/xmath"
)

type Swing struct {
	Primes  *prime.Primes
	factors []uint64
}

func New(n uint) *Swing {
	ps := new(Swing)
	ps.Primes = prime.MakePrimes(uint64(n))

	if n >= uint(len(SmallOdd)) {
		ps.factors = make([]uint64, n)
	}

	return ps
}

func (ps *Swing) swing(m uint) *big.Int {
	if uint64(m) > ps.Primes.SieveLen {
		return nil
	}
	r := ps.OddSwing(m)
	return r.Lsh(r, xmath.BitCount(uint64(m)))
}

func (ps *Swing) OddSwing(k uint) *big.Int {
	if k < uint(len(SmallOdd)) {
		return big.NewInt(SmallOdd[k])
	}

	rootK := xmath.FloorSqrt(k)
	var i int

	ps.Primes.Iterator(3, uint64(rootK), func(p uint64) {
		q := uint64(k) / p
		for q > 0 {
			if q&1 == 1 {
				ps.factors[i] = p
				i++
			}
			q /= p
		}
	})

	ps.Primes.Iterator(uint64(rootK+1), uint64(k/3), func(p uint64) {
		if (uint64(k) / p & 1) == 1 {
			ps.factors[i] = p
			i++
		}
	})

	ps.Primes.Iterator(uint64(k/2+1), uint64(k), func(p uint64) {
		ps.factors[i] = p
		i++
	})

	return xmath.Product(ps.factors[0:i])
}

var SmallOdd = []int64{1, 1, 1, 3, 3, 15, 5,
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
