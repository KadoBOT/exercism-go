package scrabble

import (
	"strings"
)

var scores = map[string]int {
	"AEIOULNRST": 1,
	"DG": 2,
	"BCMP": 3,
	"FHVWY": 4,
	"K": 5,
	"JX": 8,
	"QZ": 10,
}

func Score(word string) (result int) {
	for i := range word {
		for k, v := range scores {
			if strings.Contains(k, strings.ToUpper(string(word[i]))) {
				result += v
			}
		}
	}
	return
}
