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
	exampleStr := `"America/New_York"`
	cfgFileStr := `
	config:
	  created: true
	  filetype: {config file filetype}
	default: {default timezone you were prompted}
	name: {name}
        tasks:
	  current_number: 1 (This number is based on the number of existing tasks.)
	  future_number: 2 (This one will be the next number for the next task you create.)
	  task:
	    "1":
	      items_current_number: 0 (Same as tasks's current_number but for items)
	      items_future_number: 1 (Same)
	      status: uncompleted (<-- Please dont modify this)
	      title: Do things (The title of the task)
	`
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

			fmt.Printf("A demonstration: \n\n$ gocharm time -t %s -1\n", exampleStr)

			Continue(st)

			fmt.Printf("The %s switch is an instruction for me, to give you the time in 12 hour format.\n", "-1")

			Continue(st)

			fmt.Println(misc.Cyan("Second Command, The date command."))

			fmt.Println("When you execute, I will give you the exact date of the timezone you want")

			Continue(st)

			fmt.Printf("For example: \n\n$ gocharm date -t %s\n", "{<same as time command>}")

			Continue(st)

			fmt.Println(misc.Bold(misc.Cyan("Third command, The tasks command. (This one is long)")))
			fmt.Println("When you summon me with this command, I will request to you use one of the 4 options you have available.")
			Continue(st)

			fmt.Println(misc.Bold(misc.Cyan("Tasks command: The add subcommand.")))

			fmt.Printf("When you summon me with the tasks command and this subcommand, I will add a new task to list: \n\n$ gocharm tasks add\n")

			Continue(st)

			fmt.Println(misc.Bold(misc.Cyan("Tasks command: Add subcommand: Item Sub...subcommand?")))

			fmt.Println("When you summon me with tasks command, add subcommand and item subsubcommand (That sounds weird), I will add a item to the task provided")

			Continue(st)

			fmt.Printf("For example: \n\n$ gocharm tasks add item -t [task number]\n")

			Continue(st)

			fmt.Println(misc.Bold(misc.Cyan("Tasks command: List subcommand.")))
			fmt.Println("When you summon with tasks command and with this command, I will list you all of your pending tasks (And completed ones).")
			Continue(st)
			fmt.Printf("See: \n\n$ gocharm tasks list\n")

			fmt.Println(misc.Bold(misc.Cyan("Tasks command: View subcommand")))
			fmt.Println("When you call me with tasks command and this subcommand, I will bring you all the information of the task provided")
			Continue(st)

			fmt.Printf("Look: \n\n$ gocharm tasks view [task number]\n")

			Continue(st)

			fmt.Println(misc.Bold(misc.Cyan("Tasks command: Complete subcommand.")))

			fmt.Println("When you call me with tasks command and this subcommand (I'm tired of saying that), I will mark as completed the task provided.")

			Continue(st)

			fmt.Printf("Ehmm... Ahhh... Dont know what to say, already used look and see... Well... Look: \n\n$ gocharm tasks complete [task number]\n")

			Continue(st)

			fmt.Println(misc.Bold(misc.Cyan("Tasks command: Complete subcommand: Item... (Oh Noo... Here we go again) Subsubcommand.")))
			fmt.Println("When you... At this point you must know already, I will mark as completed the item of the task provided.")
			Continue(st)

			fmt.Printf("...See: \n\n$ gocharm tasks complete item [task number] [item number]\n")

			Continue(st)

			fmt.Println("Uff, That one was very long, here goes a bonus.")

			Continue(st)

			fmt.Println("To eliminate tasks from the list itself (Not just mark it as completed), you must go to the file where all this tasks is being saved.")
			Continue(st)

			fmt.Printf("There MUST be a file name '.gocharm.yml' (Note: .yml should be replaced by the filetype you chosed when you were prompted), The file should look something like this:\n\n")

			fmt.Printf("%s\n\n(Yeah I now, is very ugly)\nThat syntax is for YAML but applies for the rest.", cfgFileStr)

			Continue(st)

			fmt.Println("To eliminate tasks, just erase all the number of task key and rest one to tasks's current_number and future_number")

		}

		if guideAccepted == false {
			fmt.Printf("As you want %s, bye.\n", nameViper)
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
