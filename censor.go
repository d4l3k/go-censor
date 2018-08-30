package censor

import "github.com/d4l3k/go-censor/trie"

// Censor contains a trie containing the words to censor.
type Censor struct {
	t       *trie.Trie
	Pattern string
}

// New returns a new censor.
func New(words []string) Censor {
	t := trie.New()
	for _, w := range words {
		t.AddString(w)
	}
	return Censor{
		t:       t,
		Pattern: "*",
	}
}

// BadWords returns a slice of (start, end) pairs of bad words in the provided
// buf.
func (c *Censor) BadWords(buf []byte) [][]int {
	var matches [][]int
	for i := range buf {
		matchLens := c.t.Matches(buf[i:])
		for _, l := range matchLens {
			matches = append(matches, []int{i, i + l})
		}
	}
	return matches
}

// Censor censors the provided buf with c.Pattern.
func (c *Censor) Censor(buf []byte) {
	matches := c.BadWords(buf)

	for _, match := range matches {
		fillPattern(buf, c.Pattern, match[0], match[1])
	}
}

// CensorString returns a new string that has been censored. See Censor.
// This requires two memory copies of the whole string.
func (c *Censor) CensorString(in string) string {
	buf := []byte(in)
	c.Censor(buf)
	return string(buf)
}

func fillPattern(buf []byte, pat string, start, end int) {
	for i := start; i < end; i++ {
		buf[i] = pat[i%len(pat)]
	}
}
