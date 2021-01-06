// Package protein translate RNA sequences into proteins
package protein

import (
	"errors"
)

var (
	// ErrStop is the error returned when the stop signal is found
	ErrStop = errors.New("stop signal")

	// ErrInvalidBase is the error returned when a invalid base is found
	ErrInvalidBase = errors.New("invalid base")
)

var codonProteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon returns the corresponding protein, given a codon
func FromCodon(codon string) (string, error) {
	if codonProteins[codon] == "STOP" {
		return "", ErrStop
	}

	if codonProteins[codon] == "" {
		return "", ErrInvalidBase
	}

	return codonProteins[codon], nil
}

// FromRNA returns the corresponding list of proteins, given a rna sequence
func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]

		if _, ok := codonProteins[codon]; !ok {
			return proteins, ErrInvalidBase
		}

		if codonProteins[codon] == "STOP" {
			return proteins, nil
		}

		proteins = append(proteins, codonProteins[codon])
	}

	return proteins, nil
}
