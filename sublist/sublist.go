package sublist

import (
	"reflect"
)

type Relation string

func compareSlices(big, small []int) bool {
	i1 := len(small)
	i2 := len(big)
	for i := i1; i <= i2; i++ {
		if reflect.DeepEqual(small, big[i-i1:i]) {
			return true
		}
	}
	return false
}
func Sublist(l1, l2 []int) Relation {
	i1 := len(l1)
	i2 := len(l2)
	result := Relation("unequal")

	switch {
	case i1 < i2:
		ok := compareSlices(l2, l1)
		if ok {
			result = "sublist"
		}
	case i1 > i2:
		ok := compareSlices(l1, l2)
		if ok {
			result = "superlist"
		}
	case i1 == i2:
		if reflect.DeepEqual(l1, l2) {
			result = "equal"
		}
	}

	return result
}
