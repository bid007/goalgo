package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	var (
		v1 = New(1)
		v2 = New(2)
		v3 = New(3)
		v4 = New(4)
	)

	v1.DirectedEdge(v2)
	v2.DirectedEdge(v3)
	v2.DirectedEdge(v4)
	v2.DirectedEdge(New(10))
	v4.DirectedEdge(v3)
	assert.Equal(t, false, v1.HasCycle())

	v5 := New(5)
	v5.DirectedEdge(v5)
	assert.Equal(t, true, v5.HasCycle())
}

func TestGraphReset(t *testing.T) {
	var (
		v1 = New(1)
		v2 = New(2)
		v3 = New(3)
	)
	v1.color, v2.color, v3.color = BLACK, BLACK, BLACK
	v1.DirectedEdge(v2)
	v2.DirectedEdge(v3)
	v1.Reset()
	assert.Equal(t, v1.color, WHITE)
	assert.Equal(t, v2.color, WHITE)
	assert.Equal(t, v3.color, WHITE)
}
