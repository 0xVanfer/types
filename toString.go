package types

import (
	"math/big"
	"strconv"
	"strings"

	common_eth "github.com/ethereum/go-ethereum/common"
)

func ToString(input any) string {
	var str string
	switch v := input.(type) {
	case string:
		str = v
	case int:
		str = strconv.FormatInt(int64(v), 10)
	case int8:
		str = strconv.FormatInt(int64(v), 10)
	case int16:
		str = strconv.FormatInt(int64(v), 10)
	case int32:
		str = strconv.FormatInt(int64(v), 10)
	case int64:
		str = strconv.FormatInt(v, 10)
	case uint:
		str = strconv.FormatInt(int64(v), 10)
	case uint8:
		str = strconv.FormatInt(int64(v), 10)
	case uint16:
		str = strconv.FormatInt(int64(v), 10)
	case uint32:
		str = strconv.FormatInt(int64(v), 10)
	case uint64:
		str = strconv.FormatInt(int64(v), 10)
	case float64:
		str = strconv.FormatFloat(v, 'f', -1, 64)
	case *big.Int:
		str = v.String()
	case common_eth.Address:
		str = v.String()
	case []byte:
		str = string(v)
	}
	return str
}

func ToLowerString(input any) string {
	str := ToString(input)
	return strings.ToLower(str)
}

func ToUpperString(input any) string {
	str := ToString(input)
	return strings.ToUpper(str)
}
