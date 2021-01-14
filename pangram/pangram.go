// Package pangram finds if a word or sentence is a pangram.
// Pangrams are sentences or words containing every letter in the alphabet
package pangram

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuwyxz"

// IsPangram returns true if the word is a pangram and false otherwise
func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)

	letters := map[rune]bool{}

	for _, c := range sentence {
		letters[c] = true
	}

	for _, c := range alphabet {
		if !letters[c] {
			return false
		}
	}

	return true
}
