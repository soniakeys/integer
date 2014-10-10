// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package lucas_test

import (
	"fmt"

	"github.com/soniakeys/integer/lucas"
)

func ExampleU() {
	// U(1,-1) are Fibonacci numbers
	for n := 0; n <= 6; n++ {
		fmt.Println(n, lucas.U(n, 1, -1))
	}
	// Output:
	// 0 0
	// 1 1
	// 2 1
	// 3 2
	// 4 3
	// 5 5
	// 6 8
}

func ExampleV() {
	// V(1,-1) are Lucas numbers
	for n := 0; n <= 6; n++ {
		fmt.Println(n, lucas.V(n, 1, -1))
	}
	// Output:
	// 0 2
	// 1 1
	// 2 3
	// 3 4
	// 4 7
	// 5 11
	// 6 18
}
