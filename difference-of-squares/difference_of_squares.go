// Package diffsquares has functins to calculate the difference between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum calculates the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	sumOfSquares := 0

	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}

	return sumOfSquares
}

// Difference calculates the difference between the square of the sum and the sum of the squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
