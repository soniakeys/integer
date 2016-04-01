// Copyright 2014 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Addchain computes addition chains.
//
// The method is based on continued fractions, as described in "Efficient
// computation of addition chains" by F. Bergeron, J. Berstel, and S. Brlek,
// published in Journal de théorie des nombres de Bordeaux, 6 no. 1 (1994),
// p. 21-38, accessed at http://www.numdam.org/item?id=JTNB_1994__6_1_21_0.
//
// The chains produced by this method are non-optimal.
package addchain

import (
//	"fmt"
	"github.com/soniakeys/integer/xmath"
)

// A StarChain represents an addition chain where each number > 1 is the sum
// of some previous number and the immediate previous number.
//
// Representation is []int where elements represent indexes into the
// represented addition chain.  The index stored is the index of one addend
// of the sum.  The other addend is implicit.  Example,
//
//  Addition chain:  1 2 3 5
//  StarChain:         0 0 1
//
// An addition chain of length n is represented by a StarChain of length n-1.
type StarChain []int

// AddChain returns the addition chain as a list of sums.
func (s StarChain) AddChain() []int {
	a := make([]int, len(s)+1)
	a[0] = 1
	for i, x := range s {
		a[i+1] = a[i] + a[x]
	}
	return a
}

// ⊗= operator.  modifies receiver.  Not constant time as the paper claimed
// to have with a linked list representation, but the idea here is that
// adjusting slice values like this is fast and that linked lists are slow.
func (s1 *StarChain) cMul(s2 StarChain) {
	p := *s1
	n := len(p)
	p = append(p, s2...)
	for i := n; i < len(p); i++ {
		p[i] += n
	}
	*s1 = p
}

// ⊕= operator.  modifies receiver.
// note that a StarChain holds indexes--argument j is an index to add.
func (p *StarChain) cAdd(j int) {
	*p = append(*p, j)
}

// SingletonStrategy is the γ function described in the paper.  γ returns
// a set of numbers in general, but certain γ functions return only single
// values.  This extra property allows some simplifications in the code for
// a strategy of this type.
type SingletonStrategy func(uint) uint

var Binary, CoBinary, Dichotomic SingletonStrategy

type Strategy func(uint) []uint

var Total, Dyadic, Fermat Strategy

func init() {
	// singletons
	Binary = func(n uint) uint {
		return n / 2
	}
	CoBinary = func(n uint) uint {
		return (n + 1) / 2
	}
	Dichotomic = func(n uint) uint {
		return n / (1 << ((xmath.Log2(n) + 1) / 2))
	}
	// non-singletons
	Total = func(n uint) []uint {
		s := make([]uint, n-2)
		for i := range s {
			s[i] = uint(i + 2)
		}
		// fmt.Printf(indent+" Total(%d) = %v\n", n, s)
		return s
	}
	Dyadic = func(n uint) []uint {
		s := make([]uint, xmath.Log2(n)-1)
		for i := range s {
			s[i] = n / (uint(1) << uint(len(s)-i))
		}
		// fmt.Printf(indent+" Dyadic(%d) = %v\n", n, s)
		return s
	}
	Fermat = func(n uint) []uint {
        s := make([]uint, 1+xmath.Log2(xmath.Log2(n)-1))
		for i := range s {
			s[i] = n / (uint(1) << (uint(1) << uint(len(s)-1-i)))
		}
	//	fmt.Printf(indent+" Fermat(%d) = %v\n", n, s)
		return s
	}
}

var indent string

type cm struct {
	γ Strategy
	m map[uint]StarChain
}

// Chain computes an addition chain, returning it in StarChain form.
func (γ Strategy) Chain(n uint) StarChain {
	cc := &cm{γ, map[uint]StarChain{3: StarChain{0, 0}}}
	s := cc.minChain(n)
	// fmt.Println("map dump:")
	// for n, c := range cc.m {
		// fmt.Println(n, c.AddChain())
	// }
	return s
}

