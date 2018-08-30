package censor

import (
	"reflect"
	"testing"
)

func TestTrieMatches(t *testing.T) {
	tr := NewTrie()
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
