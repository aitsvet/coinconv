package coin_printer

import (
	"math/big"
	"testing"
)

type mockAmount struct {
	value *big.Float
}

func (a mockAmount) BigFloat() big.Float {
	return *a.value
}

type mockCurrency struct {
	symbol string
}

func (c *mockCurrency) Symbol() string {
	return c.symbol
}

func TestPrint(t *testing.T) {
	p := &Simple{}
	if p.Print(&mockAmount{big.NewFloat(0.01)}, &mockCurrency{"USD"}) != "0.01 USD" {
		t.Fatalf("could not print \"0.01 USD\"")
	}
}
