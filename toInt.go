package types

import (
	"math"
	"math/big"
	"strconv"
)

func ToInt64(input any) int64 {
	var num int64
	switch v := input.(type) {
	case int:
		num = int64(v)
	case int8:
		num = int64(v)
	case int16:
		num = int64(v)
	case int32:
		num = int64(v)
	case int64:
		num = v
	case uint:
		num = int64(v)
	case uint8:
		num = int64(v)
	case uint16:
		num = int64(v)
	case uint32:
		num = int64(v)
	case uint64:
		num = int64(v)
	case string:
		if (len(v) > 2) && (v[:2] == "0x") {
			num, _ = strconv.ParseInt(v[2:], 16, 64)

		} else {
			num, _ = strconv.ParseInt(v, 10, 64)
		}
	case float64:
		// str := strconv.FormatFloat(v, 'f', -1, 64)
		// num, _ = strconv.Atoi(str)
		num = int64(math.Round(v))
	case *big.Int:
		num = v.Int64()
	case bool:
		if v {
			num = 1
		} else {
			num = 0
		}
	}
	return num
}

func ToInt(input any) int {
	return int(ToInt64(input))
}
