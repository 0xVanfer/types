package types

import "math/big"

func ToBigInt(input any) *big.Int {
	var amt *big.Int
	switch v := input.(type) {
	case *big.Int:
		amt = v
	case int:
		amt = big.NewInt(int64(v))
	case int16:
		amt = big.NewInt(int64(v))
	case int32:
		amt = big.NewInt(int64(v))
	case int64:
		amt = big.NewInt(v)
	case uint:
		amt = big.NewInt(int64(v))
	case uint16:
		amt = big.NewInt(int64(v))
	case uint32:
		amt = big.NewInt(int64(v))
	case uint64:
		amt = big.NewInt(int64(v))
	case float64:
		amt = big.NewInt(ToInt64(v))
	case string:
		amt = big.NewInt(ToInt64(v))
	}
	return amt
}
