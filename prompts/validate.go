package prompts

import (
	"errors"
	"strings"
	"time"
)

func EmptyInput (input string) error {

	if len (strings.TrimSpace(input)) < 1 {

		return errors.New("The input must not be empty")

	}

	return nil
}

func isValidLocation (input string) error {
	_, err := time.LoadLocation(input)

	if err != nil {
		return errors.New("Not a valid location")
	}

	return nil
}
