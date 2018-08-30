package censor

// NewTrie returns a new Trie.
func NewTrie() *Trie {
	return &Trie{
		children: map[byte]*Trie{},
		end:      false,
	}
}

// Trie represents a prefix tree.
type Trie struct {
	children map[byte]*Trie
	end      bool
}

// AddString adds a word to the Trie. See Add.
func (t *Trie) AddString(word string) {
	t.Add([]byte(word))
}

// Add adds a word to the Trie.
func (t *Trie) Add(word []byte) {
	if len(word) == 0 {
		t.end = true
		return
	}

	subTrie, ok := t.children[word[0]]
	if !ok {
		subTrie = NewTrie()
		t.children[word[0]] = subTrie
	}
	subTrie.Add(word[1:])
}

// Matches returns the list of lengths that match the prefix of the supplied
// array.
func (t *Trie) Matches(buf []byte) []int {
	return t.matchesInternal(buf, 0, nil)
}

func (t *Trie) matchesInternal(buf []byte, lenSoFar int, matches []int) []int {
	if t.end {
		matches = append(matches, lenSoFar)
	}
	if len(buf) == 0 {
		return matches
	}

	subTrie, ok := t.children[buf[0]]
	if !ok {
		return matches
	}
	return subTrie.matchesInternal(buf[1:], lenSoFar+1, matches)
}
