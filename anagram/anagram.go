package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	// create a histogram of subject and candidates
	subject = strings.ToLower(subject)
	subjectHistogram := createHistogram(subject)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		if len(candidateLower) != len(subject) {
			continue
		}

		if candidateLower == subject {
			continue
		}

		candidateHistogram := createHistogram(strings.ToLower(candidateLower))

		if !isSameHistogram(subjectHistogram, candidateHistogram) {
			continue
		}

		anagrams = append(anagrams, candidate)
	}

	return anagrams

}

func createHistogram(s string) map[rune]int {
	h := map[rune]int{}
	for _, c := range s {
		if val, ok := h[c]; ok {
			h[c] = val + 1
		} else {
			h[c] = 1
		}
	}

	return h
}

func isSameHistogram(a, b map[rune]int) bool {
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}
