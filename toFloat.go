package types

import (
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

func ToFloat64(input any) float64 {
	var num float64
	switch v := input.(type) {
	case float64:
		return v
	case int:
		num = float64(v)
	case int8:
		num = float64(v)
	case int16:
		num = float64(v)
	case int32:
		num = float64(v)
	case int64:
		num = float64(v)
	case uint:
		num = float64(v)
	case uint8:
		num = float64(v)
	case uint16:
		num = float64(v)
	case uint64:
		num = float64(v)
	case string:
		num, _ = strconv.ParseFloat(v, 64)
	case *big.Int:
		num, _ = strconv.ParseFloat(v.String(), 64)
	case decimal.Decimal:
		num = v.InexactFloat64()
	}
	return num
}
