package gauss

import (
	"fmt"
	"math"
	"strings"
)

type Solution struct {
	Roots []float64
	Det   int
	Inv   [][]float64
}

func Solve(prec float64) Solution {
	A := [][]float64{
		{1, 2, 3, 0},
		{4, 3, 1, 2},
		{2, 1, 2, 1},
		{0, 3, 0, -5},
	}
	b := []float64{22, 30, 21, -21}

	n := len(A)
	invA := identityMatrix(n)
	det := 1.0

	for i := 0; i < n; i++ {
		if A[i][i] == 0 {
			for j := i + 1; j < n; j++ {
				if A[j][i] != 0 {
					A[i], A[j] = A[j], A[i]
					b[i], b[j] = b[j], b[i]
					invA[i], invA[j] = invA[j], invA[i]
					det = -det
					break
				}
			}
		}

		pivot := A[i][i]
		if pivot == 0 {
			return Solution{Roots: nil, Det: 0, Inv: nil}
		}
		det *= pivot
		for j := 0; j < n; j++ {
			A[i][j] /= pivot
			invA[i][j] /= pivot
		}
		b[i] /= pivot

		for j := 0; j < n; j++ {
			if j != i {
				factor := A[j][i]
				for k := 0; k < n; k++ {
					A[j][k] -= factor * A[i][k]
					invA[j][k] -= factor * invA[i][k]
				}
				b[j] -= factor * b[i]
			}
		}
	}

	roots := make([]float64, n)
	for i := 0; i < n; i++ {
		roots[i] = round(b[i], prec)
	}

	return Solution{
		Roots: roots,
		Det:   int(round(det, 0)),
		Inv:   invA,
	}
}

func identityMatrix(n int) [][]float64 {
	identity := make([][]float64, n)
	for i := 0; i < n; i++ {
		identity[i] = make([]float64, n)
		identity[i][i] = 1
	}
	return identity
}

func round(x, prec float64) float64 {
	pow := math.Pow(10, prec)
	return math.Round(x*pow) / pow
}

func FormatMatrix(matrix [][]float64) string {
	sb := strings.Builder{}
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			sb.WriteString(fmt.Sprintf("%8.4f", matrix[row][column]))
			if column < len(matrix[0])-1 {
				sb.WriteString(" ")
			}
		}
		if row < len(matrix)-1 {
			sb.WriteString(" \n")
		}
	}
	return sb.String()
}
