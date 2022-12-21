package big_amount

import (
	"testing"
)

type testCase struct {
	value, err string
}

func TestBigAmount(t *testing.T) {
	r := new(Reader)
	for n, c := range []testCase{
		{value: "0"},
		{value: "1"},
		{value: "0.1"},
		{value: "0.01"},
		{value: "1.01"},
		{value: "", err: "EOF"},
		{value: "a", err: "number has no digits"},
	} {
		a, err := r.Read(c.value)
		if err != nil {
			if err.Error() != c.err {
				t.Fatalf("case %d: expected error %s, got %v", n, c.err, err)
			}
		} else {
			f := a.BigFloat()
			if c.value != f.String() {
				t.Fatalf("case %d: expected result %s, got %s", n, c.value, f.String())
			}
		}
	}
}
