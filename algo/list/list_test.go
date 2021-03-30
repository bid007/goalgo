package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	list := New()
	list.AddToFront(10)
	list.AddToFront(11)
	list.AddToFront(12)
	list.AddToBack(9)
	list.AddToBack(8)
	assert.Equal(t, list.String(), "12->11->10->9->8")
}
