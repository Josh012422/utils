package prompts

import (
	"errors"
	"strings"
	"time"
)

func ValidateEmptyString (input string) error {
	timeAgrees := isValidLocation(input)

	if len (strings.TrimSpace(input)) < 1 {

		return errors.New("The input must not be empty")

	}

	if timeAgrees != nil {
		return errors.New("Not a valid location")
	}

	return nil
}

func isValidLocation (loc string) error {
	_, err := time.LoadLocation(loc)

	if err != nil {
		return err
	}

	return nil
}
