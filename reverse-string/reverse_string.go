package reverse

import (
	"strings"
)

func Reverse(s string) string {
	var sb strings.Builder

	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}

	return sb.String()
}
