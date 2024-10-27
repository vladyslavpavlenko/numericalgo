package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab2/internal/methods"
	"github.com/vladyslavpavlenko/numericalgo/lab2/internal/methods/thomas"
	"os"
	"time"
)

func Thomas() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgGreen(picocolors.Black(" Thomas Algorithm ")))

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution := thomas.Solve()
	dur := time.Since(start)
	s.Stop("Solved!", 0)

	prompts.Note(methods.FormatRoots(solution.Roots), prompts.NoteOptions{Title: "Solution"})

	prompts.Outro(
		fmt.Sprintf("Done in %s ✨",
			picocolors.Cyan(dur.String()),
		),
	)
}
