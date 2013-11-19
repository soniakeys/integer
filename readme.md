Integer
=======
Various integer routines.

[![Build Status](https://travis-ci.org/soniakeys/integer.png)](https://travis-ci.org/soniakeys/integer)

[![GoDoc](https://godoc.org/github.com/garyburd/gddo?status.png)](http://godoc.org/github.com/soniakeys/integer)

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/soniakeys/integer)

Prime
-----
Ways of computing prime numbers, [OEIS A000040.](http://oeis.org/A000040)
-  Sieve, a sieve of Eratosthenese.
-  PQueue, a priority queue.
-  SPRP, a strong probable-prime test.
-  Segment, a parallel segmented sieve.

Swing
-----
Computation of swinging factorials, [OEIS A056040.](http://oeis.org/A056040)

Factorial
---------
Ways of computing the factorial, [OEIS A000142.](http://oeis.org/A000142)
-  Split Recursive, very efficient and without computation of prime numbers.
-  Prime Swing, most efficient for large numbers.  Involves computing swinging factorials, which in turn involves computing prime numbers.
-  Double Factorial [OEIS A001147.](http://oeis.org/A001147)

Binomial
--------
Efficient computation of the [binomial coefficient](http://en.wikipedia.org/wiki/Binomial_coefficient).

Xmath
-----
Some simple support functions.
