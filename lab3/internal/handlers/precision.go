package handlers

import (
	"errors"
	"fmt"
	"github.com/Mist3rBru/go-clack/prompts"
	"github.com/vladyslavpavlenko/numericalgo/lab3/internal/methods"
	"strconv"
)

func getPrecision() (float64, error) {
	prec, err := prompts.Path(prompts.PathParams{
		Message:      "What should the precision be?",
		InitialValue: strconv.FormatFloat(methods.Precision, 'f', -1, 64),
		Validate: func(value string) error {
			if value == "" {
				return errors.New("precision cannot be empty")
			}

			_, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return errors.New("incorrect value")
			}
			return nil
		},
	})
	if err != nil {
		return 0, fmt.Errorf("error getting precision: %w", err)
	}

	p, err := strconv.ParseFloat(prec, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing precision: %w", err)
	}

	return p, nil
}
