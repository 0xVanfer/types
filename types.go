package types

import (
	"math/big"

	common_eth "github.com/ethereum/go-ethereum/common"
)

type Ordered interface {
	OrderedNumber | ~string
}

type Number interface {
	OrderedNumber | BigNumber
}
type OrderedNumber interface {
	Integer | Float
}

type BigNumber interface {
	~*big.Int | ~*big.Float
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type AddressTypes interface {
	string | common_eth.Address
}
