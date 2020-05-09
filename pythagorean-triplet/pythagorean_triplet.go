package pythagorean

import "math"

type Triplet struct {
	a, b, c int
}

func Sum(n int) (result []Triplet) {
	for _, triplet := range Range(1, n/2) {
		if triplet.a+triplet.b+triplet.c == n {
			result = append(result, triplet)
		}
	}
	return
}

func Range(min int, max int) (result []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; b < max; b++ {
			sqr := a*a + b*b
			c := int(math.Sqrt(float64(sqr)))
			if sqr == c*c && c <= max {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return
}
