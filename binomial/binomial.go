// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package binomial computes binomial coefficients C(n,k)
// with an efficient algorithm using prime numbers.
package binomial

import (
	"math/big"

	"github.com/soniakeys/integer/prime/sieve"
	"github.com/soniakeys/integer/xmath"
)

// Binomal computes the binomial coefficient C(n,k) leaving the result in z.
// It replaces the existing value of z and returns z.
func Binomial(z *big.Int, n, k uint) *big.Int {
	if k > n {
		return z.SetInt64(0)
	}
	return BinomialS(z, sieve.New(uint64(n)), n, k)
}

// BinomialS computes the binomial coefficient C(n,k) using prime number
// sieve p.  BinomialS returns nil if p is too small.  Otherwise it leaves
// the result in z, replacing the existing value of z, and returning z.
func BinomialS(z *big.Int, p *sieve.Sieve, n, k uint) *big.Int {
	if uint64(n) > p.Len {
		return nil
	}
	if k > n {
		return z.SetInt64(0)
	}
	if k > n/2 {
		k = n - k
	}
	if k < 3 {
		switch k {
		case 0:
			return z.SetInt64(1)
		case 1:
			return z.SetInt64(int64(n))
		case 2:
			var n1 big.Int
			return z.Rsh(z.Mul(z.SetInt64(int64(n)), n1.SetInt64(int64(n-1))), 1)
		}
	}
	rootN := uint64(xmath.FloorSqrt(n))
	var factors []uint64
	p.Iterate(2, rootN, func(p uint64) (terminate bool) {
		var r, nn, kk uint64 = 0, uint64(n), uint64(k)
		for nn > 0 {
			if nn%p < kk%p+r {
				r = 1
				factors = append(factors, p)
			} else {
				r = 0
			}
			nn /= p
			kk /= p
		}
		return
	})
	p.Iterate(rootN+1, uint64(n/2), func(p uint64) (terminate bool) {
		if uint64(n)%p < uint64(k)%p {
			factors = append(factors, p)
		}
		return
	})
	p.Iterate(uint64(n-k+1), uint64(n), func(p uint64) (terminate bool) {
		factors = append(factors, p)
		return
	})
	return xmath.Product(z, factors)
}
