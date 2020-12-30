// Package isogram contains functions to handle isograms.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter,
// however spaces and hyphens are allowed to appear multiple times.
package isogram

import (
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(s string) bool {
	letterInString := make(map[rune]bool)

	for _, c := range s {
		if unicode.IsLetter(c) {
			l := unicode.ToLower(c)

			_, ok := letterInString[l]
			if ok {
				return false
			}

			letterInString[l] = true
		}
	}

	return true
}
