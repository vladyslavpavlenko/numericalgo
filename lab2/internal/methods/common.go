package methods

import (
	"fmt"
	"strings"
)

const Precision = 0.001

type Solution struct {
	Roots      []float64
	Iterations int
}

func (s *Solution) FormatRoots() string {
	var sb strings.Builder
	for i, xi := range s.Roots {
		if i == len(s.Roots)-1 {
			sb.WriteString(fmt.Sprintf("x_%d = %.4f", i+1, xi))
			continue
		}
		sb.WriteString(fmt.Sprintf("x_%d = %.4f\n", i+1, xi))
	}

	return sb.String()
}
