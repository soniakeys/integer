// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// xmath contains functions useful for integer calculations.
package xmath

import "math/big"

// ProductSerialTreshold is the recursion cutoff.
//
// It can be tuned for best performance.
var ProductSerialThreshold = 24

// Product computes product of numbers in seq.
//
// It recursively partions large lists to reduce the number of multiplies.
// It computes the result in z to avoid unnecessary allocation.
// It returns z for convenience in chaining expressions.
func Product(z *big.Int, seq []uint64) *big.Int {
	if len(seq) <= ProductSerialThreshold {
		if len(seq) == 0 {
			return z.SetInt64(1)
		}
		var b big.Int
		z.SetInt64(int64(seq[0]))
		for _, s := range seq[1:] {
			z.Mul(z, b.SetInt64(int64(s)))
		}
		return z
	}
	halfLen := len(seq) / 2
	lprod := Product(z, seq[0:halfLen])
	rprod := Product(new(big.Int), seq[halfLen:])
	return z.Mul(lprod, rprod)
}

// BitCount32 returns the number of 1s in a uint32.
func BitCount32(w uint32) uint {
    const (
        ff    = 1<<32 - 1
        mask1 = ff / 3
        mask3 = ff / 5
        maskf = ff / 17
        maskp = ff / 255
    )
    w -= w >> 1 & mask1
    w = w&mask3 + w>>2&mask3
    w = (w + w>>4) & maskf
    return uint(w * maskp >> 24)
}

// BitCount64 returns the number of 1s in a uint64.
func BitCount64(w uint64) uint {
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

// Log2 returns the integer log base 2 of x.
//
// If x is 0, the result is the maximum value of a uint.
func Log2(x uint) (n uint) {
	for ; x > 0xff; x >>= 8 {
		n += 8
	}
	for ; x > 0; x >>= 1 {
		n++
	}
	return n - 1
}

// FloorSqrt is an integer square root function.
func FloorSqrt(n uint) uint {
	b := (n + 1) / 2
	if b >= n {
		return n
	}
	a := b
	for {
		b = (n/a + a) / 2
		if b >= a {
			break
		}
		a = b
	}
	return a
}
