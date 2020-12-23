package accumulate

// Accumulate performs operation 'f' into every string of the input array
func Accumulate(strs []string, f func(string) string) []string {
	var newStrs []string

	for _, s := range strs {
		newStrs = append(newStrs, f(s))
	}

	return newStrs
}
