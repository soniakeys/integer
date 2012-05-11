// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package double has an efficient algorithm for the double factorial. (n!!)
package double

import (
	"math/big"

	"github.com/soniakeys/integer/swing"
	"github.com/soniakeys/integer/xmath"
)

// DoubleFactorial computes the double factorial function (!!), OEIS A001147.
//
// It computes the result in parameter z, replacing the existing value of z,
// and returning z.
//
// It is uses effecient algorithm based prime numbers.
func DoubleFactorial(z *big.Int, n uint) *big.Int {
	return DoubleFactorialS(z, swing.New(n+1), n)
}

// DoubleFactorialS computes the double factorial given a swing.Swing object.
//
// It allows multiple double factorials to be computed after computing
// prime numbers just once.  The swing.Swing object encapsulates a prime
// number sieve.
//
// For computing double factorials up to n, the swing.Swing object should be
// constructed with at least n (if n is even) or n+1 (if n is odd).
//
// DoubleFactorialS returns nil if the sieve is not big enough.
func DoubleFactorialS(z *big.Int, p *swing.Swing, n uint) *big.Int {
	nEven := n&1 == 0
	if n < uint(len(smallOddDoubleFactorial)) {
		z.SetInt64(smallOddDoubleFactorial[n])
	} else {
		var nn uint
		if nEven {
			nn = n / 2
		} else {
			nn = n + 1
		}

		if uint64(nn) > p.Sieve.Lim && nn > uint(len(swing.SmallOddFactorial)) {
			return nil
		}

		var os big.Int
		var oddDoubleFactorial func(uint)
		oddDoubleFactorial = func(nn uint) {
			if nn < uint(len(swing.SmallOddFactorial)) {
				z.SetInt64(swing.SmallOddFactorial[nn])
				return
			}
			oddDoubleFactorial(nn / 2)
			if nn < n {
				z.Mul(z, z)
			}
			z.Mul(z, p.OddSwing(&os, nn))
		}
		oddDoubleFactorial(nn)
	}
	if nEven {
		z.Lsh(z, n-xmath.BitCount32(uint32(n)))
	}
	return z
}

var smallOddDoubleFactorial []int64 = []int64{1, 1, 1, 3, 1,
	15, 3, 105, 3, 945, 15, 10395, 45, 135135, 315, 2027025, 315,
	34459425, 2835, 654729075, 14175, 13749310575, 155925,
	316234143225, 467775, 7905853580625, 6081075, 213458046676875,
	42567525, 6190283353629375, 638512875, 191898783962510625,
	638512875, 6332659870762850625, 10854718875,
}
