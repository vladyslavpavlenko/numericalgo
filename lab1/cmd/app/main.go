package main

import (
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/vladyslavpavlenko/numericalgo/lab1/internal/handlers"
)

func main() {
	prompt, err := prompts.Select(prompts.SelectParams[string]{
		Message: "Select a task:",
		Options: []*prompts.SelectOption[string]{
			{Label: "x^3 + x^2 - 4x - 4 = 0 (Newton's Method)", Value: "newton"},
			{Label: "x^3 - 4x^2 - 4x + 16 = 0 (Fixed Point Iteration Method)", Value: "iteration"},
		},
	})
	if err != nil {
		return
	}

	print("\n")
	switch prompt {
	case "newton":
		handlers.HandleNewtonsMethod()
	case "iteration":
		handlers.HandleFixedPointIterationMethod()
	default:
		prompts.Error("Unknown prompt.")
	}
}
