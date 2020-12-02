// Package strand works with DNA and RNA strands
package strand

var complements = map[string]string{
	"A": "U",
	"T": "A",
	"C": "G",
	"G": "C",
}

// ToRNA should return a corresponding RNA strand, given an DNA strand input
func ToRNA(dna string) string {
	rna := ""

	for _, nucleotide := range dna {
		rna += complements[string(nucleotide)]
	}

	return rna
}
