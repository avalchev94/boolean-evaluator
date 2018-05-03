package evaluator

import (
	"fmt"
	"strings"
	"unicode"
)

type parameter string

func readParameter(r rune, reader *strings.Reader) (parameter, error) {
	if !unicode.IsLetter(r) {
		return "", fmt.Errorf("first rune not a letter")
	}

	param := parameter(r)
	for {
		r, _, err := reader.ReadRune()
		if err == nil && (unicode.IsLetter(r) || unicode.IsNumber(r)) {
			param += parameter(r)
		} else {
			reader.UnreadRune()
			break
		}
	}

	return param, nil
}
