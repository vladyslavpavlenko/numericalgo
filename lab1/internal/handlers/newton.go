package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab1/internal/methods/newton"
	"os"
	"time"
)

func HandleNewtonsMethod() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgYellow(picocolors.Black(" Newton's Method ")))

	prec, err := getPrecision()
	handleCancel(err)

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution := newton.Solve(prec)
	dur := time.Since(start)
	s.Stop("Solved!", 0)
	handleCancel(err)

	p, err := prompts.Confirm(prompts.ConfirmParams{
		Message:      "Print table?",
		InitialValue: true,
	})
	handleCancel(err)

	if p {
		prompts.Note(makeTable(solution.Steps), prompts.NoteOptions{Title: "Table"})
	}

	prompts.Outro(fmt.Sprintf("Root: %s\nIter: %s\nTime: %s",
		picocolors.Cyan(fmt.Sprintf("%.3f", solution.Root)),
		picocolors.Cyan(fmt.Sprintf("%d", len(solution.Steps))),
		picocolors.Cyan(dur.String()),
	),
	)
}
