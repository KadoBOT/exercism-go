package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int, 26)
	for k, v := range input {
		for _, l := range v {
			result[strings.ToLower(l)] = k
		}
	}
	return result
}
