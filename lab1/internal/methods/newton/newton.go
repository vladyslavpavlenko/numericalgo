package newton

import (
	"math"
)

type Solution struct {
	Root  float64
	Steps [][]float64
}

func Solve(prec float64) (s Solution) {
	x := 2.25
	s.Steps = [][]float64{}
	n := calcN(prec)

	for i := 0; i < n; i++ {
		xNext := x - f(x)/fPrime(x)
		s.Steps = append(s.Steps, []float64{xNext, f(xNext)})
		x = xNext
	}

	s.Root = s.Steps[len(s.Steps)-1][0]
	return s
}

func calcN(prec float64) int {
	return int(math.Log2(((math.Log(0.4/prec))/(math.Log(1/0.311)))+1) + 1)
}

func f(x float64) float64 {
	return math.Pow(x, 3) + math.Pow(x, 2) - 4*x - 4
}

func fPrime(x float64) float64 {
	return 3*math.Pow(x, 2) + 2*x - 4
}
