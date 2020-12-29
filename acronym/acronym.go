// Package acronym provides functions to create acronyms
package acronym

import (
	"strings"
)

// Abbreviate creates an acronym for a given string
func Abbreviate(s string) string {
	f := func(c rune) bool {
		return c == ' ' || c == '-' || c == '_'
	}

	words := strings.FieldsFunc(s, f)

	acronym := ""
	for _, w := range words {
		if len(w) > 0 {
			acronym += strings.ToUpper(string(w[0]))
		}
	}

	return acronym
}
