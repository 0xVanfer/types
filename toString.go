package types

import (
	"math/big"
	"strconv"
	"strings"

	common_eth "github.com/ethereum/go-ethereum/common"
)

func ToString(input any) string {
	switch v := input.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatInt(int64(v), 10)
	case uint8:
		return strconv.FormatInt(int64(v), 10)
	case uint16:
		return strconv.FormatInt(int64(v), 10)
	case uint32:
		return strconv.FormatInt(int64(v), 10)
	case uint64:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case *big.Int:
		return v.String()
	case common_eth.Address:
		return v.String()
	case []byte:
		return string(v)
	default:
		return ""
	}
}

func ToLowerString(input any) string {
	str := ToString(input)
	return strings.ToLower(str)
}

func ToUpperString(input any) string {
	str := ToString(input)
	return strings.ToUpper(str)
}
