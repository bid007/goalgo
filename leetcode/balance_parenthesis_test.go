package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalanceParen(t *testing.T) {
	assert.Equal(t, true, IsParenBalanced("{()(){}}"))
	assert.Equal(t, false, IsParenBalanced("{}()(){}}"))
	assert.Equal(t, true, IsParenBalanced("{}([]())"))
	assert.Equal(t, true, IsParenBalanced(""))
	assert.Equal(t, false, IsParenBalanced("}{)("))
}
