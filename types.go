package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Types that can use <, >, etc.
type Ordered interface {
	OrderedNumber | ~string
}

type Number interface {
	OrderedNumber | BigNumber | decimal.Decimal
}

// Number types that can use <, >, etc.
type OrderedNumber interface {
	Integer | Float
}

type BigNumber interface {
	~*big.Int | ~*big.Float
}

// All kinds of integer.
type Integer interface {
	Signed | Unsigned
}

// All kinds of sighed integer.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// All kinds of unsigned integer.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type AddressTypes interface {
	string | common.Address
}
