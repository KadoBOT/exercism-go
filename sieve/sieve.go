package sieve

import "math"

func Sieve(limit int) (result []int) {
	switch limit {
	case 1:
		result = []int{}
	default:
		nonPrime := map[int]bool{}

		for i := 2; i <= limit; i++ {
			ok, _ := nonPrime[i]
			if !ok {
				result = append(result, i)
				k := 0
				j := int(math.Pow(float64(i), 2))
				for l := j; l <= limit; l = j + (k * i) {
					nonPrime[l] = true
					k++
				}
			}
		}
	}
	return
}
