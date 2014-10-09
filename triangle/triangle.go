package triangle

import (
	"math"
)

const (
	Equ = "Equilateral"
	Iso = "Isosceles"
	Sca = "Scalene"
	NaT = "NotATriangle"
)

type Kind string

func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if (a+b <= c) || (b+c <= a) || (a+c <= b) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
