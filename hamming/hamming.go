package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNAs have different size")
	}

	var dif int
	for i := range a {
		if a[i] != b[i] {
			dif++
		}
	}
	return dif, nil
}
