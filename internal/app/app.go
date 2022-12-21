package app

import (
	"github.com/aitsvet/coinconv/internal/app/big_amount"
	"github.com/aitsvet/coinconv/internal/app/big_rate"
	"github.com/aitsvet/coinconv/internal/app/cmc_client"
	"github.com/aitsvet/coinconv/internal/app/cmc_currency"
	"github.com/aitsvet/coinconv/internal/app/cmc_rater"
	"github.com/aitsvet/coinconv/internal/app/coin_converter"
	"github.com/aitsvet/coinconv/internal/app/coin_printer"
	"github.com/aitsvet/coinconv/internal/app/json_extractor"
	"github.com/aitsvet/coinconv/internal/app/use_case"
)

func Configure(baseURL, apiKey string) *use_case.Convert {
	return use_case.NewConvert(
		&big_amount.Reader{},
		&cmc_currency.Reader{},
		coin_converter.New(
			cmc_rater.New(
				cmc_client.New(baseURL, apiKey),
				&json_extractor.StringField{},
				&big_rate.Reader{})),
		&coin_printer.Simple{})
}
