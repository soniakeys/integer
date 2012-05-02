// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package sieve implements a prime number sieve.
//
// The fundamental algorithm is by Eratosthenes (276-194 BC).  Representation
// is with bits, and factors of 2 and 3 are excluded with a wheel method.
package sieve

// word size dependent constants
const (
	bitsPerInt = 32
	mask       = bitsPerInt - 1
	log2Int    = 5
)

// Sieve type holds a completed sieve.
type Sieve struct {
	Len         uint64
	isComposite []uint32
}

var smallComposites = []uint32{1762821248, 848611808, 3299549660, 2510511646}
var smallCompositeLimit uint64 = 3 * bitsPerInt * uint64(len(smallComposites))

// New is the Sieve constructor, completing the sieve operation.
func New(n uint64) (ps *Sieve) {
	// Prime number sieve, Eratosthenes (276-194 b.t. )
	// This implementation considers only multiples of primes
	// greater than 3, so the smallest value has to be mapped to 5.
	// Note: There is no multiplication operation in this function
	// and *no call to a sqrt* function.

	ps = new(Sieve)
	ps.Len = n

	if n <= smallCompositeLimit {
		ps.isComposite = smallComposites
		return
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
	return
}

// Visitor function passed to Iterate method of Sieve.
//
// A visitor function should return false to continue iteration.
// Iterate interprets a true result from the visitor function as a request
// to terminate iteration.
type Visitor func(prime uint64) (terminate bool)

// Iterator iterates over primes betwing min and max inclusive, and calls
// the visitor function for each prime.
//
// Iterate returns false if max > sieve size, otherwise it returns true.
// It returns true even if no primes happen to be between the specified
// bounds or if the visitor function terminates iteration early.
func (ps *Sieve) Iterate(min, max uint64, visitor Visitor) bool {
	// isComposite[0] ... isComposite[n] includes
	// 5 <= primes numbers <= 96*(n+1)+1

	switch {
	case max > ps.Len:
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
