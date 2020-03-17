package pangram

import "strings"

func IsPangram(input string) bool {
	lowerCaseInput := strings.ToLower(input)
	for i := 1; i <= 26; i++ {
		if !strings.Contains(lowerCaseInput, string('a'-1+i)) {
			return false
		}
	}
	return true
}
