package evaluator

type operator struct {
	char     rune
	priority int8
}

var (
	or           = operator{'|', 0}
	and          = operator{'&', 1}
	not          = operator{'!', 2}
	leftBracket  = operator{'(', 3}
	rightBracket = operator{')', 3}
)

func isOperator(r rune) bool {
	return r == and.char || r == or.char || r == not.char
}

func toOperator(r rune) *operator {
  switch r {
  case or.char:
    return &or
  case and.char:
    return &and
  case not.char:
    return &not
  }

  return nil
}

func (left *operator) greater(right* operator) bool {
  return left.priority > right.priority
}
