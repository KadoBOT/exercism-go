package flatten

func recFlatten(input interface{}, c chan interface{}, main bool) {
	switch t := input.(type) {
	case []interface{}:
		for _, item := range t {
			switch item.(type) {
			case []interface{}:
				recFlatten(item, c, false)
			case int:
				c <- item
			default:
				continue
			}
		}
	}

	if main {
		close(c)
	}
}
func Flatten(input interface{}) []interface{} {
	c := make(chan interface{})
	go recFlatten(input, c, true)

	result := []interface{}{}
	for n := range c {
		result = append(result, n)
	}

	return result
}
