// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package sprp implements a Miller-Rabin deterministic strong probable-prime
// (SPRP) test.
//
// The current implementation has a limit of n = max uint32.
// In comparison to sieve and priority queue algorithms, SPRP has no up-front
// compute overhead and no space requirements, but will be ulitimately slower
// than other algorithms if a large number of primes is requested.
package sprp

import "github.com/soniakeys/integer/prime"

// SPRP doesn't actually require any state.
// Values in this struct represent a cache of last values used.
type SPRP struct {
	bx       int          // index into baseSets, index of set currently cached.
	limit    uint32       // limit of cached base set.
	bases    []uint32     // bases of cached set
	basePow2 [][31]uint64 // for each cached base, maxP2 powers.
	maxP2    []int        // number of powers of 2 currently computed.
}

// reference: http://miller-rabin.appspot.com/
var baseSets = []struct {
	limit uint32
	bases []uint32
}{
	{5329, []uint32{377687}},
	{316349281, []uint32{11000544, 31481107}},
	//	{4759123140, []uint32{2, 7, 61}}, (it's a little over 2^32)
	{1<<32 - 1, []uint32{2, 7, 61}},
}

// New constructs an SPRP object and initializes the cache to a valid state.
func New() *SPRP {
	var m SPRP
	wc := len(baseSets[len(baseSets)-1].bases)
	m.maxP2 = make([]int, wc)
	m.basePow2 = make([][31]uint64, wc)
	m.resetLimit(3)
	return &m
}

func (m *SPRP) resetLimit(n uint32) {
	m.bx = 0
	m.limit = baseSets[0].limit
	for n > m.limit {
		m.bx++
		m.limit = baseSets[m.bx].limit
	}
	m.bases = baseSets[m.bx].bases
	for i, a := range m.bases {
		m.basePow2[i][0] = uint64(a)
		m.maxP2[i] = 0
	}
}

// Limit satisfies prime.Generator.
// The current implementation has a fixed limit of max uint32.
func (m *SPRP) Limit() uint64 {
	return 1<<32 - 1
}

// Iterate satisfies prime.Generator.
func (m *SPRP) Iterate(min, max uint64, visitor prime.Visitor) bool {
	switch {
	case max > m.Limit():
		return false
	case max < 2:
		return true
	case min <= 2:
		min = 2
		if visitor(2) {
			return true
		}
	}
	c := uint32(min | 1)
	m.resetLimit(c)
	for max32 := uint32(max); c <= max32; c += 2 {
		if m.Prime(c) && visitor(uint64(c)) {
			return true
		}
	}
	return true
}

// Prime returns true if n is prime.
func (m *SPRP) Prime(n uint32) bool {
	for n > m.limit {
		m.bx++
		m.limit = baseSets[m.bx].limit
	}
	m.bases = baseSets[m.bx].bases
	for i, a := range m.bases {
		m.maxP2[i] = 0
		m.basePow2[i][0] = uint64(a)
	}

	nm1 := n - 1
	s := deBruijn32Zeros(nm1)
	d := nm1 >> s
	for i, _ := range m.bases {
		// compute x := a^d % n
		n64 := uint64(n)
		x := uint64(1)
		p2 := m.basePow2[i][0:]
		bit := 0
		for dr := d; dr > 0; dr >>= 1 {
			if bit > m.maxP2[i] {
				p1 := p2[bit-1]
				p2[bit] = p1 * p1 % n64
				m.maxP2[i]++
			}
			if dr&1 != 0 {
				x = x * p2[bit] % n64
			}
			bit++
		}
		if x == 1 || uint32(x) == nm1 {
			continue
		}
		for r := byte(1); ; r++ {
			if r == s {
				return false
			}
			x = x * x % n64
			if x == 1 {
				return false
			}
			if uint32(x) == nm1 {
				break
			}
		}
	}
	return true
}

// reference: http://graphics.stanford.edu/~seander/bithacks.html
const deBruijn32Multiple = 0x077CB531
const deBruijn32Shift = 27

var deBruijn32Bits = []byte{
	0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
}

func deBruijn32Zeros(v uint32) byte {
	return deBruijn32Bits[v&-v*deBruijn32Multiple>>deBruijn32Shift]
}
