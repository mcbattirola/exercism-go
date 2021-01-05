package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(s string) string {
	var sb strings.Builder

	runes := utf8.RuneCountInString(s)

	for i := runes - 1; i >= 0; i-- {
		sb.WriteRune([]rune(s)[i])
	}

	return sb.String()
}
