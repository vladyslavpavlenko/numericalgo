package iteration

import (
	"math"
)

type Solution struct {
	Root  float64
	Steps [][]float64
}

func Solve(prec float64) (s Solution) {
	x := 50.0
	s.Steps = [][]float64{}

	for {
		xNext := x - f(x)/fPrime(x)
		s.Steps = append(s.Steps, []float64{x, f(x), fPrime(x), xNext})

		if math.Abs(xNext-x) < prec {
			s.Root = xNext
			return s
		}
		x = xNext
	}
}

// f(x) = x^3 + x^2 - 4x - 4
func f(x float64) float64 {
	return math.Pow(x, 3) + math.Pow(x, 2) - 4*x - 4
}

// f'(x) = 3x^2 + 2x - 4
func fPrime(x float64) float64 {
	return 3*math.Pow(x, 2) + 2*x - 4
}
