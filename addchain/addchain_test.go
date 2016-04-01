package addchain_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/soniakeys/integer/addchain"
)

/*
func ExampleStrategy_Chain() {
	a := addchain.Total.Chain(1000).AddChain()
	fmt.Println(a)
	fmt.Println("len", len(a))
	// Output:
}
*/

func Example_exponentiation() {
	x, y, m := 123, uint(456), 789
	s := addchain.Dichotomic.Chain(y)
	// Compare to source code for StarChain.AddChain.
	// Other applications can be coded similarly.
	e := make([]int, len(s)+1)
	e[0] = x
	for i, x := range s {
		e[i+1] = (e[i] * e[x]) % m
	}
	fmt.Println(e[len(s)])

	xx := big.NewInt(int64(x))
	yy := big.NewInt(int64(y))
	mm := big.NewInt(int64(m))
	fmt.Println(xx.Exp(xx, yy, mm))
	// Output:
	// 699
	// 699
}

func ExampleStarChain_AddChain() {
	s := addchain.StarChain{0, 0, 1}
	fmt.Println(s.AddChain())
	// Output:
	// [1 2 3 5]
}

func ExampleSingletonStrategy_Chain() {
	fmt.Println(addchain.Binary.Chain(87).AddChain())
	fmt.Println(addchain.CoBinary.Chain(87).AddChain())
	fmt.Println(addchain.Dichotomic.Chain(87).AddChain())
	// Output:
	// [1 2 4 5 10 20 21 42 43 86 87]
	// [1 2 3 5 10 11 21 22 43 44 87]
	// [1 2 3 6 7 10 20 40 80 87]
}

func TestBinary(t *testing.T) {
	n := 0
	for i := uint(2); i <= 1000; i++ {
		n += len(addchain.Binary.Chain(i))
	}
	t.Log("Binary:    ", n)
}
func TestDichotomic(t *testing.T) {
	n := 0
	for i := uint(2); i <= 1000; i++ {
		n += len(addchain.Dichotomic.Chain(i))
	}
	t.Log("Dichotomic:", n)
}
func TestDyadic(t *testing.T) {
	n := 0
	for i := uint(2); i <= 1000; i++ {
		n += len(addchain.Dyadic.Chain(i))
	}
	t.Log("Dyadic:    ", n)
}
func TestFermat(t *testing.T) {
	n := 0
	for i := uint(2); i <= 1000; i++ {
		n += len(addchain.Fermat.Chain(i))
	}
	t.Log("Fermat:    ", n)
}

func TestTotal(t *testing.T) {
	n := 0
	for i := uint(2); i <= 50; i++ {
		n += len(addchain.Total.Chain(i))
	}
	t.Log("Total:    ", n)
}

const min = 1000
const max = 9999

var rp = rand.Perm(max - min + 1)

func BenchmarkBinary____(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addchain.Binary.Chain(min + uint(rp[i%(max-min+1)]))
	}
}

func BenchmarkCoBinary__(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addchain.CoBinary.Chain(min + uint(rp[i%(max-min+1)]))
	}
}

func BenchmarkDichotomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addchain.Dichotomic.Chain(min + uint(rp[i%(max-min+1)]))
	}
}
