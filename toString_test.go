package types

import (
	"fmt"
)

func ExampleToString() {
	fmt.Println(ToString(1))
	fmt.Println(ToString(1.1))
	fmt.Println(ToString(true))
	// Output:
	// 1
	// 1.1
	// true
}
