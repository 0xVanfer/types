package types

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// Convert number or string into *big.Int.
//
// Input must be number or number string, otherwise will return 0.
func ToBigInt(input any) *big.Int {
	switch v := input.(type) {
	case *big.Int:
		return v
	case int:
		return big.NewInt(int64(v))
	case int16:
		return big.NewInt(int64(v))
	case int32:
		return big.NewInt(int64(v))
	case int64:
		return big.NewInt(v)
	case uint:
		return big.NewInt(int64(v))
	case uint16:
		return big.NewInt(int64(v))
	case uint32:
		return big.NewInt(int64(v))
	case uint64:
		return big.NewInt(int64(v))
	case float32:
		return big.NewInt(ToInt64(v))
	case float64:
		return big.NewInt(ToInt64(v))
	case string:
		x := big.NewInt(0)
		big, _ := x.SetString(v, 10)
		return big
	case decimal.Decimal:
		return v.BigInt()
	default:
		return big.NewInt(0)
	}
}
