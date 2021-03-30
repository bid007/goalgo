package stack

import (
	"errors"
)

// Stack is LIFO data structure
type Stack struct {
	data []interface{}
}

// New creates a new instance of stack
func New(size int) *Stack {
	return &Stack{data: make([]interface{}, 0, size)}
}

// Len return the current length of stack
func (s *Stack) Len() int {
	return len(s.data)
}

// Push adds an element to the top of the stack
func (s *Stack) Push(elem interface{}) error {
	if len(s.data) < cap(s.data) {
		s.data = append(s.data, elem)
		return nil
	}
	return errors.New("stack is full")
}

// Pop removes an element from the top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("popping off empty stack")
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

// Peek returns an element from the top of the stack
func (s *Stack) Peek() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("peeking from empty stack")
	}
	return s.data[len(s.data)-1], nil
}
