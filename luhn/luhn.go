package luhn

func Valid(s string) bool {
	var count, sum int
	for i := range s {
		v := s[len(s)-i-1]
		if v == ' ' {
			continue
		}

		count++
		if v >= '0' && v <= '9'{
			num := int(v - '0')
			if count % 2 == 0 {
				num *= 2
			}
			if num > 9 {
				num -= 9
			}
			sum += num
			continue
		}

		return false
	}
	return count > 1 && sum % 10 == 0
}