// minChain returns a chain that is stored in the map.
func (cc *cm) minChain(n uint) StarChain {
	// fmt.Println(indent, "minChain(", n, ")")
	ns := indent
	indent += "  "
	defer func() { indent = ns }()
	if s, ok := cc.m[n]; ok {
		// fmt.Println(indent, "memoized")
		return s
	}
	// This is the function minChain as described in the paper.
	switch a := xmath.Log2(n); {
	case n == 1<<a:
		r := make(StarChain, a)
		for i := range r {
			r[i] = i
		}
		cc.m[n] = r
		// fmt.Println(indent, "It's a power of 2.  Returning", r.AddChain())
		return r
		//	case n == 3:
		//		// fmt.Println(indent, "It's 3.  Returning", StarChain{0, 0}.AddChain())
		//		return StarChain{0, 0}
	}
	min := int(n)
	var minS StarChain
	// var bestk uint
	// fmt.Println(indent, "Considering some possibilities...")
	for _, k := range cc.γ(n) {
		s, _ := cc.chain1(n, k)
		if len(s) < min {
			// fmt.Println(indent, "oh,", k, "was a good one")
			// bestk = k
			min = len(s)
			minS = s
		}
	}
	cc.m[n] = minS
	// fmt.Println(indent, "yeah,", bestk, "was best.  Returning", minS.AddChain())
	return minS
}

// chain1 is chain as described in the paper.
// return value c may or may not be stored in the map yet.
// return value x2 is index of n2 in c.AddChain() on return value c.
// X2 is not described in the paper but seems helpful for efficiency.
func (cc *cm) chain1(n ...uint) (c StarChain, x2 int) {
	// fmt.Println(indent, "chain1(", n, ")")
	ns := indent
	indent += "  "
	// ndbg := append([]uint{}, n...) // copy for debug output
	defer func() {
		// fmt.Println(indent, "for chain1(", ndbg, "), returning", c.AddChain(), "with x2 =", x2)
		indent = ns
	}()
	n1 := n[0]
	if len(n) == 1 {
		// fmt.Println(indent, "huh.  just one target,", n1)
		c := cc.minChain(n1)
		return c, len(c)
	}
	n2 := n[1]
	if n2 == 1 {
		// fmt.Println(indent, "huh.  n2 is 1, other target is", n1)
		return cc.minChain(n1), 0 // n2 was 1, index of 1 is always 0
	}
	q, r := n1/n2, n1%n2
	if r == 0 {
		// fmt.Println(indent, n2, "divides", n1, ".  Recursing for", n[1:])
		c, x2 = cc.chain1(n[1:]...)
		// fmt.Println(indent, "x2 from that was", x2, ".  Recursing for", q)
		c.cMul(cc.minChain(q))
		return
	}
	// fmt.Println(indent, "the trickiest case,", n2, "doesn't divide", n1, ".  r =", r)
	copy(n, n[1:])
	n[len(n)-1] = r
	// fmt.Println(indent, "Recursing with r tacked on the end")
	c, xr := cc.chain1(n...)
	// fmt.Println(indent, "xr from that is", xr)
	x2 = len(c)
	// fmt.Println(indent, "Recursing for", q, "now")
	c.cMul(cc.minChain(q))
	// fmt.Println(indent, "and now tacking r =", r, "on the end of the chain")
	c.cAdd(xr)
	return
}

// Chain computes an addition chain, returning it in StarChain form.
func (γ SingletonStrategy) Chain(n uint) StarChain {
	// This is the function minChain as described in the paper.
	switch a := xmath.Log2(n); {
	case n == 1<<a:
		r := make(StarChain, a)
		for i := range r {
			r[i] = i
		}
		return r
	case n == 3:
		return StarChain{0, 0}
	}
	s, _ := γ.chain1(n, γ(n))
	return s
}

// chain1 is chain as described in the paper.
// return value x2 is index of n2 in c.AddChain() on return value c.
// X2 is not described in the paper but seems helpful for efficiency.
func (γ SingletonStrategy) chain1(n1, n2 uint) (c StarChain, x2 int) {
	q, r := n1/n2, n1%n2
	if r == 0 {
		c = γ.Chain(n2)
		x2 = len(c)
		c.cMul(γ.Chain(q))
		return
	}
	c, xr := γ.chain1(n2, r)
	x2 = len(c)
	c.cMul(γ.Chain(q))
	c.cAdd(xr)
	return
}
