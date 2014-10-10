// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

// Package lucas computes terms of Lucas sequences.
package lucas

// U computes term n of Lucas sequence of the first kind U(P,Q)
//
// Recurrence relation, using U for U(P,Q):
//
//  U₀ := 0
//  U₁ := 1
//  Uₙ := P*Uₙ₋₁ - Q*Uₙ₋₂
func U(n, p, q int) int {
	return lucas(n, p, q, 0, 1)
}

// V computes term n of Lucas sequence of the second kind V(P,Q)
//
// Recurrence relation, using V for V(P,Q):
//
//  V₀ := 2
//  V₁ := P
//  Vₙ := P*Vₙ₋₁ - Q*Vₙ₋₂
func V(n, p, q int) int {
	return lucas(n, p, q, 2, p)
}

func lucas(n, p, q, n0, n1 int) int {
	if n < 1 {
		return n0
	}
	for n--; n > 0; n-- {
		n0, n1 = n1, p*n0-q*n1
	}
	return n1
}
