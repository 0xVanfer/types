package types

import (
	"fmt"
	"math/big"
)

func ExampleToDecimal() {
	fmt.Println(ToDecimal(big.NewInt(9)))
	fmt.Println(ToDecimal(9.9))
	fmt.Println(ToDecimal("9.99"))
	// Output:
	// 9
	// 9.9
	// 9.99
}
