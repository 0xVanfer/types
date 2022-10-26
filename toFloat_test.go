package types

import (
	"fmt"
	"math/big"
)

func ExampleToFloat64() {
	fmt.Println(ToFloat64(1))
	fmt.Println(ToFloat64(1.4))
	fmt.Println(big.NewFloat(1.8))
	// Output:
	// 1
	// 1.4
	// 1.8
}
