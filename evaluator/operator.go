package evaluator

import (
	"fmt"
	"github.com/avalchev94/boolean-evaluator/stack"
)

type operator struct {
	char     rune
	priority int8
	arity    int
}

var (
	or           = operator{'|', 0, 2}
	and          = operator{'&', 1, 2}
	not          = operator{'!', 2, 1}
	leftBracket  = operator{'(', 3, 0}
	rightBracket = operator{')', 3, 0}
)

func (op operator) greater(other operator) bool {
	return op.priority > other.priority
}

func (op operator) equal(other operator) bool {
	return op.char == other.char
}

func (op operator) calculate(parameters *stack.Stack) error {
	if op.arity <= 0 {
		return fmt.Errorf("operator %c's arity is zero or less", op.char)
	}

	if parameters.Len() < op.arity {
		return fmt.Errorf("insufficient parameters for operator %c", op.char)
	}

	args := make([]bool, op.arity)
	for i := 0; i < op.arity; i++ {
		args[i] = parameters.Pop().(bool)
	}

	switch {
	case op.equal(not):
		parameters.Push(!args[0])
	case op.equal(and):
		parameters.Push(args[0] && args[1])
	case op.equal(or):
		parameters.Push(args[0] || args[1])
	default:
		return fmt.Errorf("missing implementation for operator %c", op.char)
	}

	return nil
}
