// Package proverb is a package to output proverbs
package proverb

import "fmt"

// Proverb generate the relevant proverb given a list of inputs
func Proverb(rhyme []string) []string {
	var proverb []string

	for i, str := range rhyme {
		var verse string
		if i > len(rhyme)-2 {
			verse = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			verse = fmt.Sprintf("For want of a %s the %s was lost.", str, rhyme[i+1])
		}

		proverb = append(proverb, verse)
	}

	return proverb
}
