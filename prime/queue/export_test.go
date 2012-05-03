package queue_test

// File contains tests for the SievePQ type and related functions.
// At a minimum, tests should cover PrimeGenerator methods.

import (
	"testing"

	"github.com/soniakeys/integer/prime/queue"
)

func TestPQLimit(t *testing.T) {
	if (queue.PQueue{}).Limit() != 1<<64-1 {
		t.Error("queue.PQueue Limit() fail")
	}
}

func TestPQSmall(t *testing.T) {
	// reproduce 20 primes hardcoded right here.
	twenty := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71}

	// check that the 20 are right
	var i int
	queue.PQueue{}.Iterate(0, twenty[19], func(prime uint64) bool {
		if prime != twenty[i] {
			t.Fatal("Incorrect prime.  Expected", twenty[i], "got", prime)
		}
		i++
		return false
	})
}
