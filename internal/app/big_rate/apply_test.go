package big_rate

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

func mock(v float64) mockAmount {
	return mockAmount{big.NewFloat(v)}
}

type testCase struct {
	from, to  mockAmount
	rate, err string
}

func TestBigRate(t *testing.T) {
	r := new(Reader)
	for n, c := range []testCase{
		{from: mock(1), rate: "5", to: mock(5)},
		{rate: "", err: "EOF"},
		{rate: "a", err: "number has no digits"},
	} {
		a, err := r.Read(c.rate)
		if err != nil {
			if err.Error() != c.err {
				t.Fatalf("case %d: expected error %s, got %v", n, c.err, err)
			}
		} else {
			result := a.ApplyTo(c.from).BigFloat()
			if c.to.value.String() != result.String() {
				t.Fatalf("case %d: expected result %s, got %s", n, c.to.value, &result)
			}
		}
	}
}
