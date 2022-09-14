package types

import (
	common_eth "github.com/ethereum/go-ethereum/common"
)

func ToAddress(input any) common_eth.Address {
	var address common_eth.Address
	switch v := input.(type) {
	case common_eth.Address:
		address = v
	case string:
		address = common_eth.HexToAddress(v)
	}
	return address
}
