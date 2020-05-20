package command

import (
	"fmt"
	"os"
//	strc "strconv"
	str "strings"

	"github.com/spf13/viper"
	"github.com/Josh012422/gocharm/misc"
	sur "github.com/AlecAivazis/survey/v2"
)

var Name string
var DefaultTz = viper.GetString("default")
var Id int

func Info (user string) {
	viper.New()
	name := ""
	nameViperExists := viper.Get("users." + user + ".user.name")
	userExistsViper := viper.Get("users." + user)
	nameExists := true

//	fmt.Println(nameViperExists, userExistsViper, "users." + user + ".user.name", "users." + user)

	if userExistsViper == nil {
		fmt.Println(misc.Red("That user does not exists."))
		os.Exit(5)
		return
	}

	if nameViperExists == nil {
		nameExists = false
	}

	if nameViperExists != nil {
		name = viper.GetString("users." + user + ".user.name")
	}

	if nameExists == false {
		fmt.Println(misc.Bold(misc.Red("We do not have a register of your name,")))
		nameQs := &sur.Input{
			Message: "Your name is...",
		}

		_ = sur.AskOne(nameQs, &name)
		Name = name
		viper.Set("users." + user + ".user.name", name)
		viper.WriteConfig()
	}

	var (
		nameInfoPrint = fmt.Sprintf(misc.Bold(misc.Cyan("Your name:")) + " %s\n", name)
		defaultTzInfoPrint = fmt.Sprintf(misc.Bold(misc.Blue("Your default timezone:")) + " %s\n", DefaultTz)
	)

	fmt.Printf("\n%s\n%s\n", nameInfoPrint, defaultTzInfoPrint)
}

func Create () {
	viper.New()
	idRaw := 1
	viper.Set("users_current_number", idRaw)
	idRawViper := viper.GetInt("users_future_number")
	name := ""
	defaultTimeZone := ""

	defaultTzQs := &sur.Input{
		Message: "Default timezone for the new user:",
	}

	nameQs := &sur.Input{
		Message: "Name for the new user:",
	}

	err := sur.AskOne(nameQs, &name)
	_ = sur.AskOne(defaultTzQs, &defaultTimeZone)

	if err != nil {
		fmt.Println(misc.Red("There was an error:"), err)
	}

	Name = name
	Id = idRaw
	viper.Set("users." + name + ".user.current", idRawViper)
	viper.Set("users." + name + ".user.id", idRawViper)
	viper.Set("users." + name + ".user.name", name)
	viper.Set("users." + name + ".user.defaulttimezone", defaultTimeZone)
	viper.Set("user_logged_in", true)
	viper.Set("user_current", name)
	viper.Set("users_current_number", idRaw)
	viper.WriteConfig()
	idRaw += 1
	viper.Set("users_future_number", idRaw)
	viper.WriteConfig()
}

func Change (user string) {
//	fmt.Println(misc.Bold(misc.Yellow("Proccessing...")))
	viper.New()
	currName := viper.GetString("users." + user + ".user.name")
	currDefaultTz := viper.GetString("users." + user + ".user.defaultTimezone")
	userExistsViper := viper.Get("users." + user)
	var newName string
	var confirmed bool
	var setting string
	var newDefaultTz string

	if userExistsViper == nil {
		fmt.Println(misc.Bold(misc.Red("Error: That user does not exists.")))
		fmt.Println(misc.Bold(misc.Red("Error was FATAL, Exiting...")))
		os.Exit(1)
		return
	}

	settingQs := &sur.Select{
		Message: "Choose a setting to modify:",
		Options: []string{
			"NAME",
			"DEFAULT TIMEZONE",
		},
	}

	errPromptRaw := sur.AskOne(settingQs, &setting)
	errPrompt := errPromptRaw.Error()

	if errPrompt == "interrupt" {
		fmt.Println(misc.Bold(misc.Red("Proccess was interruted...")))
		fmt.Println("Exiting...")
		os.Exit(1)
		return
	}

	if errPromptRaw != nil {
		fmt.Println(misc.Bold(misc.Red("There was a error:")), errPromptRaw)
		fmt.Println("Exiting...")
		os.Exit(1)
		return
	}

	confirm := &sur.Confirm{
		Message: "Are you sure of changing " + setting + ":",
	}

	_ = sur.AskOne(confirm, &confirmed)

	if confirmed == true {
		setting = str.ToLower(setting)
		switch setting{
		case "name":
			var confirmedChange bool
			nameQs := &sur.Input{
				Message: "New name for user " + user + ":",
			}

			_ = sur.AskOne(nameQs, &newName)
			newNameTxt := misc.Bold(misc.Green("New name for user " + user + ": " + newName))

			fmt.Printf("\nCurrent Name of " + user + ": %s\n%s\n\n", currName, newNameTxt)
			confirmChange := &sur.Confirm{
				Message: "Are you sure?",
			}

			_ = sur.AskOne(confirmChange, &confirmedChange)

			if confirmedChange == true {
				userLower := str.ToLower(user)
				viper.Set("users." + userLower + ".user.name", newName)
				viper.WriteConfig()

				fmt.Println(misc.Green("✓ Successfully changed " + user + "'s name"))
				return
			} else {
				fmt.Println("Okay bye!")
				return
			}


		case "default timezone":
			var confirmedChange bool
			defaultTzQs := &sur.Input{
				Message: "New default timezone for user " + user + ":",
			}

			_ = sur.AskOne(defaultTzQs, &newDefaultTz)

			newDTzTxt := misc.Bold(misc.Green("New default timezone for " + user + ": " + newDefaultTz))

			fmt.Printf("Current default timezone of " + user + ": %s\n%s\n\n", currDefaultTz, newDTzTxt)

			confirmChange := &sur.Confirm{
				Message: "Are you sure?",
			}

			_ = sur.AskOne(confirmChange, &confirmedChange)

			if confirmedChange == true {
				userLower := str.ToLower(user)
				viper.Set("users." + userLower + ".user.defaulttimezone", newDefaultTz)
				viper.WriteConfig()
				fmt.Println(misc.Bold(misc.Green("✓ Successfully changed " + user + "'s default timezone")))

				return
			} else {
				fmt.Println("Okay bye!")
				return
			}

			return
		}

	}

	return

}

func Use (user string) {
	viper.New()
	userExistsViper := viper.Get("users." + user)
	currUser := viper.GetString("user_current")
	userTxt := misc.Bold(misc.Green("New current user"))
	var confirmed bool

	if userExistsViper == nil {
		fmt.Println(misc.Bold(misc.Red("Error: That user does not exists")))
		os.Exit(1)
		return
	}

	confirmQst := &sur.Confirm{
		Message: "Are you sure of changing user?",
	}


	fmt.Printf("\nCurrent user: " + currUser + "\n%s\n\n", userTxt)

	err := sur.AskOne(confirmQst, &confirmed)

	if err.Error() == "interrupt" {
		fmt.Println(misc.Bold(misc.Red("Proccess was interrupted...")))
		fmt.Println("Exiting...")
		os.Exit(1)
		return
	}

	if err != nil {
		fmt.Println(misc.Bold(misc.Red("Sorry there was a error:")), err)
		fmt.Println("Exiting...")
		os.Exit(1)
		return
	}


	if confirmed == true {
		viper.Set("user_current", user)
		viper.WriteConfig()

		fmt.Println("Current user is now " + user + ".")
		fmt.Println(misc.Bold(misc.Green("✓ Successfully changed current user")))
		os.Exit(0)
		return
	} else {
		fmt.Println("Okay bye!")
	}
}


