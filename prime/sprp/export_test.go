package sprp_test

import (
	"testing"

	"github.com/soniakeys/integer/prime/sprp"
)

func TestLimit(t *testing.T) {
	if l := sprp.New().Limit(); l != 1<<32-1 {
		t.Errorf("Limit() returned %d. 1<<32-1 expected.", l)
	}
}

func TestIterate(t *testing.T) {
	s := sprp.New()

	// reproduce 10 primes hardcoded right here.
	ten := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// check that the 10 are right
	var i int
	ok := s.Iterate(0, ten[9], func(prime uint64) bool {
		if prime != ten[i] {
			t.Fatal("Incorrect prime.  Expected", ten[i], "got", prime)
		}
		i++
		return false
	})
	if !ok {
		t.Error("Iterate returned false on early termination.")
	}

	// some primes across the base set 0/1 boundary
	// source: prime pages
	b01 := []uint64{1373557, 1373563, 1373591, 1373611, 1373627,
		1373639, 1373677, 1373683, 1373689, 1373717,
		1373761, 1373777, 1373789, 1373803, 1373819}

	// check that they are right
	i = 0
	s.Iterate(b01[0], b01[len(b01)-1], func(prime uint64) bool {
		if prime != b01[i] {
			t.Fatal("Incorrect prime.  Expected", b01[i], "got", prime)
		}
		i++
		return false
	})
}
