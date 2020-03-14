package reverse

import (
	"strings"
)

func Reverse(s string) string {
	var b strings.Builder
	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteRune(runes[i])
	}
	return b.String()
}
