package stack

// Stack contain params which need for stack.
type Stack struct {
	stages []float64
}

// NewStack return new exemplar of stack.
func New() *Stack {
	return &Stack{
		stages: make([]float64, 0),
	}
}

// Push add new stage.
func (s *Stack) Push(stage float64) {
	s.stages = append(s.stages, stage)
}

// Pop return and delete last stage.
func (s *Stack) Pop() float64 {
	if len(s.stages) == 0 {
		return 0
	}

	stage := s.stages[len(s.stages)-1]
	s.stages = s.stages[:len(s.stages)-1]
	return stage
}
