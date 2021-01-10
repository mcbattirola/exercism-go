// Package summultiples finds the sum of all the unique multiples of particular numbers up to but not including that number.
// If we list all the natural numbers below 20 that are multiples of 3 or 5, we get 3, 5, 6, 9, 10, 12, 15, and 18.
// The sum of these multiples is 78.
package summultiples

// SumMultiples find the sum of all the unique multiples of particular numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int) int {
	multiples := map[int]bool{}

	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d == 0 {
				continue
			}

			if i%d == 0 {
				multiples[i] = true
			}
		}
	}

	sum := 0
	for k := range multiples {
		sum += k
	}

	return sum
}
