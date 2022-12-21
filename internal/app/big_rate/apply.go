package big_rate

import (
	"math/big"

	"github.com/aitsvet/coinconv/internal/app/big_amount"
	"github.com/aitsvet/coinconv/internal/app/model"
)

type Rate struct {
	value big.Float
}

func (r *Rate) ApplyTo(what model.Amount) (to model.Amount) {
	result := what.BigFloat()
	result.Mul(&result, &r.value)
	return big_amount.New(result)
}

type Reader struct{}

func (f *Reader) Read(from string) (model.Rate, error) {
	result := &Rate{}
	_, _, err := result.value.Parse(from, 10)
	return result, err
}
