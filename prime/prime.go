// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package prime contains a variety of prime number generators and
// a common interface for accessing them.
//
// Although primes numbers are prime numbers, an interface allows
// for generators with tradeoffs of time, space, and complexity.
package prime

// Generator is a minimal interface for prime number generators.
// Types implementing Generator may of course have additional useful methods.
type Generator interface {
	// Implementations typically generate primes up to some specified limit.
	// If an implementation can generate an arbitrarily long sequence of
	// primes, Limit() should return the maximum value for a uint64.
	Limit() uint64

	// Iterate shoulde iterate over the range of min to max,
	// inclusive, and call the visitor function for each prime number.
	//
	// It should return an ok status of false if iteration is not possible
	// because the specified max is greater than Limit().
	//
	// It should generally return true if iteration is possible, even if
	// no primes happen to be between the specified bounds or if the visitor
	// function terminates iteration early.
	//
	// The visitor function normally returns false.  Iterate should
	// interpret a true return from the visitor function as a request to
	// terminate iteration.  Iterate should then return true.
	Iterate(min, max uint64, visitor Visitor) (ok bool)
}

// Visitor function passed to Iterate method of a Generator.
//
// A visitor function should return false to continue iteration.
// Iterate interprets a true result from the visitor function as a request
// to terminate iteration.
type Visitor func(prime uint64) (terminate bool)

// Iterator returns a channel suitable for iteration, for example, with a
// range clause.  The channel returns primes in the range between a minimum
// and maximum value, inclusive.
//
// The variadic bounds can be 0, 1, or 2 paramaters.  The first, if present,
// is the minimum value of prime number.  The second, if present, is the
// maximum.  If no maximum is given, the default is the value returned by
// Limit().  If no minimum is given, the default is 2.
//
// Iterator returns nil if more than 3 parameters total are given,
// or if the specified maximum is greater than Limit().
func Iterator(pg Generator, bounds ...uint64) chan uint64 {
	min, max, ok := parameterBounds(pg, bounds)
	if !ok {
		return nil
	}
	r := make(chan uint64)
	go func() {
		pg.Iterate(min, max, func(prime uint64) bool {
			r <- prime
			return false
		})
		close(r)
	}()
	return r
}

// Primes returns a slice containing prime numbers in the specified range.
//
// The variadic bounds can be 0, 1, or 2 paramaters.  The first, if present,
// is the minimum value of prime number.  The second, if present, is the
// maximum.  If no maximum is given, the default is the value returned by
// Limit().  If no minimum is given, the default is 0.
//
// Primes returns nil if more than 3 parameters total are given,
// or if the specified maximum is greater than Limit().
//
// Note well!  In the case of a stream generator with no limit,
// calling Primes with no specified maximum is an attempt to generate
// all primes < MaxInt64.
func Primes(pg Generator, bounds ...uint64) []uint64 {
	min, max, ok := parameterBounds(pg, bounds)
	if !ok {
		return nil
	}
	var i int64
	pg.Iterate(min, max, func(_ uint64) bool {
		i++
		return false
	})
	r := make([]uint64, i)
	i = 0
	pg.Iterate(min, max, func(prime uint64) bool {
		r[i] = prime
		i++
		return false
	})
	return r
}

func parameterBounds(pg Generator, bounds []uint64) (min, max uint64, ok bool) {
	switch len(bounds) {
	case 0:
		max = pg.Limit()
	case 1:
		min, max = bounds[0], pg.Limit()
	case 2:
		min, max = bounds[0], bounds[1]
		if max > pg.Limit() {
			return
		}
	default:
		return
	}
	return min, max, true
}
