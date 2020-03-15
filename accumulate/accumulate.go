package accumulate

func Accumulate(s []string, f func(string) string) []string {
	y := s[:0]
	for _, n := range s {
		y = append(y, f(n))
	}
	return y
}
