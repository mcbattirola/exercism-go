package strain

// Ints represent an array of integers
type Ints []int

// Lists represent an array of integer arrays
type Lists [][]int

// Strings represent an array of strings
type Strings []string

// Keep returns a new collection containing elements where the predicate f is true
func (ints Ints) Keep(f func(int) bool) Ints {
	var kepts Ints
	for _, i := range ints {
		if f(i) {
			kepts = append(kepts, i)
		}
	}

	return kepts
}

// Discard returns a new collection containing elements where the predicate f is false
func (ints Ints) Discard(f func(int) bool) Ints {
	var discards Ints
	for _, i := range ints {
		if !f(i) {
			discards = append(discards, i)
		}
	}

	return discards
}

// Keep returns a new collection containing elements where the predicate f is true
func (lists Lists) Keep(f func([]int) bool) Lists {
	var kepts Lists
	for _, i := range lists {
		if f(i) {
			kepts = append(kepts, i)
		}
	}

	return kepts
}

// Keep returns a new collection containing elements where the predicate f is true
func (strings Strings) Keep(f func(string) bool) Strings {
	var discards Strings
	for _, i := range strings {
		if f(i) {
			discards = append(discards, i)
		}
	}

	return discards
}
