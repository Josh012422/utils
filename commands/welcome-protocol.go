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
	nameViper := viper.GetString("name")
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

		_ = sur.AskOne(nameQt, &name)

		viper.Set("name", name)
		viper.WriteConfig()
	}

	if nameViper == "" {
		nameViper = name
	}


		fmt.Printf("Hello %s, I am your personal assistant, Gocharm.\n", nameViper)
		fmt.Println(misc.Yellow("Press enter to continue."))

		_ = waitForEnter(st)

		fmt.Println("I learn new things I can do for you every certain time.")
		Continue(st)

		fmt.Println("I can give you a guide through the things I can do for you right now.")

		confirmGuide := &sur.Confirm{
			Message: "Accept guide?:",
		}

		_ = sur.AskOne(confirmGuide, &guideAccepted)

		if guideAccepted == true {
			fmt.Println(misc.Cyan("Well, hold tight"))

			fmt.Println("First of all, the time command.")
			fmt.Println(misc.Yellow("Press enter to continue"))

			_ = waitForEnter(st)

			fmt.Println("When you execute this command, I will give you the exact time of the timezone you desire.")

			fmt.Println(misc.Yellow("Press enter to continue"))

			fmt.Printf("A demonstration: \n\n$ gocharm time -t %s -1\n", "{<string in format Continent/City>}")

			Continue(st)

			fmt.Printf("The %s switch is an instruction for me, to give you the time in 12 hour format.\n", "-1")

			Continue(st)

			fmt.Println(misc.Cyan("Second Command, The date command."))

			fmt.Println("When you execute, I will give you the exact date of the timezone you want")

			Continue(st)

			fmt.Printf("For example: \n\n$ gocharm date -t %s\n", "{<same as time command>}")

			Continue(st)
		}

		if guideAccepted == false {
			fmt.Printf("As you want, %s, bye.\n", nameViper)
		}


	return


}

func waitForEnter (r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Err()
}

func Continue (r io.Reader) error {
	fmt.Println(misc.Yellow("Press enter to continue"))
	err := waitForEnter(r)

	if err != nil {
		return err
	}

	return nil
}
