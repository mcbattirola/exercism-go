// Package scrabble helps with scrabble computations
package scrabble

import "strings"

// Score provides a scrabble score for a given word
func Score(s string) int {
	points := 0

	sUpperCased := strings.ToUpper(s)
	for _, c := range sUpperCased {
		points += getRunePoints(c)
	}

	return points

}

func getRunePoints(c rune) int {
	if strings.ContainsRune("AEIOULNRST", c) {
		return 1
	}

	if strings.ContainsRune("DG", c) {
		return 2
	}

	if strings.ContainsRune("BCMP", c) {
		return 3
	}

	if strings.ContainsRune("FHVWY", c) {
		return 4
	}

	if strings.ContainsRune("K", c) {
		return 5
	}

	if strings.ContainsRune("JX", c) {
		return 8
	}

	if strings.ContainsRune("QZ", c) {
		return 10
	}

	return 0

}
