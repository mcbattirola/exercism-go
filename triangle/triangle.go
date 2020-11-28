// Package triangle implements utility routines for triangles
package triangle

import (
	"math"
)

// Kind of triangles
type Kind string

// Triangle types
const (
	NaT = "Not a Triangle"
	Equ = "Equilateral"
	Iso = "Isosceles"
	Sca = "Scalene"
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene,
// given its sides
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if (a == b) && (b == c) {
		return Equ
	}
	if (a == b) || (a == c) || (b == c) {
		return Iso
	}

	return Sca
}

func isTriangle(a, b, c float64) bool {
	if isAnyInfinite(a, b, c) || isAnyNaN(a, b, c) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if (a+b < c) || (a+c < b) || (b+c < a) {
		return false
	}

	return true
}

func isAnyInfinite(values ...float64) bool {
	for _, value := range values {
		if math.IsInf(value, 0) {
			return true
		}
	}
	return false
}

func isAnyNaN(values ...float64) bool {
	for _, value := range values {
		if math.IsNaN(value) {
			return true
		}
	}
	return false
}
