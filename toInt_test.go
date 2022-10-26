package types

import (
	"fmt"
	"math/big"
)

func ExampleToInt64() {
	fmt.Println(ToInt64(1.4))
	fmt.Println(ToInt64(1.8))
	fmt.Println(ToInt64(big.NewInt(2)))
	fmt.Println(ToInt64(big.NewFloat(2.6)))
	fmt.Println(ToInt64("0x11"))
	fmt.Println(ToInt64("0o66"))
	fmt.Println(ToInt64(true))
	// Output:
	// 1
	// 2
	// 2
	// 3
	// 17
	// 54
	// 1
}
