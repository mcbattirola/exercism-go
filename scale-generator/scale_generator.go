package scale

import (
	"fmt"
	"strings"
)

// Notes represent an array of notes
type Notes []string

var notesWithSharp = Notes{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var notesWithFlats = Notes{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

// GetNthNote returns the nth note given a initial note and an interval
func (notes Notes) GetNthNote(note string, interval rune) string {
	n := intervalSizes[interval]
	noteIndex, err := notes.GetNoteIndex(note)

	if err != nil {
		panic(err)
	}

	n += noteIndex

	for n > len(notes)-1 {
		n -= len(notes)
	}

	return notes[n]
}

// GetNoteIndex returns the smallest index i at which note == notes[i],
// or -1 if there is no such index.
func (notes Notes) GetNoteIndex(note string) (int, error) {
	for i, n := range notes {
		if note == n {
			return i, nil
		}
	}
	return -1, fmt.Errorf("note not found: %s", note)
}

var intervalSizes = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// Scale generates the musical scale starting with the goven tonic and following the specified interval pattern.
// Intervals are specified using M for major step, m for minor step and A for augmented first
func Scale(tonic string, intervals string) []string {
	notes := getTonicNotes(tonic)
	tonic = strings.Title(tonic)

	if len(intervals) == 0 {
		intervals = "mmmmmmmmmmm"
	}

	scale := []string{tonic}

	currentNote := tonic

	for _, interval := range intervals {
		currentNote = notes.GetNthNote(currentNote, interval)
		if currentNote != tonic {
			scale = append(scale, currentNote)
		}
	}

	return scale
}

func getTonicNotes(tonic string) Notes {
	flatTonics := []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb minor"}

	for _, t := range flatTonics {
		if tonic == t {
			return notesWithFlats
		}
	}

	return notesWithSharp
}
