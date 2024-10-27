package main

import (
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/vladyslavpavlenko/numericalgo/lab2/internal/handlers"
)

func main() {
	prompt, err := prompts.Select(prompts.SelectParams[string]{
		Message: "Select a method:",
		Options: []*prompts.SelectOption[string]{
			{Label: "Gauss Elimination", Value: "gauss"},
			{Label: "Thomas Algorithm", Value: "thomas"},
			{Label: "Jacobi Method", Value: "jacobi"},
		},
	})
	if err != nil {
		return
	}

	print("\n")
	switch prompt {
	case "jacobi":
		handlers.Jacobi()
	case "gauss":
		handlers.Gauss()
	case "thomas":
		handlers.Thomas()
	default:
		prompts.Error("Unknown option.")
	}
}
