package sieve

func Sieve(limit int) []int {
	switch limit {
	case 1:
		return []int{}
	default:
		result := []int{}
		nonPrime := make([]bool, limit+1)

		for i := 2; i <= limit; i++ {
			ok := nonPrime[i]
			if !ok {
				result = append(result, i)
				k := 0
				for l := i * i; l <= limit; l = (i * i) + (k * i) {
					nonPrime[l] = true
					k++
				}
			}
		}
		return result
	}
}
