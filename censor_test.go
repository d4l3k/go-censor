package censor

import (
	"testing"
)

func TestCensor(t *testing.T) {
	ce := New([]string{
		"test",
		"testfoo",
		"bar",
	})
	ce.Pattern = "1234"

	cases := []struct {
		input string
		want  string
	}{
		{
			"test",
			"1234",
		},
		{
			"testfoobah",
			"1234123bah",
		},
		{
			"",
			"",
		},
		{
			"bah",
			"bah",
		},
		{
			"foobarfoo",
			"foo412foo",
		},
	}

	for i, c := range cases {
		out := ce.CensorString(c.input)
		if out != c.want {
			t.Fatalf("%d. Matches(%q) = %+v; not %+v", i, c.input, out, c.want)
		}
	}
}
