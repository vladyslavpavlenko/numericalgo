package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab2/internal/methods/jacobi"
	"os"
	"strings"
	"time"
)

func Jacobi() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgMagenta(picocolors.Black(" Jacobi Method ")))

	prec, err := getPrecision()
	handleCancel(err)

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution, iterations := jacobi.Solve(prec)
	dur := time.Since(start)
	s.Stop("Solved!", 0)
	handleCancel(err)

	var sb strings.Builder
	for i, xi := range solution {
		if i == len(solution)-1 {
			sb.WriteString(fmt.Sprintf("x_%d = %.4f", i+1, xi))
			continue
		}
		sb.WriteString(fmt.Sprintf("x_%d = %.4f\n", i+1, xi))
	}

	prompts.Note(sb.String(), prompts.NoteOptions{Title: "Solution"})

	prompts.Outro(
		fmt.Sprintf("Done in %s (%d iterations) ✨",
			picocolors.Cyan(dur.String()), iterations,
		),
	)
}
