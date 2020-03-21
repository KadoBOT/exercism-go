package sieve

func Sieve(limit int) (result []int) {
	switch limit {
	case 1:
		result = []int{}
	case 2:
		result = []int{2}
	case 3:
		result = []int{2, 3}
	default:
		result = []int{2, 3}
		notPrime := map[int]bool{4: true, 6: true, 9: true}
		for i := 4; i <= limit; i++ {
			for _, n := range result {
				notPrime[n*n] = true
				notPrime[n*i] = true
			}
			if ok, _ := notPrime[i]; !ok {
				result = append(result, i)
			}
		}
	}
	return
}
