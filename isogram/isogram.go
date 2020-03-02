package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for _, c := range s {
		if unicode.IsLetter(c) && strings.Count(s, string(c)) > 1{
			return false
		}
	}
	return true
}
