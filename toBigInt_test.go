package types

import (
	"fmt"
)

func ExampleToBigInt() {
	fmt.Println(ToBigInt(1))
	fmt.Println(ToBigInt(1.4))
	fmt.Println(ToBigInt(1.8))
	fmt.Println(ToBigInt("2"))
	fmt.Println(ToBigInt("2.3"))
	// Output:
	// 1
	// 1
	// 2
	// 2
	// <nil>
}
