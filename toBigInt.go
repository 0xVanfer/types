package types

import "math/big"

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
	case float64:
		return big.NewInt(ToInt64(v))
	case string:
		return big.NewInt(ToInt64(v))
	default:
		return big.NewInt(0)
	}
}
