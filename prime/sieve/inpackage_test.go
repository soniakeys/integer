package sieve

import (
	"testing"

	"github.com/soniakeys/integer/prime"
)

// Validate that the package variable smallComposites can be reproduced
// with the sieve algorithm.
func TestSmallComposites(t *testing.T) {
	sc := smallComposites
	smallComposites = nil
	scl := smallCompositeLimit
	smallCompositeLimit = 0
	s := New(scl)
	if len(s.isComposite) < len(sc) {
		goto bad
	}
	for i, b := range sc {
		if b != s.isComposite[i] {
			goto bad
		}
	}
	// put things back the way they were for other tests
	smallComposites = sc
	smallCompositeLimit = scl
	return // ok
bad:
	t.Log("bad small composite list.")
	t.Log("valid:", s.isComposite)
	t.Log("found:", sc)
	t.Fail()
}

// Some more tests.
// Nothing here should invoke the Eratosthenes algorithm.
func TestSmall(t *testing.T) {
	// reproduce 10 primes hardcoded right here.
	ten := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	n := ten[9]
	s := New(n)
	if s.Limit() < n {
		t.Fatal("Limit of", n, "expected.  Got", s.Limit())
	}
	// check that the 10 are right
	var i int
	s.Iterate(0, n, func(prime uint64) (termiate bool) {
		if prime != ten[i] {
			t.Fatal("Incorrect prime.  Expected", ten[i], "got", prime)
		}
		i++
		return
	})

	// resieve, getting all precomputed primes
	n = 3 * bitsPerInt * uint64(len(smallComposites))
	s.Init(n)

	// s.isComposites shouldn't have grown
	if len(s.isComposite) != len(smallComposites) {
		t.Error("Sieve didn't recognize small n")
	}

	// use handy PrimeGenerator function to save primes for next test
	smallPrimes := prime.Primes(s)

	// test InitPi on boundary case of small np
	s.InitPi(uint64(len(smallPrimes)))
	if s.Lim != n {
		t.Error("SievePi didn't seem to work on small np")
	}

	// asking for one more prime than was precomputed should trigger sieve.
	s.InitPi(1 + uint64(len(smallPrimes)))
	if len(s.isComposite) <= len(smallComposites) {
		t.Fatal("sieve didn't seem to work", len(s.isComposite), len(smallComposites))
	}

	// check that primes from sieve match precomputed primes from
	// previous test.  test IterateFunc early termination feature
	// in the process.
	var j uint64
	ok := s.Iterate(0, s.Lim, func(prime uint64) bool {
		if int(j) == len(smallPrimes) {
			return true
		}
		if prime != smallPrimes[j] {
			t.Fatalf("sieve failed.  expected %d for prime[%d].  got %d.\n",
				smallPrimes[j], j, prime)
		}
		j++
		return false
	})
	if !ok {
		t.Error("Iterate returned false on early termination.")
	}
}
