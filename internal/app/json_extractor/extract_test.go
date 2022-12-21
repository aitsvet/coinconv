package json_extractor

import "testing"

type testCase struct {
	path              []string
	json, result, err string
}

func TestExtract(t *testing.T) {
	e := StringField{}
	for n, c := range []testCase{
		{[]string{"a"}, `{"a": "b"}`, "b", ""},
		{[]string{"a"}, `{"a": 1.01}`, "1.01", ""},
		{[]string{"a", "b", "c"}, `{"a": {"b": {"c": 1.01}}}`, "1.01", ""},
		{[]string{"a"}, `{"a": {"b": 1.01}}`, "", "could not convert to string .a"},
		{[]string{"a"}, `{"b": 1.01}`, "", "could not extract key at .a"},
		{[]string{"a", "b"}, `{"a": 1.01}`, "", "could not convert to map at .a"},
		{[]string{"a", "b"}, ``, "", "EOF"},
	} {
		result, err := e.Extract(c.json, c.path)
		if err != nil {
			if err.Error() != c.err {
				t.Fatalf("case %d: expected error %s, got %s", n, c.err, err)
			}
		} else {
			if result != c.result {
				t.Fatalf("case %d: expected result %s, got %s", n, c.err, err)
			}
		}
	}
}
