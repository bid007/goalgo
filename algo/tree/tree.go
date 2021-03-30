package tree

import (
	"strconv"
	"strings"
)

// Tree represents a binary tree data structure
type Tree struct {
	val   int
	right *Tree
	left  *Tree
}

// Deserialize deserializes the tree in string format
func (t *Tree) Deserialize() string {
	if t == nil {
		return "<nil>"
	}
	return strconv.Itoa(t.val) + "," + t.left.Deserialize() + "," + t.right.Deserialize()
}

// Serialize takes a deserialized tree string and converts into tree
func Serialize(tStr string) *Tree {
	tokens := strings.Split(tStr, ",")
	var serialize func() *Tree

	serialize = func() *Tree {
		if len(tokens) > 0 {
			token := tokens[0]
			// Handle error here
			val, _ := strconv.Atoi(token)
			tokens = tokens[1:]

			if token == "<nil>" {
				return nil
			}

			t := New(val, nil, nil)
			t.left, t.right = serialize(), serialize()

			return t
		}
		return nil
	}
	return serialize()
}

// New creates a tree instance with given value
func New(val int, left, right *Tree) *Tree {
	return &Tree{
		val:   val,
		left:  left,
		right: right,
	}
}
