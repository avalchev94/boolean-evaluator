package evaluator

import (
	"fmt"

	"github.com/avalchev94/boolean-evaluator/stack"
)

// Evaluator is a struct for boolean evaluation. Parameters is a map that keeps
// all found variables. All keys are set to false by default. Change to effect
// the evaluation.
type Evaluator struct {
	Parameters map[string]bool
	expression string
}

// New creates Evaluator. Beware that New makes only basic validations.
// Evaluation could fail later with better error.
func New(expression string) (*Evaluator, error) {
	parameters := make(map[string]bool)
	brackets := 0

	reader := newReader(expression)
	for reader.len() > 0 {
		reader.clear(' ')

		if param, err := reader.readParameter(); err == nil {
			parameters[param] = false
		} else if op, err := reader.readOperator(); err == nil {
			if op.equal(leftBracket) {
				brackets++
			} else if op.equal(rightBracket) {
				brackets--
			}
		} else {
			ch, _ := reader.seek()
			return nil, fmt.Errorf("unexpected character %c", ch)
		}
	}

	if brackets != 0 {
		return nil, fmt.Errorf("brackets mismatch")
	}

	return &Evaluator{parameters, expression}, nil
}

// Evaluate the expression using the Parameters map.
// If the expression is not valid, an error is returned.
func (e *Evaluator) Evaluate() (bool, error) {
	parameters := stack.New()
	operators := stack.New()

	reader := newReader(e.expression)
	for reader.len() > 0 {
		reader.clear(' ')

		if op, err := reader.readOperator(); err == nil {
			switch {
			case op.equal(and) || op.equal(or):
				if operators.Len() > 0 {
					prevOp := operators.Top().(operator)
					if !prevOp.equal(leftBracket) && !op.greater(prevOp) {
						if err := operators.Pop().(operator).calculate(parameters); err != nil {
							return false, err
						}
					}
				}
				operators.Push(op)

			case op.equal(leftBracket) || op.equal(not):
				operators.Push(op)

			case op.equal(rightBracket):
				for !operators.Empty() && !operators.Top().(operator).equal(leftBracket) {
					if err := operators.Pop().(operator).calculate(parameters); err != nil {
						return false, err
					}
				}
				operators.Pop()
			}
		} else if param, err := reader.readParameter(); err == nil {
			parameters.Push(e.Parameters[param])
		}
	}

	for operators.Len() > 0 {
		if err := operators.Pop().(operator).calculate(parameters); err != nil {
			return false, err
		}
	}

	if parameters.Len() != 1 {
		return false, fmt.Errorf("missing operator?")
	}

	return parameters.Pop().(bool), nil
}
