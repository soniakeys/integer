// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package xmath contains functions useful for integer calculations.
package xmath

import "math/big"

// ProductSerialTreshold is the recursion cutoff.
//
// It can be tuned for best performance.
var ProductSerialThreshold = 40

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

// FloorSqrt32 is an integer square root function.
func FloorSqrt32(n uint32) uint32 {
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

// FloorSqrt is an integer square root function.
func FloorSqrt64(n uint64) uint64 {
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

// TrailingZeros returns the number of trailing 0 bits in v.
//
// If v is 0, it returns 0.
func TrailingZeros(v uint) (c byte) {
	// seqential algorithm
	if v != 0 {
		for v&1 == 0 {
			v >>= 1
			c++
		}
	}
	return
}

// reference: http://graphics.stanford.edu/~seander/bithacks.html
const deBruijn32Multiple = 0x077CB531
const deBruijn32Shift = 27

var deBruijn32Bits = []byte{
	0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
}

// TrailingZeros32 returns the number of trailing 0 bits in v.
//
// If v is 0, it returns 0.
func TrailingZeros32(v uint32) byte {
	return deBruijn32Bits[v&-v*deBruijn32Multiple>>deBruijn32Shift]
}

const deBruijn64Multiple = 0x03f79d71b4ca8b09
const deBruijn64Shift = 58

var deBruijn64Bits = []byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

// TrailingZeros64 returns the number of trailing 0 bits in v.
//
// If v is 0, it returns 0.
func TrailingZeros64(v uint64) byte {
	return deBruijn64Bits[v&-v*deBruijn64Multiple>>deBruijn64Shift]
}

// works for 32 and 64 anyway
const wordBits = int((^big.Word(0))>>32&1+1) * 32

var tzw func(big.Word) int

func init() {
	switch wordBits {
	case 32:
		tzw = func(w big.Word) int { return int(TrailingZeros32(uint32(w))) }
	case 64:
		tzw = func(w big.Word) int { return int(TrailingZeros64(uint64(w))) }
	}
}

// TrailingZerosBig returns the number of trailing 0 bits in v.
//
// If v is 0, it returns 0.
func TrailingZerosBig(v *big.Int) int {
	for i, b := range v.Bits() {
		if b != 0 {
			return i*wordBits + tzw(b)
		}
	}
	return 0
}

// TrailingOnesBig returns the number of trailing 1 bits in v.
func TrailingOnesBig(v *big.Int) int {
	words := v.Bits()
	for i, b := range words {
		if b != ^big.Word(0) {
			return i*wordBits + tzw(^b)
		}
	}
	return len(words) * wordBits
}
