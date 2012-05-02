Integer
=======
Various integer routines.

Prime
-----
Ways of computing prime numbers, [OEIS A000040.](http://oeis.org/A000040)
-  Sieve, a sieve of Eratosthenese.

Swing
-----
Computation of swinging factorials, [OEIS A056040.](http://oeis.org/A056040)

Factorial
---------
Ways of computing the factorial, [OEIS A000142.](http://oeis.org/A000142)
-  Naive algorithm, for comparison with more efficient ones.
-  Split Recursive, very efficient and without computation of prime numbers.
-  Swing, using precomputed swinging factorials, but otherwise not performing computation of prime numbers.  More efficient that split recursive for large numbers.
-  Prime Swing, most efficient for large numbers.  Involves computing swinging factorials, which in turn involves computing prime numbers.

-  Double Factorial [OEIS A001147.](http://oeis.org/A001147)

Binomial
--------
Efficient computation of the [binomial coefficient](http://en.wikipedia.org/wiki/Binomial_coefficient).

Xmath
-----
Some simple support functions.
