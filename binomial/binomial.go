// Adapted with permission from code by Peter Luschny
// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package binomial

import (
	"math/big"

	"github.com/soniakeys/integer/prime"
	"github.com/soniakeys/integer/xmath"
)

func Binomial(p *prime.Primes, n, k uint) *big.Int {
	if uint64(n) > p.SieveLen {
		return nil
	}

	var r big.Int
	if k > n {
		return &r
	}

	if k > n/2 {
		k = n - k
	}

	if k < 3 {
		switch k {
		case 0:
			return r.SetInt64(1)
		case 1:
			return r.SetInt64(int64(n))
		case 2:
			var n1 big.Int
			return r.Rsh(r.Mul(r.SetInt64(int64(n)), n1.SetInt64(int64(n-1))), 1)
		}
	}

	var i int
	rootN := uint64(xmath.FloorSqrt(n))
	factors := make([]uint64, n)
	p.Iterator(2, rootN, func(p uint64) {
		var r, nn, kk uint64 = 0, uint64(n), uint64(k)
		for nn > 0 {
			if nn%p < kk%p+r {
				r = 1
				factors[i] = p
				i++
			} else {
				r = 0
			}
			nn /= p
			kk /= p
		}
	})

	p.Iterator(rootN+1, uint64(n/2), func(p uint64) {
		if uint64(n)%p < uint64(k)%p {
			factors[i] = p
			i++
		}
	})

	p.Iterator(uint64(n-k+1), uint64(n), func(p uint64) {
		factors[i] = p
		i++
	})

	return xmath.Product(factors[0:i])
}
