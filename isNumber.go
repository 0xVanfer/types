package types

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func IsNumber(input any) bool {
	switch input.(type) {
	case int:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case float32:
		return true
	case float64:
		return true
	case big.Int:
		return true
	case big.Float:
		return true
	case decimal.Decimal:
		return true
	}
	return false
}
