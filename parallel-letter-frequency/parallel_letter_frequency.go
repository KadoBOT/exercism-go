package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func createFrequency(s string, ch chan FreqMap) {
	ch <- Frequency(s)
}

func ConcurrentFrequency(a []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap)
	for _, s := range a {
		go createFrequency(s, ch)
	}
	for _ = range a {
		chMap := <- ch
		for k, v := range chMap {
			m[k] += v
		}
	}
	return m
}
