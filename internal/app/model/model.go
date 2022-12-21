package model

import "math/big"

type Amount interface {
	BigFloat() big.Float
}

type Currency interface {
	Symbol() string
}

type Rate interface {
	ApplyTo(what Amount) (result Amount)
}
