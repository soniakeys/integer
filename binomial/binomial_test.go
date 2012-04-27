package binomial

import (
	"math/big"
	"testing"

	"github.com/soniakeys/integer/prime"
)

func TestBinomial(t *testing.T) {
	const n = 234567
	const k = 120033
	p := prime.MakePrimes(n)
	a := Binomial(p, n, k)

	dtrunc := int64(float64(a.BitLen())*.30103) - 10
	var first, rest big.Int
	rest.Exp(first.SetInt64(10), rest.SetInt64(dtrunc), nil)
	first.Quo(a, &rest)
	fstr := first.String()
	if len(fstr)+int(dtrunc) != 70581 {
		t.Error("wrong number of digits in answer")
	}
	if fstr != "83947130061" {
		t.Error("first digits of answer are wrong")
	}
}
