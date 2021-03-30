package stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	elem, _ := s.Pop()
	assert.Equal(t, 3, elem)
}

func TestStackEmptyPop(t *testing.T) {
	s := New(5)
	assert.Equal(t, 0, s.Len())
	_, err := s.Pop()
	assert.ErrorIs(t, err, err, errors.New("popping off empty stack"))
}

func BenchmarkStack(b *testing.B) {
	n := b.N
	s := New(n)

	for i := 0; i < n; i++ {
		s.Push(i)
	}

	for i := 0; i < n; i++ {
		s.Pop()
	}
}
