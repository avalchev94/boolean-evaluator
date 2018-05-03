package evaluator

import (
	"fmt"

	"github.com/avalchev94/boolean-evaluator/stack"
)

type operator struct {
	char     rune
	priority int8
	arity    int8
}

var (
	or           = &operator{'|', 0, 2}
	and          = &operator{'&', 1, 2}
	not          = &operator{'!', 2, 1}
	leftBracket  = &operator{'(', 3, 0}
	rightBracket = &operator{')', 3, 0}
)

func readOperator(r rune) *operator {
	switch r {
	case or.char:
		fmt.Println("Read operator |")
		return or
	case and.char:
		fmt.Println("Read operator &")
		return and
	case not.char:
		fmt.Println("Read operator !")
		return not
	case leftBracket.char:
		fmt.Println("Read operator (")
		return leftBracket
	case rightBracket.char:
		fmt.Println("Read operator )")
		return rightBracket
	}

	return nil
}

func (this *operator) greater(op *operator) bool {
	return this.priority > op.priority
}

func (this *operator) equal(op *operator) bool {
	return this.char == op.char
}

func (this *operator) calculate(parameters *stack.Stack) (bool, error) {
	fmt.Printf("Calculation for %c\n", this.char)
	if this.arity <= 0 {
		return false, fmt.Errorf("operator arity is zero or less")
	}

	if parameters.Len() < int(this.arity) {
		return false, fmt.Errorf("operator %c does not have enough parameters", this.char)
	}

	var result bool
	switch {
	case this.equal(not):
		result = !parameters.Pop().(bool)
	case this.equal(and):
		result = parameters.Pop().(bool) && parameters.Pop().(bool)
	case this.equal(or):
		result = parameters.Pop().(bool) || parameters.Pop().(bool)
	default:
		return false, fmt.Errorf("operator %c does not have calculate algorithm implemented", this.char)
	}
	fmt.Printf("Calculation end with result %v\n", result)

	return result, nil
}
