package types

import (
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

// Convert anything to boolean.
func ToBool(input any) bool {
	switch v := input.(type) {
	case string:
		return strings.EqualFold(v, "true")
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	case *big.Int:
		return v.String() != "0"
	case decimal.Decimal:
		return v.String() != "0"
	case []byte:
		return strings.EqualFold(string(v), "true")
	case bool:
		return v
	default:
		return false
	}
}
