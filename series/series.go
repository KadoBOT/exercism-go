package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	var result []string
	for i := 0; i+n <= len(s); i++ {
		result = append(result, s[i:n+i])
	}
	return result
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
