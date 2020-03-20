package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

type Line string

func Encode(input string) string {
	reg, _ := regexp.Compile("[^-a-zA-Z0-9]+")
	str := reg.ReplaceAllString(strings.ToLower(input), "")
	sqr := math.Sqrt(float64(len(str)))
	cols := int(math.Ceil(sqr))
	size := int(math.RoundToEven(sqr))

	result := make([]string, cols)

	for i, v := range str {
		row := i % cols
		col := i / cols
		if col == 0 {
			result[i] = strings.Repeat(" ", size)
		}
		out := []rune(result[row])
		out[col] = v
		result[row] = string(out)
	}

	return strings.Join(result, " ")
}
