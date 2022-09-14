package types

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func ToDecimal(input any) decimal.Decimal {
	var num decimal.Decimal
	switch v := input.(type) {
	case decimal.Decimal:
		num = v
	case float64:
		num = decimal.NewFromFloat(v)
	case int:
		num = decimal.NewFromInt(int64(v))
	case int8:
		num = decimal.NewFromInt(int64(v))
	case int32:
		num = decimal.NewFromInt(int64(v))
	case int64:
		num = decimal.NewFromInt(v)
	case uint:
		num = decimal.NewFromInt(int64(v))
	case uint8:
		num = decimal.NewFromInt(int64(v))
	case uint32:
		num = decimal.NewFromInt(int64(v))
	case uint64:
		num = decimal.NewFromInt(int64(v))
	case *big.Int:
		num = decimal.NewFromBigInt(v, 0)
	case string:
		num, _ = decimal.NewFromString(v)
	}
	return num
}
