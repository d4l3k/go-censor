package trie

import (
	"reflect"
	"testing"
)

func TestMatches(t *testing.T) {
	tr := New()
	tr.AddString("test")
	tr.AddString("testfoo")
	tr.AddString("bar")

	cases := []struct {
		input string
		want  []int
	}{
		{
			"test",
			[]int{4},
		},
		{
			"testfoobar",
			[]int{4, 7},
		},
		{
			"",
			nil,
		},
		{
			"bah",
			nil,
		},
	}

	for i, c := range cases {
		out := tr.Matches([]byte(c.input))
		if !reflect.DeepEqual(out, c.want) {
			t.Fatalf("%d. Matches(%q) = %+v; not %+v", i, c.input, out, c.want)
		}
	}
}
