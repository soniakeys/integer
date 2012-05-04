// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package sieve implements a prime number sieve.
//
// The fundamental algorithm is by Eratosthenes (276-194 BC).  Representation
// is with bits, and factors of 2 and 3 are excluded with a wheel method.
// Compared to SPRP and priority queue methods, this is relatively speedy,
// but does have a space and time overhead before any primes can be returned.
package sieve

import (
	"math"

	"github.com/soniakeys/integer/prime"
	"github.com/soniakeys/integer/xmath"
)

// word size dependent constants
const (
	bitsPerInt = 32
	mask       = bitsPerInt - 1
	log2Int    = 5
)

// Sieve type holds a completed sieve.
type Sieve struct {
	Lim         uint64
	isComposite []uint32
}

func (ps *Sieve) Limit() uint64 {
	return ps.Lim
}

var smallComposites = []uint32{1762821248, 848611808, 3299549660, 2510511646}
var smallCompositeLimit uint64 = 3 * bitsPerInt * uint64(len(smallComposites))
var smallPiLimit uint64

func init() {
	smallPiLimit = 2
	for _, w := range smallComposites {
		smallPiLimit += uint64(32 - xmath.BitCount32(w))
	}
}

// New is the Sieve constructor, completing the sieve operation.
func New(n uint64) *Sieve {
	return new(Sieve).Init(n)
}

// Init initializes the Sieve object by allocating memory and running
// the sieve algorithm.  This will find prime numbers less than
// or equal to the parameter n.
//
// If n <= 0, the passed object is set to the zero object.
//
// The function returns its reciever.
func (ps *Sieve) Init(n uint64) *Sieve {
	// Prime number sieve, Eratosthenes (276-194 b.t. )
	// This implementation considers only multiples of primes
	// greater than 3, so the smallest value has to be mapped to 5.
	// Note: There is no multiplication operation in this function
	// and *no call to a sqrt* function.

	ps.Lim = n

	if n <= smallCompositeLimit {
		ps.isComposite = smallComposites
		return ps
	}

	ps.isComposite = make([]uint32, n/(3*bitsPerInt)+1)
	var (
		d1, d2, p1, p2, s, s2 uint64 = 8, 8, 3, 7, 7, 3
		l, c, max, inc        uint64 = 0, 1, n / 3, 0
		toggle                bool
	)

	for s < max { // --  scan the sieve
		// --  if a primes is found ...
		if (ps.isComposite[l>>log2Int] & (1 << (l & mask))) == 0 {
			inc = p1 + p2 // --  ... cancel its multiples

			// --  ... set c as composite
			for c = s; c < max; c += inc {
				ps.isComposite[c>>log2Int] |= 1 << (c & mask)
			}

			for c = s + s2; c < max; c += inc {
				ps.isComposite[c>>log2Int] |= 1 << (c & mask)
			}
		}

		l++
		toggle = !toggle
		if toggle {
			s += d2
			d1 += 16
			p1 += 2
			p2 += 2
			s2 = p2
		} else {
			s += d1
			d2 += 8
			p1 += 2
			p2 += 6
			s2 = p1
		}
	}
	return ps
}

// Iterate iterates over primes betwing min and max inclusive, and calls
// the visitor function for each prime.
//
// Iterate returns false if max > sieve size, otherwise it returns true.
// It returns true even if no primes happen to be between the specified
// bounds or if the visitor function terminates iteration early.
func (ps *Sieve) Iterate(min, max uint64, visitor prime.Visitor) bool {
	// isComposite[0] ... isComposite[n] includes
	// 5 <= primes numbers <= 96*(n+1)+1

	switch {
	case max > ps.Lim:
		return false
	case max < 2:
		return true
	case min <= 2:
		min = 2
		if visitor(2) {
			return true
		}
	}
	switch {
	case max < 3:
		return true
	case min <= 3:
		if visitor(3) {
			return true
		}
	}

	absPos := uint64((min+(min+1)%2)/3 - 1)
	index := absPos / bitsPerInt
	bitPos := absPos % bitsPerInt
	toggle := bitPos&1 == 1
	prime := uint64(5 + 3*(bitsPerInt*index+bitPos) - bitPos&1)

	for prime <= max {
		bitField := ps.isComposite[index] >> bitPos
		index++
		for ; bitPos < bitsPerInt; bitPos++ {
			if bitField&1 == 0 && visitor(prime) {
				return true
			}
			toggle = !toggle
			if toggle {
				prime += 2
			} else {
				prime += 4
			}
			if prime > max {
				return true
			}
			bitField >>= 1
		}
		bitPos = 0
	}
	return true
}

// InitPi similar to Init, but parameter is a minimum number of
// prime numbers to find rather than a maximum value of primes.
//    
// Mathematically, π(n), is the prime counting function, the number of primes
// less than or equal to n.  InitPi selects a bound for n, such that π(n) ≥ pn.
func (ps *Sieve) InitPi(pn uint64) {
	var n uint64
	if pn <= smallPiLimit {
		n = smallCompositeLimit
	} else {
		ln := math.Log(float64(pn))
		n = pn * uint64(ln+math.Log(ln))
	}
	ps.Init(n)
}
