package thomas

type Solution struct {
	Roots []float64
}

func Solve() Solution {
	matrix := [][]int{
		{3, 2, 0},
		{2, 4, 1},
		{0, 1, 5},
	}
	rhs := []int{9, 19, 28}

	n := len(matrix)

	a := make([]float64, n) // lower diagonal
	b := make([]float64, n) // main diagonal
	c := make([]float64, n) // upper diagonal
	d := make([]float64, n) // right parts

	for i := 0; i < n; i++ {
		b[i] = float64(matrix[i][i]) // main diagonal
		d[i] = float64(rhs[i])       // right parts

		if i > 0 {
			a[i] = float64(matrix[i][i-1]) // lower diagonal
		}

		if i < n-1 {
			c[i] = float64(matrix[i][i+1]) // upper diagonal
		}
	}

	alpha := make([]float64, n)
	beta := make([]float64, n)

	// forward sweep
	alpha[0] = -c[0] / b[0]
	beta[0] = d[0] / b[0]

	for i := 1; i < n; i++ {
		denom := b[i] + a[i]*alpha[i-1]
		alpha[i] = -c[i] / denom
		beta[i] = (d[i] - a[i]*beta[i-1]) / denom
	}

	// backward sweep
	x := make([]float64, n)
	x[n-1] = beta[n-1]

	for i := n - 2; i >= 0; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i]
	}

	return Solution{
		Roots: x,
	}
}
