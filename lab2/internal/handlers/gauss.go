package handlers

import (
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/Mist3rBru/go-clack/third_party/picocolors"
	"github.com/Mist3rBru/go-clack/third_party/sisteransi"
	"github.com/vladyslavpavlenko/numericalgo/lab2/internal/methods"
	gauss "github.com/vladyslavpavlenko/numericalgo/lab2/internal/methods/gauss"
	"os"
	"time"
)

func GaussElimination() {
	os.Stdout.Write([]byte(sisteransi.EraseDown()))
	prompts.Intro(picocolors.BgYellow(picocolors.White(" Gauss Elimination ")))

	prec, err := getPrecision()
	handleCancel(err)

	s := prompts.Spinner(prompts.SpinnerOptions{})
	s.Start("Solving")
	time.Sleep(50 * time.Millisecond)
	start := time.Now()
	solution := gauss.Solve(prec)
	dur := time.Since(start)
	s.Stop("Solved!", 0)
	handleCancel(err)

	prompts.Note(methods.FormatRoots(solution.Roots), prompts.NoteOptions{Title: "Solution"})

	prompts.Note(fmt.Sprintf("%d", solution.Det), prompts.NoteOptions{Title: "Determinant"})

	prompts.Note(gauss.FormatMatrix(solution.Inv), prompts.NoteOptions{Title: "Matrix Inverse"})

	prompts.Outro(
		fmt.Sprintf("Done in %s ✨",
			picocolors.Cyan(dur.String()),
		),
	)
}
