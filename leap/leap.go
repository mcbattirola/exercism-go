// Package leap reports if a given year is a leap year or not.
package leap

// IsLeapYear returns true if year is leap year, false otherwise.
func IsLeapYear(year int) bool {
	if isDivisible(year, 4) {
		if isDivisible(year, 100) {
			if isDivisible(year, 400) {
				return true
			}

			return false
		}

		return true
	}
	return false
}

func isDivisible(numerator int, denominator int) bool {
	return numerator%denominator == 0
}
