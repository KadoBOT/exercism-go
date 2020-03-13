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

	switch {
	case i1 < i2 && compareSlices(l2, l1):
		return "sublist"
	case i1 > i2 && compareSlices(l1, l2):
		return "superlist"
	case i1 == i2 && reflect.DeepEqual(l1, l2):
		return "equal"
	}

	return "unequal"
}
