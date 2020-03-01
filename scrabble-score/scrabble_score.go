package scrabble

import (
	"strings"
)

func getValue(s uint8) int {
	letter := strings.ToUpper(string(s))
	switch {
	case strings.Contains("AEIOULNRST", letter):
		return 1
	case strings.Contains("DG", letter):
		return 2
	case strings.Contains("BCMP", letter):
		return 3
	case strings.Contains("FHVWY", letter):
		return 4
	case strings.Contains("K", letter):
		return 5
	case strings.Contains("JX", letter):
		return 8
	case strings.Contains("QZ", letter):
		return 10
	default:
		return 0
	}
}
func Score(word string) (result int) {
	for i := range word {
		value := getValue(word[i])
		result += value
	}
	return
}
