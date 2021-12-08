package stack_test

import (
	"testing"

	"github.com/l0s0s/stack-calculator/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	testStack := stack.New()
	testStack.Push(0)
	assert.Equal(t, testStack.Pop(), float64(0))
	testStack.Pop()
	assert.Equal(t, testStack.Pop(), float64(0))
}
