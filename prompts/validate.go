package prompts

import (
	"errors"
	"strings"
)

func ValidateEmptyInput (input string) error {
	if len (strings.TrimSpace(input)) < 1 {
		return errors.New("The input must not be empty")
	}

	return nil
}
