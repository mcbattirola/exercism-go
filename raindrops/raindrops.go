/*
Package raindrops convert a number into a string that contains raindrop sounds corresponding to certain potential factors.

The rules of raindrops are that if a given number:

has 3 as a factor, add 'Pling' to the result.
has 5 as a factor, add 'Plang' to the result.
has 7 as a factor, add 'Plong' to the result.
does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.*/
package raindrops

import "strconv"

// Convert returns the raindrop sound string from input
func Convert(input int) string {
	raindrop := ""

	if input%3 == 0 {
		raindrop += "Pling"
	}
	if input%5 == 0 {
		raindrop += "Plang"
	}
	if input%7 == 0 {
		raindrop += "Plong"
	}

	if raindrop == "" {
		return strconv.Itoa(input)
	}

	return raindrop
}
