package jacobi

import (
	"fmt"
	"math"
	"strings"
)

func Solve() []float64 {
	A := [][]float64{
		{5, 0, 2, 1},
		{0, 4, 0, 1},
		{2, 0, 2, 0},
		{1, 1, 0, 3},
	}

	n := len(A)

	for iter := 0; iter < 4; iter++ {
		var maximum float64
		var p, q int
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if math.Abs(A[i][j]) > maximum {
					maximum = math.Abs(A[i][j])
					p = i
					q = j
				}
			}
		}

		// angle
		if A[p][p] == A[q][q] {
			theta := math.Pi / 4
			c := math.Cos(theta)
			s := math.Sin(theta)

			// rotation
			for i := 0; i < n; i++ {
				if i != p && i != q {
					Aip := A[i][p]
					Aiq := A[i][q]
					A[i][p] = c*Aip - s*Aiq
					A[p][i] = A[i][p]
					A[i][q] = s*Aip + c*Aiq
					A[q][i] = A[i][q]
				}
			}

			App := A[p][p]
			Aqq := A[q][q]
			Apq := A[p][q]

			A[p][p] = c*c*App - 2*c*s*Apq + s*s*Aqq
			A[q][q] = s*s*App + 2*c*s*Apq + c*c*Aqq
			A[p][q] = 0
			A[q][p] = 0
		} else {
			tau := (A[q][q] - A[p][p]) / (2 * A[p][q])
			t := math.Copysign(1.0/(math.Abs(tau)+math.Sqrt(1+tau*tau)), tau)
			c := 1 / math.Sqrt(1+t*t)
			s := c * t

			// rotation
			for i := 0; i < n; i++ {
				if i != p && i != q {
					Aip := A[i][p]
					Aiq := A[i][q]
					A[i][p] = c*Aip - s*Aiq
					A[p][i] = A[i][p] // Since A is symmetric
					A[i][q] = s*Aip + c*Aiq
					A[q][i] = A[i][q]
				}
			}

			App := A[p][p]
			Aqq := A[q][q]
			Apq := A[p][q]

			A[p][p] = c*c*App - 2*c*s*Apq + s*s*Aqq
			A[q][q] = s*s*App + 2*c*s*Apq + c*c*Aqq
			A[p][q] = 0
			A[q][p] = 0
		}
	}

	eigenvalues := make([]float64, n)
	for i := 0; i < n; i++ {
		eigenvalues[i] = A[i][i]
	}

	return eigenvalues
}

func FormatRoots(roots []float64) string {
	var sb strings.Builder
	for i, xi := range roots {
		if i == len(roots)-1 {
			sb.WriteString(fmt.Sprintf("%.4f", xi))
			continue
		}
		sb.WriteString(fmt.Sprintf("%.4f\n", xi))
	}

	return sb.String()
}
