package leetcode

import (
	"sample/algo/stack"
)

func IsParenBalanced(paren string) bool {
	parenMap := map[interface{}]interface{}{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	s := stack.New(len(paren))
	for _, token := range paren {
		if s.Len() == 0 {
			s.Push(token)
		} else {
			top, _ := s.Peek()
			if token == parenMap[top] {
				s.Pop()
			} else {
				s.Push(token)
			}

		}
	}
	return s.Len() == 0
}
