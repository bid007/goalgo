package trie

import "reflect"

// Trie to represent trie data structure
type Trie struct {
	isWord bool
	root   map[rune]*Trie
}

// New creates an instance of trie
func New() *Trie {
	return &Trie{root: make(map[rune]*Trie)}
}

func (t *Trie) addWord(word string) {
	parent := t
	for _, ch := range word {
		child, ok := parent.root[ch]
		if !ok {
			parent.root[ch] = &Trie{root: make(map[rune]*Trie)}
			child = parent.root[ch]
		}
		parent = child
	}
	parent.isWord = true
}

// AddWords adds words in the trie
func (t *Trie) AddWords(words ...string) {
	for _, word := range words {
		t.addWord(word)
	}
}

// Exists takes a word and returns true if the word is in trie
func (t *Trie) Exists(word string) bool {
	parent := t
	for _, ch := range word {
		child, ok := parent.root[ch]
		if !ok {
			return false
		}
		parent = child
	}
	return parent.isWord
}

func dfsVisit(t *Trie, suffixes *[]string, suffix string) {

	if t.isWord {
		suffPtr := reflect.ValueOf(suffixes)
		suff := suffPtr.Elem()
		suff.Set(reflect.Append(suff, reflect.ValueOf(suffix)))
	}

	for ch, trie := range t.root {
		dfsVisit(trie, suffixes, suffix+string(ch))
	}
}

// Search takes a prefix and returns all the word starting with the given prefix
func (t *Trie) Search(prefix string) []string {
	parent := t
	for _, ch := range prefix {
		child, ok := parent.root[ch]
		if !ok {
			return []string{}
		}
		parent = child
	}

	var suffixes []string
	dfsVisit(parent, &suffixes, "")

	for idx, suffix := range suffixes {
		suffixes[idx] = prefix + suffix
	}
	return suffixes
}

// String method to print the trie
func (t *Trie) String() string {
	repr := "["
	for k, v := range t.root {
		repr += string(k) + ": " + v.String()
	}
	repr += "]"
	return repr
}
