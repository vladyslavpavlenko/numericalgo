package handlers

import (
	"bytes"
	"fmt"
	"github.com/aquasecurity/table"
	"strconv"
)

func makeTable(steps [][]float64) string {
	var buf bytes.Buffer

	t := table.New(&buf)
	t.SetRowLines(false)
	t.SetHeaders("i", "x_n", "f(x_n)", "f'(x_n)", "x_(n+1)")
	t.SetAlignment(table.AlignLeft, table.AlignRight, table.AlignRight, table.AlignRight, table.AlignRight)
	t.SetDividers(table.UnicodeRoundedDividers)
	t.SetLineStyle(table.StyleBrightBlack)
	t.SetHeaderStyle(table.StyleBold)

	for i, step := range steps {
		t.AddRow(
			strconv.Itoa(i+1),
			fmt.Sprintf("%.3f", step[0]),
			fmt.Sprintf("%.3f", step[1]),
			fmt.Sprintf("%.3f", step[2]),
			fmt.Sprintf("%.3f", step[3]),
		)
	}

	t.Render()

	return buf.String()
}
