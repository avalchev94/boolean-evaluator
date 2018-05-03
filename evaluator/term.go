package evaluator

import (
	"strings"
	"unicode"
)

func isTerm(r rune) bool {
	return unicode.IsLetter(r)
}

func readTerm(reader *strings.Reader) string {

}
