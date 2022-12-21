package cmc_currency

import (
	"testing"
)

func TestCMCCurrency(t *testing.T) {
	r := new(Reader)
	for n, c := range []string{"USD", "BTC"} {
		a, err := r.Read(c)
		if err != nil {
			t.Fatalf("case %d: expected no error, got %v", n, err)
		}
		if c != a.Symbol() {
			t.Fatalf("case %d: expected result %s, got %s", n, c, a.Symbol())
		}
	}
}
