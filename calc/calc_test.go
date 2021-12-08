package calc_test

import (
	"testing"

	"github.com/l0s0s/stack-calculator/calc"
	"github.com/l0s0s/stack-calculator/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalc(t *testing.T) {
	calculator := calc.NewCalculator([]rune("+*"), stack.New())
	result, err := calculator.Calculate([]string{"1", "1", "+", "3", "*"}) // (1 + 1) * 3
	require.NoError(t, err)

	assert.Equal(t, result, 6)
}
