package coin_converter

import (
	"github.com/aitsvet/coinconv/internal/app/model"
)

type Rater interface {
	GetRate(from, to model.Currency) (model.Rate, error)
}

type Converter struct {
	rater Rater
}

func New(rater Rater) *Converter {
	return &Converter{rater: rater}
}

func (c *Converter) Convert(what model.Amount, from, to model.Currency) (model.Amount, error) {
	r, err := c.rater.GetRate(from, to)
	if err != nil {
		return nil, err
	}
	return r.ApplyTo(what), nil
}
