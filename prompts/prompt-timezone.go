package prompts

import (
	"github.com/manifoldco/promptui"
)

func PromptTimezone1 (name string) (string, error) {
	promptTz1 := promptui.Prompt{
		Label: name,
		Validate: ValidateEmptyString,
	}

	return promptTz1.Run()
}
