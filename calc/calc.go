package calc

import (
	"strconv"
)

// Calculator describe dependencies for calculator.
type Calculator struct {
	operators []rune
	stack     stack
}

// NewCalculator return new instance of Calculator.
func NewCalculator(operators []rune, stack stack) *Calculator {
	return &Calculator{
		operators: operators,
		stack:     stack,
	}
}

type stack interface {
	Push(stage float64)
	Pop() float64
}

func (calc *Calculator) count(symbol string) (float64, error) {
	arg2, arg1 := calc.stack.Pop(), calc.stack.Pop()

	switch symbol {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "/":
		return arg1 / arg2, nil
	}

	return 0, InvalidSymbolError(symbol)
}

// Calculate return answer to expression.
func (calc *Calculator) Calculate(expression []string) (float64, error) {
	for _, symbol := range expression {
		if len(symbol) > 1 {
			return 0, InvalidSymbolError(symbol)
		}

		if func() bool {
			for _, o := range calc.operators {
				if string(o) == symbol {
					return true
				}
			}
			return false
		}() {
			result, err := calc.count(symbol)
			if err != nil {
				return 0, err
			}

			calc.stack.Push(result)
		} else {
			number, err := strconv.ParseFloat(symbol, 64)
			if err != nil {
				return 0, err
			}
			calc.stack.Push(number)
		}
	}

	return calc.stack.Pop(), nil
}
