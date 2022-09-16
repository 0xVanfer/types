package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Return go-ethereum address.
func ToAddress(input any) common.Address {
	switch v := input.(type) {
	case common.Address:
		return v
	case string:
		return common.HexToAddress(v)
	default:
		return common.Address{}
	}
}
