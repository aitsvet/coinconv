package big_amount

import (
	"math/big"

	"github.com/aitsvet/coinconv/internal/app/model"
)

type Amount struct {
	value big.Float
}

func New(value big.Float) *Amount {
	return &Amount{value: value}
}

func (a *Amount) BigFloat() big.Float {
	return a.value
}

type Reader struct{}

func (f *Reader) Read(from string) (model.Amount, error) {
	result := &Amount{}
	_, _, err := result.value.Parse(from, 10)
	return result, err
}
