/*
Package raindrops convert a number into a string that contains raindrop sounds corresponding to certain potential factors.

The rules of raindrops are that if a given number:

has 3 as a factor, add 'Pling' to the result.
has 5 as a factor, add 'Plang' to the result.
has 7 as a factor, add 'Plong' to the result.
does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.

*/
package raindrops

import "fmt"

// Convert returns the raindrop sound string from input
func Convert(input int) string {
	raindrop := ""
	if hasFactor(input, 3) {
		raindrop += "Pling"
	}
	if hasFactor(input, 5) {
		raindrop += "Plang"
	}
	if hasFactor(input, 7) {
		raindrop += "Plong"
	}

	if raindrop == "" {
		raindrop = fmt.Sprint(input)
	}

	return raindrop
}

func hasFactor(num int, factor int) bool {
	return num%factor == 0
}
