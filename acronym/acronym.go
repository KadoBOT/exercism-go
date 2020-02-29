package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var prev rune
	acronymGenerator := func(r rune) rune {
		defer func() {
			prev = r
		}()
		isDashOrSpace := prev == '-' || unicode.IsSpace(prev)
		isLetterAfterDashOrSpace := isDashOrSpace && unicode.IsLetter(r)
		isUpperLetterAfterSymbol := unicode.IsUpper(r) && !unicode.IsLetter(prev)
		if isUpperLetterAfterSymbol || isLetterAfterDashOrSpace {
			return unicode.ToUpper(r)
		}
		return -1
	}
	return strings.Map(acronymGenerator, s)
}
