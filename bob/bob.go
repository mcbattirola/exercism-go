// Package bob simulates a conversation with a teenager
package bob

import (
	"strings"
	"unicode"
)

// Hey returns a lackadaisical response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if len(remark) > 1 {
		isQuestion := remark[len(remark)-1] == '?'
		isYell := isAllCaps(remark) && isAnyLower(remark)

		if isQuestion {
			if isYell {
				return "Calm down, I know what I'm doing!"
			}
			return "Sure."
		}

		if isYell {
			return "Whoa, chill out!"
		}

		return "Whatever."
	}

	return "Fine. Be that way!"
}

func isAllCaps(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isAnyLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
