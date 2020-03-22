package secret

func reverse(result []string) []string {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func Handshake(code uint) []string {
	var result []string
	if 1<<0&code > 0 {
		result = append(result, "wink")
	}
	if 1<<1&code > 0 {
		result = append(result, "double blink")
	}
	if 1<<2&code > 0 {
		result = append(result, "close your eyes")
	}
	if 1<<3&code > 0 {
		result = append(result, "jump")
	}
	if 1<<4&code > 0 {
		result = reverse(result)
	}

	return result
}
