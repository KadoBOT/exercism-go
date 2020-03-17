package pangram

import "strings"

func IsPangram(input string) bool {
	lowerCaseInput := strings.ToLower(input)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(lowerCaseInput, i) {
			return false
		}
	}
	return true
}
