package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	words := []string{"Hell o", "Hellone", "Ballon", "Cat"}
	trie := New()
	for _, word := range words {
		trie.AddWords(word)
	}

	for _, word := range words {
		assert.True(t, trie.Exists(word))
	}
	assert.False(t, trie.Exists("Hello"))
	assert.False(t, trie.Exists("Ball"))
	assert.False(t, trie.Exists("False"))
}

func TestSearch(t *testing.T) {
	trie := New()
	trie.AddWords("Hell o", "Hellone", "Hello", "Cat")
	prefResult := trie.Search("Hel")
	for _, word := range []string{"Hell o", "Hellone", "Hello"} {
		assert.Contains(t, prefResult, word)
	}
}
