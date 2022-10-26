package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Convert string into go-ethereum address.
//
// Input must be go-ethereum address or length 42 string,
// otherwise will return 0x0.
func ToAddress(input any) common.Address {
	switch v := input.(type) {
	case common.Address:
		return v
	case string:
		if len(v) != 42 {
			return common.Address{}
		}
		return common.HexToAddress(v)
	default:
		return common.Address{}
	}
}
