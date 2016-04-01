package prime_test

import (
	"reflect"
	"testing"

	"github.com/soniakeys/integer/prime"
	"github.com/soniakeys/integer/prime/queue"
	"github.com/soniakeys/integer/prime/segment"
	"github.com/soniakeys/integer/prime/sieve"
	"github.com/soniakeys/integer/prime/sprp"
)

// Exercise Limit and Iterate methods of each implemented generator.
func TestMethods(t *testing.T) {
	t100(t, segment.New(100))
	t100(t, sieve.New(100))
	t100(t, queue.PQueue{})
	t100(t, sprp.New())
}

func t100(t *testing.T, pg prime.Generator) {
	if lim := pg.Limit(); lim < 100 {
		t.Errorf("%s Limit() returned %d.  result >= 100 expected.",
			reflect.TypeOf(pg), lim)
	}
	// Test boundary conditions and test results against known primes.
	k100 := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
		53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	ti(t, pg, 0, 100, k100) // Iterate should handle 0 and 1 as min
	ti(t, pg, 1, 100, k100)
	ti(t, pg, 2, 100, k100)     // min = first prime, but 2 is often special
	ti(t, pg, 3, 100, k100[1:]) // min = first odd prime, still often special
	ti(t, pg, 4, 97, k100[2:])  // min = first composite, max is prime
	ti(t, pg, 5, 97, k100[2:])  // min = typical prime
}

func ti(t *testing.T, pg prime.Generator, min, max uint64, known []uint64) {
	var i int
	pg.Iterate(min, max, func(prime uint64) (terminate bool) {
		if i == len(known) {
			t.Errorf("%s.Iterate(%d, %d) returning too many primes.",
				reflect.TypeOf(pg), min, max)
			return true
		}
		if prime != known[i] {
			t.Errorf("%s.Iterate(%d, %d) = %d for %dth prime.  %d expected.",
				reflect.TypeOf(pg), min, max, prime, i+1, known[i])
			return true
		}
		i++
		return
	})
	if i != len(known) {
		t.Errorf("%s.Iterate(%d, %d) didn't find all primes.",
			reflect.TypeOf(pg), min, max)
	}
}

// Test the two generic access functions in prime.go, Primes() and
// Iterator().  Also test that the generators agree on all primes under
// a somewhat bigger limit.
//
// Use basic sieve as a reference, compare it to other generators.
func TestGeneric(t *testing.T) {
	const limit = 1000
	// exercise generic function Primes on sieve.
	reference := prime.Primes(sieve.New(limit), 0, limit)
gloop:
	for _, gen := range []prime.Generator{
		segment.New(limit),
		queue.PQueue{},
		sprp.New(),
	} {
		// exercise generic function Iterator on other generators.
		ch := prime.Iterator(gen, 0, limit)
		for i, pRef := range reference {
			switch p, ok := <-ch; {
			case !ok:
				t.Errorf("%s.Iterator(0, %d) didn't find all primes.",
					reflect.TypeOf(gen), limit)
				continue gloop
			case p != pRef:
				t.Errorf("%dth prime?  sieve says %d, %s says %d\n",
					i+1, pRef, reflect.TypeOf(gen), p)
				continue gloop
			}
		}
		if _, ok := <-ch; ok {
			t.Errorf("%s.Iterator(0, %d) returning too many primes.\n",
				reflect.TypeOf(gen), limit)
		}
	}
}
