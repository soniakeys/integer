// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT
// Algorithm credit, Melissa E. O’Neill.  Inspiration, Anh Hai Trinh.
// Heap code lifted from Go package container/heap, subject to following:
/*
Copyright (c) 2012 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Package queue implements the prime.Generator interface with a priority queue
// based sequential algorithm.  In comparison with a sieve algorithm, it has
// no up-front computation overhead, but time and space requirements will
// increase as primes are generated.
package queue

import (
	"math"

	"github.com/soniakeys/integer/prime"
)

// PQueue has no state.  Memory is only used when primes are requested.
type PQueue struct{}

// Limit is a method of the prime.Generator interface.
// Semantics are per prime.Generator documentation.
func (s PQueue) Limit() uint64 {
	return math.MaxUint64
}

type pMult struct {
	prime uint64
	pMult uint64
}

type wheel func() uint64

func makeWheel(half []uint64) wheel {
	last := len(half) - 1
	var forwards bool
	pos := 1
	return func() uint64 {
		if forwards {
			if pos < last {
				pos++
			} else {
				forwards = false
				pos--
			}
		} else {
			if pos > 0 {
				pos--
			} else {
				forwards = true
				pos++
			}
		}
		return half[pos]
	}
}

// Iterate is a method of the prime.Generator interface.
// Semantics are per prime.Generator documentation.
func (s PQueue) Iterate(min, max uint64, visitor prime.Visitor) (ok bool) {
	if min < 2 {
		min = 2
	}
	if min > max {
		return true
	}
	if min == 2 && max >= 2 && visitor(2) {
		return true
	}
	if min <= 3 && max >= 3 && visitor(3) {
		return true
	}

	// estimate number of primes < sqrt(max)
	// this will determine the storage required by the wheel and the heap.
	var piSM int
	sqrtMaxF := math.Sqrt(float64(max))
	sqrtMaxI := uint64(sqrtMaxF)
	if max == math.MaxInt64 {
		// this is the case of unbounded iteration.  so really, let's
		// limit this to something reasonable and reallocate as needed.
		// a limit of 1000 limits the inital heap to 16k of memory
		piSM = 1000
	} else {
		// a quick estimate for pi(sqrt(max))
		ln := math.Log(sqrtMaxF)
		piSM = int(sqrtMaxF / ln * (1 + 1.2762/ln))
	}

	// build the wheel
	half := []uint64{2, 4}
	wheelPrimes := 2
	var prime uint64
	for {
		prime = half[1] + 1
		if prime > max {
			return true
		}
		if prime >= min && visitor(prime) {
			return true
		}
		break
	}
	wheel := makeWheel(half)
	wheel()
	wheel()

	pq := make([]pMult, 1, piSM-wheelPrimes)
	pq[0].prime = prime
	pq[0].pMult = prime * prime

	for k := prime + wheel(); k <= max; {
		if k < pq[0].pMult {
			if k >= min {
				if visitor(k) {
					return true
				}
			}
			if k <= sqrtMaxI {
				n := len(pq)
				if cap(pq) == n {
					bq := make([]pMult, n+1, n*2)
					copy(bq, pq)
					pq = bq
				} else {
					pq = pq[0 : n+1]
				}
				pq[n].prime = k
				pq[n].pMult = k * k
			}
			k += wheel()
		} else {
			for {
				if k == pq[0].pMult {
					k += wheel()
				}
				pq[0].pMult += pq[0].prime
				// sift down
				for i := 0; ; {
					j1 := 2*i + 1
					if j1 >= len(pq) {
						break
					}
					j := j1 // left child
					j2 := j1 + 1
					if j2 < len(pq) && pq[j1].pMult >= pq[j2].pMult {
						j = j2 // = 2*i + 2  // right child
					}
					if pq[i].pMult < pq[j].pMult {
						break
					}
					pq[i], pq[j] = pq[j], pq[i]
					i = j
				}
				if pq[0].pMult > k {
					break
				}
			}
		}
	}
	return true
}
