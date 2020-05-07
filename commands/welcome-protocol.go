package command

import (
	"fmt"
	"io"
	"bufio"
	"os"

	"github.com/spf13/viper"
	"github.com/Josh012422/gocharm/misc"
	sur "github.com/AlecAivazis/survey/v2"
)

func Welcome () {
	viper.New()
	name := ""
	guideAccepted := false
	st := os.Stdin
	nameExistsViper := viper.Get("name")
	var nameExists bool

	if nameExistsViper == nil {
		nameExists = false
	} else {
		nameExists = true
	}

	if nameExists == false {
		nameQt := &sur.Input{
			Message: "Hello,",
		}

		_ := sur.AskOne(nameQt, &name)

		fmt.Printf("Hello %s, I am your personal assistant, Gocharm.", name)
		fmt.Println(misc.Yellow("Press enter to continue."))

		_ := waitForEnter(st)

		fmt.Println("I learn new things I can do for you every certain time.")
		fmt Println("I can give you a guide through the things I can do for you right now.")

		confirmGuide := &sur.Confirm{
			Message: "Accept guide?:",
		}

		_ := sur.AskOne(confirmGuide, &guide)

		if guide == true {
			fmt.Println(misc.Cyan("Well, hold tight")

			fmt.Println("First of all, the time command")
			fmt.Println(misc.Yellow("Press enter to continue"))

			_ := waitForEnter(st)

			fmt.Println("When you execute this command, I will give you the exact time of the timezone you desire.")

			fmt.Println(misc.Yellow("Press enter to continue"))
		}

	}

	return


}

func waitForEnter (i io.Reader) error {
	scanner := bufio.NewScanner(i)
	scanner.Scan()
	return scanner.Err()
}
