package triangle

import "math"

type Kind string

const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

func KindFromSides(a, b, c float64) Kind {
	isZeroOrLess := a <= 0 || b <= 0 || c <= 0
	isDegenerate := a+b < c || a+c < b || b+c < a
	isNaN := math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
	isInf := math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)

	if isZeroOrLess || isDegenerate || isNaN || isInf {
		return NaT
	}
	switch {
	case a == b && a == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
