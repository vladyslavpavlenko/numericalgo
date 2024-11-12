package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab3/internal/methods/newton"
	"os"
	"time"
)

func Newton() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgGreen(picocolors.Black(" Newton Method ")))

	prec, err := getPrecision()
	handleCancel(err)

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution := newton.Solve(prec)
	dur := time.Since(start)
	s.Stop("Solved!", 0)

	prompts.Note(
		fmt.Sprintf("x = %.4f\ny = %.4f", solution.X, solution.Y),
		prompts.NoteOptions{Title: "Solution"},
	)

	prompts.Outro(
		fmt.Sprintf("Done in %s ✨ (%d iterations)",
			picocolors.Cyan(dur.String()),
			solution.Iter,
		),
	)
}
