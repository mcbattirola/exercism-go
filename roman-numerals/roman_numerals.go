// Package romannumerals provides functions to convert decimals (arabic) into roman
package romannumerals

import (
	"errors"
	"strings"
)

type conversion struct {
	decimal int
	roman   string
}

var decToRoman = []conversion{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral turns a given number between 1 and 3000 (inclusive) its corresponding roman numeral
func ToRomanNumeral(n int) (string, error) {
	if n > 3000 || n < 1 {
		return "", errors.New("input n must be between 1 and 3000")
	}

	var b strings.Builder

	for _, conversion := range decToRoman {
		for ; n >= conversion.decimal; n -= conversion.decimal {
			b.WriteString(conversion.roman)
		}
	}
	return b.String(), nil
}
