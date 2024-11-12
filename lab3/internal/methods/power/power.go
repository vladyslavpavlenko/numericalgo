package power

import "math"

func Solve(prec float64) float64 {
	A := [][]float64{
		{5, 0, 2, 1},
		{0, 4, 0, 1},
		{2, 0, 2, 0},
		{1, 1, 0, 3},
	}

	n := len(A)
	b := make([]float64, n)
	for i := range b {
		b[i] = 1.0
	}
	var lambdaOld float64

	for {
		bNext := make([]float64, n)
		for i := 0; i < n; i++ {
			bNext[i] = 0
			for j := 0; j < n; j++ {
				bNext[i] += A[i][j] * b[j]
			}
		}

		lambda := bNext[0]
		for i := 1; i < n; i++ {
			if math.Abs(bNext[i]) > math.Abs(lambda) {
				lambda = bNext[i]
			}
		}

		for i := 0; i < n; i++ {
			bNext[i] /= lambda
		}

		if math.Abs(lambda-lambdaOld) < prec {
			return 1 / lambda
		}

		b = bNext
		lambdaOld = lambda
	}
}
