package prompts

import (
	"github.com/manifoldco/promptui"
)

func PromptTimezone1 (name string) (string, error) {
	promptTz1 := promptui.Prompt{
		Label: name,
		Validate: isValidLocation,
	}

	return promptTz1.Run()
}

func PromptTimezone2 (name string) (string, error) {
        promptTz2 := promptui.Prompt{
                Label: name,
                Validate: isValidLocation,
        }

        return promptTz2.Run()
}

func PromptTitle (name string) (string, error) {
	promptT := promptui.Prompt{
		Label: name,
		Validate: EmptyInput,
	}

	return promptT.Run()
}
