package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab3/internal/methods/jacobi"
	"os"
	"time"
)

func Jacobi() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgYellow(picocolors.Black(" Jacobi Eigenvalue Algorithm ")))

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution := jacobi.Solve()
	dur := time.Since(start)
	s.Stop("Solved!", 0)

	prompts.Note(jacobi.FormatRoots(solution), prompts.NoteOptions{Title: "Solution"})

	prompts.Outro(
		fmt.Sprintf("Done in %s ✨",
			picocolors.Cyan(dur.String()),
		),
	)
}
