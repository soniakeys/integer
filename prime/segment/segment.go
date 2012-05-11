// Copyright 2012 Peter Luschny, Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package segment computes prime numbers with a parallel segmented sieve.
package segment

import (
	"math"
	"runtime"

	"github.com/soniakeys/integer/prime"
)

type Sieve struct {
	Lim         uint64
	isComposite []uint64
}

func (ps *Sieve) Limit() uint64 {
	return ps.Lim
}

const (
	bitsPerWord = 64 // word size of isComposite
	log2Int     = 6

	mask = bitsPerWord - 1

	density = 3 // integers represented per bit
	wordCap = bitsPerWord * density

	smallCompositeLimit = 192 // wordCap * len(smallComposites)
)

var smallComposites []uint64 = []uint64{0x3294c9e069128480}

// New is the Sieve constructor, completing the sieve operation.
func New(n uint64) *Sieve {
	return new(Sieve).Init(n)
}

// Iterate iterates over primes betwing min and max inclusive, and calls
// the visitor function for each prime.
//
// Iterate returns false if max > sieve size, otherwise it returns true.
// It returns true even if no primes happen to be between the specified
// bounds or if the visitor function terminates iteration early.
func (s *Sieve) Iterate(min, max uint64, visitor prime.Visitor) bool {
	// isComposite[0] ... isComposite[n] includes
	// 5 <= primes numbers <= 96*(n+1)+1

	switch {
	case max > s.Lim:
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

	absPos := (min+(min+1)%2)/density - 1
	index := absPos / bitsPerWord
	bitPos := absPos % bitsPerWord
	prime := 5 + density*(bitsPerWord*index+bitPos) - bitPos&1
	inc := bitPos&1*2 + 2

	for prime <= max {
		bitField := s.isComposite[index] >> uint64(bitPos)
		index++
		for ; bitPos < bitsPerWord; bitPos++ {
			if bitField&1 == 0 && visitor(prime) {
				return true
			}
			prime += inc
			if prime > max {
				return true
			}
			inc = 6 - inc
			bitField >>= 1
		}
		bitPos = 0
	}
	return true
}

// Prime number sieve, Eratosthenes (276-194 b.c.)
// Adapted from code by Peter Luschny.  Luschny algorithm implements
// 2,3 wheel logic, bit representation, and precomputed small primes.
// Go coding and parallelization by Sonia Keys
//
// sieve initializes the Sieve object by allocating memory and
// running the Luschny sieve algorithm.  This will find prime numbers
// less than or equal to the parameter n.
// If n <= 0, the passed object is set to the zero object
func (ps *Sieve) Init(n uint64) *Sieve {
	if n <= 0 {
		*ps = Sieve{}
		return ps
	}
	ps.Lim = n

	if n <= smallCompositeLimit {
		ps.isComposite = smallComposites
		return ps
	}

	ps.isComposite = make([]uint64, (n+wordCap-1)/wordCap)

	// it would be nice to query for L2 cache size.
	const l2cacheSize = 4e6

	// leave some cache for other purposes.
	const l2quota = l2cacheSize / 2

	// determine bound for single threaded computation
	var maxB, wordsq uint64
	all := uint64(n) / density
	if n < l2quota*density*8 {
		// if sieve fits in L2, no need for segmentation,
		// bound is entire range
		maxB = all
	} else {
		// run single threaded only to square root of n
		wordsq = uint64(math.Ceil(math.Sqrt(float64(len(ps.isComposite)))))
		maxB = wordsq * bitsPerWord
	}

	var d1, d2, p1, p2, s, s2 uint64 = 8, 8, 3, 7, 7, 3
	var toggle bool

	for bx := uint64(0); s < maxB; bx++ {
		if (ps.isComposite[bx>>log2Int] & (1 << (bx & mask))) == 0 {
			inc := p1 + p2
			for c := s; c < maxB; c += inc {
				ps.isComposite[c>>log2Int] |= 1 << (c & mask)
			}
			for c := s + s2; c < maxB; c += inc {
				ps.isComposite[c>>log2Int] |= 1 << (c & mask)
			}
		}

		if toggle {
			toggle = false
			s += d1
			d2 += 8
			p1 += 2
			p2 += 6
			s2 = p1
		} else {
			toggle = true
			s += d2
			d1 += 16
			p1 += 2
			p2 += 2
			s2 = p2
		}
	}

	if maxB == all {
		return ps
	}

	// parallelize the rest
	wordsNonSq := uint64(len(ps.isComposite)) - wordsq
	nCpu := runtime.GOMAXPROCS(0)
	segQuota := uint64(l2quota / nCpu)
	bytesToSegment := wordsNonSq * bitsPerWord / 8
	segments := (bytesToSegment + segQuota - 1) / segQuota
	wordsPerSegment := (wordsNonSq + segments - 1) / segments

	type segCS struct {
		start, end uint64
	}
	segCh := make(chan *segCS)
	doneCh := make(chan int)

	// dispatcher
	go func() {
		// start the workers
		for i := 0; i < nCpu; i++ {
			// workers
			go func() {
				for {
					seg := <-segCh
					if seg == nil {
						return
					}
					minB := bitsPerWord * seg.start
					maxB := bitsPerWord * seg.end

					var d1, d2, p1, p2, s, s2 uint64 = 8, 8, 3, 7, 7, 3
					var toggle bool

					for bx := uint(0); s < maxB; bx++ {
						if (ps.isComposite[bx>>log2Int] & (1 << (bx & mask))) == 0 {
							inc := p1 + p2

							c := s
							if c < minB {
								c = minB + inc - 1 - (minB-c-1)%inc
							}
							for ; c < maxB; c += inc {
								ps.isComposite[c>>log2Int] |= 1 << (c & mask)
							}
							c = s + s2
							if c < minB {
								c = minB + inc - 1 - (minB-c-1)%inc
							}
							for ; c < maxB; c += inc {
								ps.isComposite[c>>log2Int] |= 1 << (c & mask)
							}
						}

						if toggle {
							toggle = false
							s += d1
							d2 += 8
							p1 += 2
							p2 += 6
							s2 = p1
						} else {
							toggle = true
							s += d2
							d1 += 16
							p1 += 2
							p2 += 2
							s2 = p2
						}
					}
					doneCh <- 0
				}
			}()
		}

		// dispatch
		start, end := wordsq, wordsq+wordsPerSegment
		for i := uint64(1); i < segments; i++ {
			segCh <- &segCS{start, end}
			start, end = end, end+wordsPerSegment
		}
		segCh <- &segCS{start, uint64(len(ps.isComposite))}
	}()

	// count completions
	for i := uint64(0); i < segments; i++ {
		<-doneCh
	}
	close(segCh)
	return ps
}
