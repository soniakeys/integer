package sieve

import "testing"

func TestSmallComposites(t *testing.T) {
	sc := smallComposites
	smallComposites = nil
	scl := smallCompositeLimit
	smallCompositeLimit = 0
	s := New(scl)
	if len(s.isComposite) < len(sc) {
		goto bad
	}
	for i, b := range sc {
		if b != s.isComposite[i] {
			goto bad
		}
	}
	return // ok
bad:
	t.Log("bad small composite list.")
	t.Log("valid:", s.isComposite)
	t.Log("found:", sc)
	t.Fail()
}
