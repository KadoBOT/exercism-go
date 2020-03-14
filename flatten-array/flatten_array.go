package flatten

func recFlatten(input interface{}, c chan interface{}, main bool) {
	for _, item := range input.([]interface{}) {
		switch item.(type) {
		case []interface{}:
			recFlatten(item, c, false)
		case nil:
			continue
		default:
			c <- item
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
