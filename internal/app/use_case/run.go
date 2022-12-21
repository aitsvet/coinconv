package use_case

import (
	"fmt"

	"github.com/aitsvet/coinconv/internal/app/model"
)

type AmountReader interface {
	Read(from string) (to model.Amount, err error)
}

type CurrencyReader interface {
	Read(from string) (to model.Currency, err error)
}

type Converter interface {
	Convert(what model.Amount, from, to model.Currency) (model.Amount, error)
}

type CurrencyPrinter interface {
	Print(amount model.Amount, currency model.Currency) string
}

type Convert struct {
	amountReader   AmountReader
	currencyReader CurrencyReader
	converter      Converter
	printer        CurrencyPrinter
}

func NewConvert(a AmountReader, c CurrencyReader, v Converter, p CurrencyPrinter) *Convert {
	return &Convert{amountReader: a, currencyReader: c, converter: v, printer: p}
}

func (c *Convert) Run(args []string) (string, error) {
	if len(args) != 4 {
		return "", fmt.Errorf("usage: %s <amount> <source> <dest>", args[0])
	}
	what, err := c.amountReader.Read(args[1])
	if err != nil {
		return "", fmt.Errorf("parsing amount: %w", err)
	}
	from, err := c.currencyReader.Read(args[2])
	if err != nil {
		return "", fmt.Errorf("parsing source symbol: %w", err)
	}
	to, err := c.currencyReader.Read(args[3])
	if err != nil {
		return "", fmt.Errorf("parsing dest symbol: %w", err)
	}
	result, err := c.converter.Convert(what, from, to)
	if err != nil {
		return "", fmt.Errorf("converting: %w", err)
	}
	return c.printer.Print(result, to), nil
}
