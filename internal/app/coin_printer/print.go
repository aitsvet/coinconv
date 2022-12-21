package coin_printer

import (
	"github.com/aitsvet/coinconv/internal/app/model"
)

type Simple struct{}

func (p *Simple) Print(amount model.Amount, currency model.Currency) string {
	value := amount.BigFloat()
	return value.String() + " " + currency.Symbol()
}
