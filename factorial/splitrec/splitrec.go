// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Factorial by split-recursive algorithm.
package splitrec

import (
	"math/big"

	"github.com/soniakeys/integer/xmath"
)

// Factorial computes n!, leaving result in z and returning z.
func Factorial(z *big.Int, n uint) *big.Int {
	currentN := int64(1)
	var product func(*big.Int, uint) *big.Int
	product = func(z *big.Int, n uint) *big.Int {
		switch n {
		case 1:
			currentN += 2
			return z.SetInt64(currentN)
		case 2:
			currentN += 2
			r := currentN
			currentN += 2
			r *= currentN
			return z.SetInt64(r)
		}
		m := n / 2
		var r big.Int
		return z.Mul(product(z, m), product(&r, n-m))
	}

	var p, pr big.Int
	p.SetInt64(1)
	z.SetInt64(1)

	var h, shift uint
	var high uint = 1
	log2n := xmath.Log2(n)

	for h != n {
		shift += h
		h = n >> log2n
		log2n--
		length := high
		high = (h - 1) | 1
		length = (high - length) / 2

		if length > 0 {
			z.Mul(z, p.Mul(&p, product(&pr, length)))
		}
	}

	return z.Lsh(z, shift)
}
