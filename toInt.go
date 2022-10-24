package types

import (
	"math"
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

func ToInt64(input any) int64 {
	switch v := input.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case string:
		if (len(v) > 2) && (v[:2] == "0x") {
			num, _ := strconv.ParseInt(v[2:], 16, 64)
			return num
		}
		if (len(v) > 2) && (v[:2] == "0o") {
			num, _ := strconv.ParseInt(v[2:], 8, 64)
			return num
		}
		if (len(v) > 2) && (v[:2] == "0b") {
			num, _ := strconv.ParseInt(v[2:], 2, 64)
			return num
		}
		num, _ := strconv.ParseInt(v, 10, 64)
		return num
	case float32:
		return int64(math.Round(float64(v)))
	case float64:
		return int64(math.Round(v))
	case *big.Int:
		return v.Int64()
	case bool:
		// true 1; false 0
		if v {
			return 1
		} else {
			return 0
		}
	case decimal.Decimal:
		return v.IntPart()
	default:
		return 0
	}

}

func ToInt(input any) int {
	return int(ToInt64(input))
}
