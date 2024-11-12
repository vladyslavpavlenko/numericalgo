package main

import (
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/vladyslavpavlenko/numericalgo/lab3/internal/handlers"
)

func main() {
	prompt, err := prompts.Select(prompts.SelectParams[string]{
		Message: "Select a method:",
		Options: []*prompts.SelectOption[string]{
			{Label: "Power Iteration", Value: "power"},
			{Label: "Jacobi Eigenvalue Algorithm", Value: "jacobi"},
			{Label: "Newton Method", Value: "newton"},
		},
	})
	if err != nil {
		return
	}

	print("\n")
	switch prompt {
	case "power":
		handlers.Power()
	case "jacobi":
		handlers.Jacobi()
	case "newton":
		handlers.Newton()
	default:
		prompts.Error("Unknown option.")
	}
}
