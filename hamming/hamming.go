/*
Package hamming calculates the Hamming Distance between two strings.

The Hamming distance is only defined for sequences of equal length.
*/
package hamming

import "errors"

// Distance calculates the hamming distance between two strings.
// An attempt to calculate it between sequences of different lengths will return error.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the Hamming distance is only defined for sequences of equal length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
