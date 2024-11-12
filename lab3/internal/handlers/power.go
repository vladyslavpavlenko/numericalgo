package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab3/internal/methods/power"
	"os"
	"time"
)

func Power() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgBlue(picocolors.White(" Power Iteration ")))

	prec, err := getPrecision()
	handleCancel(err)

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution := power.Solve(prec)
	dur := time.Since(start)
	s.Stop("Solved!", 0)

	prompts.Note(fmt.Sprintf("%.4f", solution), prompts.NoteOptions{Title: "Solution"})

	prompts.Outro(
		fmt.Sprintf("Done in %s ✨",
			picocolors.Cyan(dur.String()),
		),
	)
}
