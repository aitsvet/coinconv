package cmc_currency

import (
	"github.com/aitsvet/coinconv/internal/app/model"
)

type CMCCurrency struct {
	symbol string
}

func (c *CMCCurrency) Symbol() string {
	return c.symbol
}

type Reader struct{}

func (f *Reader) Read(from string) (model.Currency, error) {
	return &CMCCurrency{symbol: from}, nil
}
