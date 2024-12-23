package jacobi

import (
	"math"
)

type Solution struct {
	Roots []float64
	Iter  int
}

func Solve(prec float64) Solution {
	a := [][]float64{
		{6, 0, 2, 3},
		{0, 4, 2, 1},
		{2, 2, 5, 0},
		{1, 1, 0, 3},
	}
	b := []float64{24, 18, 21, 15}

	x := []float64{0, 0, 0, 0}
	n := len(x)
	tmpX := make([]float64, n)
	maxDiff := 0.0
	iter := 0

	for {
		maxDiff = 0
		copy(tmpX, b)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					tmpX[i] -= a[i][j] * x[j]
				}
			}

			xUpdated := tmpX[i] / a[i][i]
			diff := math.Abs(x[i] - xUpdated)
			x[i] = xUpdated

			if diff > maxDiff {
				maxDiff = diff
			}
		}
		iter++

		if maxDiff <= prec {
			break
		}
	}

	return Solution{
		Roots: x,
		Iter:  iter,
	}
}
