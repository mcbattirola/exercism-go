// Package dna has utility functions to count DNA strands
package dna

import "errors"

// Histogram represents each items appearence on strand
type Histogram map[rune]int

// DNA represents a DNA strand
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := map[rune]int{
		'A': 0,
		'T': 0,
		'C': 0,
		'G': 0,
	}

	for _, nucleotide := range d {
		if !isValidNucleotide(nucleotide) {
			return h, errors.New("Invalid nucleotide")
		}
		h[nucleotide]++
	}

	return h, nil
}

func isValidNucleotide(nucleotide rune) bool {
	return nucleotide == 'A' || nucleotide == 'T' || nucleotide == 'C' || nucleotide == 'G'
}
