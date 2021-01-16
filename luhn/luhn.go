// Package luhn performs a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a given input is valid
func Valid(digits string) bool {
	digits = strings.ReplaceAll(digits, " ", "")

	if digits == "0" {
		return false
	}

	shouldDouble := false
	sum := 0

	// iterate from right to left
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == ' ' {
			continue
		}

		digit, err := strconv.Atoi(string(digits[i]))

		if err != nil {
			return false
		}

		if shouldDouble {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		shouldDouble = !shouldDouble
		sum += digit

	}

	return sum%10 == 0
}
