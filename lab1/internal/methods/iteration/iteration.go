package iteration

import (
	"math"
)

type Solution struct {
	Root  float64
	Steps [][]float64
}

func Solve(prec float64) (s Solution) {
	x := 3.75
	s.Steps = [][]float64{}
	n := calcN(prec)

	for i := 0; i < n; i++ {
		xNext := phi(x)
		s.Steps = append(s.Steps, []float64{xNext, f(xNext)})
		x = xNext
	}

	if len(s.Steps) != 0 {
		s.Root = s.Steps[len(s.Steps)-1][0]
	}

	return s
}

func calcN(prec float64) int {
	return int((math.Log((5-3)/((1-0.48)*prec)))/(math.Log(1/0.48))) + 1
}

func phi(x float64) float64 {
	return -(16 / math.Pow(x, 2)) + 4/x + 4
}

func f(x float64) float64 {
	return math.Pow(x, 3) - 4*math.Pow(x, 2) - 4*x + 16
}
