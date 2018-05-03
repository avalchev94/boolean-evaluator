package evaluator

import (
	"fmt"
	"strings"

	"github.com/avalchev94/boolean-evaluator/stack"
)

// Evaluator is a struct for boolean evaluation. Parameters is a map that keeps
// all found variables. All keys are set to false by default. Change to effect
// the evaluation.
type Evaluator struct {
	Parameters map[parameter]bool
	expression string
}

// New creates evaluator. Beware that your expression might not be valid, but
// new still will work. However, later the evaluation might fail.
func New(expression string) (*Evaluator, error) {
	parameters := make(map[parameter]bool)

	reader := strings.NewReader(expression)
	for reader.Len() > 0 {
		r, _, err := reader.ReadRune()
		if err != nil {
			return nil, fmt.Errorf("failed to read next rune")
		}

		if param, err := readParameter(r, reader); err == nil {
			parameters[param] = false
		}
	}

	return &Evaluator{parameters, expression}, nil
}

// Evaluate the expression using the Parameters map. If the expression is not valid,
// a verbose errors are returned.
func (e *Evaluator) Evaluate() (bool, error) {
	parameters := stack.New()
	operators := stack.New()

	reader := strings.NewReader(e.expression)
	for reader.Len() > 0 {
		r, i, err := reader.ReadRune()
		if err != nil {
			return false, fmt.Errorf("failed to read next rune")
		}

		if op, err := readOperator(r); err == nil {
			switch {
			case op.equal(and) || op.equal(or):
				if operators.Len() > 0 {
					prevOp := operators.Top().(operator)
					if !prevOp.equal(leftBracket) && !op.greater(prevOp) {
						if err := operators.Pop().(operator).calculate(parameters); err != nil {
							return false, fmt.Errorf("Error on column %d: %s", i, err.Error())
						}
					}
				}
				operators.Push(op)

			case op.equal(leftBracket) || op.equal(not):
				operators.Push(op)
			case op.equal(rightBracket):
				for operators.Len() > 0 && !operators.Top().(operator).equal(leftBracket) {
					if err := operators.Pop().(operator).calculate(parameters); err != nil {
						return false, fmt.Errorf("Error on column %d: %s", i, err.Error())
					}
				}

				if !operators.Top().(operator).equal(leftBracket) {
					return false, fmt.Errorf("Error on column %d: brackets mismatch", i)
				}
				operators.Pop()
			}
		} else if param, err := readParameter(r, reader); err == nil {
			parameters.Push(e.Parameters[param])
		}
	}

	for operators.Len() > 0 {
		if err := operators.Pop().(operator).calculate(parameters); err != nil {
			return false, fmt.Errorf("Parsing failed with error: %s", err.Error())
		}
	}

	if parameters.Len() != 1 {
		return false, fmt.Errorf("Parsing failed. Missing operator?")
	}

	return parameters.Pop().(bool), nil
}
