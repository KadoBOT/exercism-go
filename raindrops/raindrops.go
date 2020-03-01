package raindrops

import (
	"fmt"
	"math"
)

type Num int
func (n Num) hasRemainder(d int) bool {
	fn, fd := float64(n), float64(d)
	return fn >= fd && math.Remainder(fn, fd) == 0
}

func Convert(num int) string  {
	var result string
	n := Num(num)

	if n.hasRemainder(3) {
		result += "Pling"
	}
	if n.hasRemainder(5) {
		result += "Plang"
	}
	if n.hasRemainder(7) {
		result += "Plong"
	}

	if result == ""{
		result = fmt.Sprint(num)
	}
	return result
}
