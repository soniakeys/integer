package prime_test

import (
	"testing"

	"github.com/soniakeys/integer/prime"
	"github.com/soniakeys/integer/prime/queue"
	"github.com/soniakeys/integer/prime/sieve"
	"github.com/soniakeys/integer/prime/sprp"
)

// Test the two high level access functions in prime.go,
// Primes() and Iterator().  Also test that the generators agree
// on all primes under some fairly bit limit.

func TestPrime(t *testing.T) {
	// test each implemented generator
	t100("Sieve", sieve.New(100), t)
	t100("SPRP", sprp.New(), t)
	t100("PQueue", &queue.PQueue{}, t)

	const bigtest = 1e3
	// use sieve as the reference
	reference := prime.Primes(sieve.New(bigtest))
	// check that queue agrees
	var i int
	queue.PQueue{}.Iterate(0, bigtest, func(p uint64) bool {
		if p != reference[i] {
			t.Errorf("%dth prime?  sieve says %d, queue says %d\n",
				i, reference[i], p)
		}
		i++
		return false
	})
	// check that sprp agrees
	i = 0
	sprp.New().Iterate(0, bigtest, func(p uint64) bool {
		if p != reference[i] {
			t.Errorf("%dth prime?  23 says %d, pq says %d\n",
				i, reference[i], p)
		}
		i++
		return false
	})
}

func t100(p string, pg prime.Generator, t *testing.T) {
	k100 := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
		53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	// test Primes()
	pa := prime.Primes(pg, 0, 100)

	if len(pa) != len(k100) {
		t.Errorf("%s.Primes() gives wrong number of primes under 100.\n", p)
		t.Log(pa)
		return
	}

	for i, prime := range pa {
		if prime != k100[i] {
			t.Errorf("%s.Primes() gives %d for %dth prime.  %d expected.\n",
				p, prime, i+1, k100[i])
			return
		}
	}

	// test Iterator()
	var i int
	for prime := range prime.Iterator(pg, 0, 100) {
		if prime != k100[i] {
			t.Errorf("%s.Iterator() gives %d for %dth prime.  %d expected.\n",
				p, prime, i+1, k100[i])
			return
		}
		i++
	}

	if i != len(k100) {
		t.Errorf("%s.Iterator() returns wrong number of primes under 100.\n", p)
		return
	}
}
