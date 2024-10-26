package methods

import (
	"fmt"
	"strings"
)

const Precision = 0.001

func FormatRoots(roots []float64) string {
	var sb strings.Builder
	for i, xi := range roots {
		if i == len(roots)-1 {
			sb.WriteString(fmt.Sprintf("x_%d = %.4f", i+1, xi))
			continue
		}
		sb.WriteString(fmt.Sprintf("x_%d = %.4f\n", i+1, xi))
	}

	return sb.String()
}
