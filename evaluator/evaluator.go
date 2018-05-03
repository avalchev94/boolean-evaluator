package evaluator

import (
	"fmt"
	"strings"

	"github.com/avalchev94/boolean-evaluator/stack"
)

// Evaluator keeps the needed data for evaluating
type Evaluator struct {
	Parameters map[parameter]bool
	expression string
}

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

func (this *Evaluator) Evaluate() (bool, error) {
	parameters := stack.New()
	operators := stack.New()

	reader := strings.NewReader(this.expression)
	for reader.Len() > 0 {
		r, i, err := reader.ReadRune()
		if err != nil {
			return false, fmt.Errorf("failed to read next rune")
		}

		if op := readOperator(r); op != nil {
			switch {
			case op.equal(not):
				b, err := op.calculate(parameters)
				if err != nil {
					return false, fmt.Errorf("Error on column %d: %s", i, err.Error())
				}
				parameters.Push(b)

			case op.equal(and) || op.equal(or):
				if operators.Len() > 0 {
					prevOp := operators.Top().(*operator)
					if !prevOp.equal(leftBracket) && !op.greater(prevOp) {
						b, err := operators.Pop().(*operator).calculate(parameters)
						if err != nil {
							return false, fmt.Errorf("Error on column %d: %s", i, err.Error())
						}
						parameters.Push(b)
					}
				}
				operators.Push(op)

			case op.equal(leftBracket):
				operators.Push(op)
			case op.equal(rightBracket):
				for operators.Top().(*operator) != leftBracket && operators.Len() > 0 {
					b, err := operators.Pop().(*operator).calculate(parameters)
					if err != nil {
						return false, fmt.Errorf("Error on column %d: %s", i, err.Error())
					}
					parameters.Push(b)
				}

				if operators.Top().(*operator) != leftBracket {
					return false, fmt.Errorf("Brackets mismatch on column %d.", i)
				}
				operators.Pop()
			}
		} else if param, err := readParameter(r, reader); err == nil {
			parameters.Push(this.Parameters[param])
		}
	}

	for operators.Len() > 0 {
		b, err := operators.Pop().(*operator).calculate(parameters)
		if err != nil {
			return false, fmt.Errorf("Parsing failed with error: %s", err.Error())
		}
		parameters.Push(b)
	}

	return parameters.Pop().(bool), nil
}
