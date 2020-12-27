package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps it takes to reach 1 given an input
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be breater than zero")
	}

	steps := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		steps++
	}

	return steps, nil
}
