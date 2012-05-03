// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package prime computes factorials from swinging factorials, computed
// from prime numbers.
package prime

import (
	"math/big"

	"github.com/soniakeys/integer/swing"
	"github.com/soniakeys/integer/xmath"
)

// Factorial computes n!, leaving the result in z.
func Factorial(z *big.Int, n uint) *big.Int {
	return FactorialS(z, swing.New(n), n)
}

// FactorialS computes n! given an existing swing.Swing object.
//
// It allows multiple factorials to be computed after computing prime numbers
// just once.  The swing.Swing object encapsulates a prime number sieve.
//
// For computing factorials up to n, the swing.Swing object should be
// constructed with with the same or greater n.
//
// FactorialS returns nil if the sieve is not big enough.
func FactorialS(z *big.Int, ps *swing.Swing, n uint) *big.Int {
	var oddFactorial func(*big.Int, uint) *big.Int
	oddFactorial = func(z *big.Int, n uint) *big.Int {
		if n < uint(len(swing.SmallOddFactorial)) {
			return z.SetInt64(swing.SmallOddFactorial[n])
		}

		oddFactorial(z, n/2)
		var os big.Int
		return z.Mul(z.Mul(z, z), ps.OddSwing(&os, n))
	}
	return z.Lsh(oddFactorial(z, n), n-xmath.BitCount32(uint32(n)))
}
