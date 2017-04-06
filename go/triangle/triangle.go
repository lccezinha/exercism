package triangle

import "math"

const testVersion = 3

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

type Kind string

func discover(sideA, sideB, sideC float64) Kind {
	if notTriangle(sideA, sideB, sideC) {
		return NaT
	}

	if equilateral(sideA, sideB, sideC) {
		return Equ
	}

	if isoseles(sideA, sideB, sideC) {
		return Iso
	}

	if scalene(sideA, sideB, sideC) {
		return Sca
	}

	return NaT
}

func notTriangle(sideA, sideB, sideC float64) bool {
	return (sideA == 0 || sideB == 0 || sideC == 0) ||
		(sideA+sideB < sideC || sideA+sideC < sideB || sideB+sideC < sideA) ||
		(math.IsNaN(sideA) || math.IsNaN(sideB) || math.IsNaN(sideC)) ||
		(math.IsInf(sideA, 1) || math.IsInf(sideB, 1) || math.IsInf(sideC, 1)) ||
		(math.IsInf(sideA, -1) || math.IsInf(sideB, -1) || math.IsInf(sideC, -1))
}

func equilateral(sideA, sideB, sideC float64) bool {
	return sideA == sideB && sideA == sideC
}

func isoseles(sideA, sideB, sideC float64) bool {
	return sideA == sideB || sideB == sideC || sideA == sideC
}

func scalene(sideA, sideB, sideC float64) bool {
	return sideA != sideB && sideB != sideC && sideC != sideA
}

func KindFromSides(sideA, sideB, sideC float64) Kind {
	kind := discover(sideA, sideB, sideC)
	return kind
}
