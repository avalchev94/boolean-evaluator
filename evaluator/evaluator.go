package evaluator

import (
	"fmt"
	"github.com/avalchev94/boolean-evaluator/stack"
	"github.com/avalchev94/boolean-evaluator/tree"
	"strings"
	"unicode"
)

// Evaluator keeps the needed data for evaluating
type Evaluator struct {
	tree *tree.Tree
}

func findBracketMatch(expression string) int {
	brackets := 1

	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case lbracket:
			brackets++
		case rbracket:
			brackets--
		}

		if brackets == 0 {
			return i
		} else if brackets < 0 {
			return -1
		}
	}

	return -1
}

func toTree(expression string) (*tree.Tree, error) {
	termStack := stack.New()
  operatorStack := stack.New()

	reader := strings.NewReader(expression)
	for reader.Len() > 0 {
		r, i, err := reader.ReadRune()
		if err != nil {
			break
		}

		if op := toOperator(r); op != nil {
      switch operator {
      case not:
        if termStack.Len() <= 0 {
					return nil, fmt.Errorf("Operator \"not\" on column %d does not have term.", i)
				}
        newTree := tree.New(not.char)
        newTree.SetChild(termStack.Pop().(*tree.Tree))
				term := termStack.Pop().(*tree.Tree)
				treeStack.Push(tree.New(not.char).SetChild(term))

      case and, or:
        if op.greater(operatorStack.Top().(*operator)) {
          if termStack.Len() <= 1 {
            return nil, fmt.Errorf("Operator %s on column %d does not have enough terms.", op.char, i)
          }

        }
      }
			if r == not.char {


		} else if isTerm(r) {
			termStack.Push(readTerm(reader))
		}
	}
}

func New(expression string) (*Evaluator, error) {
	tree, err := expressionToTree(expression)
	return &Evaluator{tree}, err
}
