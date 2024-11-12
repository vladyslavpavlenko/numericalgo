package newton

import (
	"math"
)

type Solution struct {
	X    float64
	Y    float64
	Iter int
}

func f(x, y float64) float64 {
	return 5*x - 6*y + 20*math.Log10(x) + 16
}

func g(x, y float64) float64 {
	return 2*x + y - 10*math.Log10(y) - 4
}

func fx(x float64) float64 {
	return 5 + 20/(x*math.Log(10))
}

func fy() float64 {
	return -6
}

func gx() float64 {
	return 2
}

func gy(y float64) float64 {
	return 1 - 10/(y*math.Log(10))
}

func Solve(prec float64) *Solution {
	x, y := 1.5, 2.5
	maxIter := 5
	for i := 0; i < maxIter; i++ {
		F := []float64{f(x, y), g(x, y)}

		// Jacobian matrix
		J := [][]float64{
			{fx(x), fy()},
			{gx(), gy(y)},
		}

		detJ := J[0][0]*J[1][1] - J[0][1]*J[1][0]
		if detJ == 0 {
			return &Solution{0, 0, i}
		}

		deltaX := (-F[0]*J[1][1] + F[1]*J[0][1]) / detJ
		deltaY := (-F[1]*J[0][0] + F[0]*J[1][0]) / detJ

		x += deltaX
		y += deltaY

		if math.Sqrt(deltaX*deltaX+deltaY*deltaY) < prec {
			return &Solution{x, y, i + 1}
		}
	}

	return &Solution{0, 0, 0}
}
