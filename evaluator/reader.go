package evaluator

import (
	"fmt"
	"unicode"
)

type reader struct {
	data      []byte
	readIndex int
}

func newReader(expression string) *reader {
	return &reader{
		data:      []byte(expression),
		readIndex: 0,
	}
}

func (r *reader) len() int {
	return len(r.data) - r.readIndex
}

func (r *reader) clear(char rune) error {
	for r.len() > 0 {
		ch, _ := r.read()

		if ch != char {
			r.unread()
			break
		}
	}

	return nil
}

func (r *reader) read() (rune, error) {
	if r.len() <= 0 {
		return -1, fmt.Errorf("reader's data has finished")
	}

	defer func() { r.readIndex++ }()
	return rune(r.data[r.readIndex]), nil
}

func (r *reader) seek() (rune, error) {
	if r.len() <= 0 {
		return -1, fmt.Errorf("reader's data has finished")
	}

	return rune(r.data[r.readIndex]), nil
}

func (r *reader) unread() {
	r.readIndex--
}

func (r *reader) readOperator() (operator, error) {
	ch, err := r.read()
	if err != nil {
		return operator{0, -1, -1}, fmt.Errorf("failed to read rune")
	}

	switch ch {
	case or.char:
		return or, nil
	case and.char:
		return and, nil
	case not.char:
		return not, nil
	case leftBracket.char:
		return leftBracket, nil
	case rightBracket.char:
		return rightBracket, nil
	}

	r.unread()
	return operator{0, -1, -1}, fmt.Errorf("not an operator")
}

func (r *reader) readParameter() (string, error) {
	ch, err := r.seek()
	if err != nil {
		return "", fmt.Errorf("failed to read rune")
	}

	if !unicode.IsLetter(ch) {
		return "", fmt.Errorf("first rune not a letter")
	}

	var parameter string
	for r.len() > 0 {
		ch, err := r.read()
		if err == nil && (unicode.IsLetter(ch) || unicode.IsNumber(ch)) {
			parameter += string(ch)
		} else {
			r.unread()
			break
		}
	}

	return parameter, nil
}
