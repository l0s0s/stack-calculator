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
	Push(stage int)
	Pop() int
}

func (calc *Calculator) count(symbol string) (int, error) {
	if func() bool {
		for _, o := range calc.operators {
			if string(o) == symbol {
				return true
			}
		}
		return false
	}() {
		arg1, arg2 := calc.stack.Pop(), calc.stack.Pop()

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
	}

	number, err := strconv.Atoi(symbol)
	if err != nil {
		return 0, err
	}

	return number, nil
}

// Calculate return answer to expression.
func (calc *Calculator) Calculate(expression []string) (int, error) {
	for _, symbol := range expression {
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
			number, err := strconv.Atoi(symbol)
			if err != nil {
				return 0, err
			}
			calc.stack.Push(number)
		}
	}

	return calc.stack.Pop(), nil
}
