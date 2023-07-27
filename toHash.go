package types

import (
	"math/big"

	common_eth "github.com/ethereum/go-ethereum/common"
)

// Convert anything to hash.
func ToHash(input any) common_eth.Hash {
	switch v := input.(type) {
	case string:
		return common_eth.HexToHash(v)
	case *big.Int:
		return common_eth.BigToHash(v)
	case []byte:
		return common_eth.BytesToHash(v)
	default:
		return common_eth.HexToHash(ToString(v))
	}
}
