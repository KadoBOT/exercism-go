package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(s string) bool {
	trimmedStr := strings.TrimSpace(s)
	if len(trimmedStr) <= 1 {
		return false
	}

	out := []rune(trimmedStr)
	var count int
	var sum int
	for i := range out {
		v := out[len(out)-i-1]
		if unicode.IsSpace(v) {
			continue
		}

		count++
		num, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		if count % 2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	return sum % 10 == 0
}
