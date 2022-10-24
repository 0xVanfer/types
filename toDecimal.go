package types

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func ToDecimal(input any) decimal.Decimal {
	switch v := input.(type) {
	case decimal.Decimal:
		return v
	case float32:
		return decimal.NewFromFloat32(v)
	case float64:
		return decimal.NewFromFloat(v)
	case int:
		return decimal.NewFromInt(int64(v))
	case int8:
		return decimal.NewFromInt(int64(v))
	case int32:
		return decimal.NewFromInt32(v)
	case int64:
		return decimal.NewFromInt(v)
	case uint:
		return decimal.NewFromInt(int64(v))
	case uint8:
		return decimal.NewFromInt(int64(v))
	case uint32:
		return decimal.NewFromInt(int64(v))
	case uint64:
		return decimal.NewFromInt(int64(v))
	case *big.Int:
		return decimal.NewFromBigInt(v, 0)
	case *big.Float:
		return decimal.NewFromFloat(ToFloat64(v))
	case string:
		num, _ := decimal.NewFromString(v)
		return num
	default:
		return decimal.NewFromInt(0)
	}
}
