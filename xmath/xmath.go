// Adapted with permission from code by Peter Luschny
// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package xmath

import "math/big"

const productSerialThreshold = 24

func Product(seq []uint64) *big.Int {
	if len(seq) <= productSerialThreshold {
		var b big.Int
		sprod := big.NewInt(int64(seq[0]))
		for _, s := range seq[1:] {
			b.SetInt64(int64(s))
			sprod.Mul(sprod, &b)
		}
		return sprod
	}
	halfLen := len(seq) / 2
	lprod := Product(seq[0:halfLen])
	rprod := Product(seq[halfLen:])
	return lprod.Mul(lprod, rprod)
}

func BitCount(w uint64) uint {
	const (
		ff    = 1<<64 - 1
		mask1 = ff / 3
		mask3 = ff / 5
		maskf = ff / 17
		maskp = maskf >> 3 & maskf
	)
	w -= w >> 1 & mask1
	w = w&mask3 + w>>2&mask3
	w = (w + w>>4) & maskf
	return uint(w * maskp >> 56)
}

func FloorSqrt(n uint) uint {
	for b := n; ; {
		a := b
		b = (n/a + a) / 2
		if b >= a {
			return a
		}
	}
	panic("unreachable")
}
