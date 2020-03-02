package reverse

import "strings"

func Reverse(s string) string {
	result := make([]string, len(s))
	for i, v := range s {
		result[len(s)-i-1] = string(v)
	}
	return strings.Join(result, "")
}
