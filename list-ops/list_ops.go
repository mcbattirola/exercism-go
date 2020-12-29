package listops

type IntList []int
type binFunc func(item, acc int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

// Append appends lists into a single list
func (list IntList) Append(items IntList) IntList {
	if len(list) < 1 && len(items) < 1 {
		return list
	}

	appended := list
	for _, item := range items {
		appended = append(appended, item)
	}

	return appended
}

// Concat adds a list into the original
func (list IntList) Concat(intLists []IntList) IntList {
	if len(list) < 1 && len(intLists) < 1 {
		return list
	}

	newList := list
	for _, intList := range intLists {
		for _, item := range intList {
			newList = append(newList, item)
		}
	}

	return newList
}

// Foldl reduces each item into the accumulator from the left using a given function
func (list IntList) Foldl(f binFunc, accumulator int) int {
	if len(list) < 1 {
		return accumulator
	}

	list[0] = f(accumulator, list[0])

	for i := 0; i < len(list)-1; i++ {
		list[i+1] = f(list[i], list[i+1])
	}

	return list[len(list)-1]
}

// Foldr folds (reduce) each item into the accumulator from the right using a given function
func (list IntList) Foldr(f binFunc, accumulator int) int {
	if len(list) < 1 {
		return accumulator
	}

	list[len(list)-1] = f(list[len(list)-1], accumulator)

	for i := 0; i < len(list)-1; i++ {
		list[i+1] = f(list[i], list[i+1])
	}

	return list[len(list)-1]
}

// Filter returns the list of all items for which predicate(item) is True
func (list IntList) Filter(f predFunc) IntList {
	if len(list) < 1 {
		return list
	}

	var filtered IntList

	for _, item := range list {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

// Map returns the list of the results of applying function(item) on all items
func (list IntList) Map(f unaryFunc) IntList {
	if len(list) < 1 {
		return list
	}

	var mapped IntList

	for _, item := range list {
		mapped = append(mapped, f(item))
	}

	return mapped
}

// Length returns the length of the list
func (list IntList) Length() int {
	return len(list)
}

// Reverse returns the original list in reverse order (last element becomes the first, etc)
func (list IntList) Reverse() IntList {
	if len(list) < 1 {
		return list
	}

	var reverse IntList
	maxIndex := len(list) - 1
	for i := maxIndex; i >= 0; i-- {
		reverse = append(reverse, list[i])
	}

	return reverse
}
