package etl

import "strings"

func Transform(oldMap map[int][]string) map[string]int {
	var newMap = make(map[string]int)

	for k, v := range oldMap {
		for _, letter := range v {
			newMap[strings.ToLower(letter)] = k
		}
	}

	return newMap
}
