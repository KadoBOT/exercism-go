package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var result int
	computed := make([]bool, limit)
	for _, n := range divisors {
		if n > limit || n == 0 {
			continue
		}

		for j := n; j < limit; j += n {
			if computed[j] {
				continue
			}

			if j%n == 0 {
				computed[j] = true
				result += j
			}
		}
	}

	return result
}
