package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	uf := NewSet(1, 2, 3, 4, 5, 6, 7)
	assert.Equal(t, 7, uf.Find(7))

	uf.Union(7, 5)
	assert.Equal(t, 5, uf.Find(7))
	assert.Equal(t, 2, uf.size[5])

	uf.Union(5, 3)
	assert.Equal(t, 5, uf.Find(3))
	assert.Equal(t, 3, uf.size[5])
	assert.Equal(t, 1, uf.size[3])
}
