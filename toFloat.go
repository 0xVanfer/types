package types

import (
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

// Convert number or string into float64.
//
// Input must be number or number string, otherwise will return 0.
func ToFloat64(input any) float64 {
	switch v := input.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case string:
		num, _ := strconv.ParseFloat(v, 64)
		return num
	case *big.Int:
		num, _ := strconv.ParseFloat(v.String(), 64)
		return num
	case *big.Float:
		num, _ := strconv.ParseFloat(v.String(), 64)
		return num
	case decimal.Decimal:
		return v.InexactFloat64()
	default:
		return 0
	}
}
