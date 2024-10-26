package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab2/internal/methods/jacobi"
	"os"
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
	solution := jacobi.Solve(prec)
	dur := time.Since(start)
	s.Stop("Solved!", 0)
	handleCancel(err)

	prompts.Note(solution.FormatRoots(), prompts.NoteOptions{Title: "Solution"})

	prompts.Outro(
		fmt.Sprintf("Done in %s (%d iterations) ✨",
			picocolors.Cyan(dur.String()), solution.Iterations,
		),
	)
}
