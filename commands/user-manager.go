package command

import (
	"fmt"

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
	nameViperExists := viper.Get(user + ".name")
	nameExists := true

	if nameViperExists == nil {
		nameExists = false
	}

	if nameViperExists != nil {
		name = viper.GetString("users." + user)
	}

	if name == "" && nameExists == false {
		fmt.Println(misc.Bold(misc.Red("We do not have a register of your name,")))
		nameQs := &sur.Input{
			Message: "Your name is...:",
		}

		_ = sur.AskOne(nameQs, &name)
		Name = name
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

	nameQs := &sur.Input{
		Message: "Name for the new user?:",
	}

	err := sur.AskOne(nameQs, &name)

	if err != nil {
		fmt.Println(misc.Red("There was an error:"), err)
	}

	Name = name
	Id = idRaw
	viper.Set("users." + name + ".user.current", idRawViper)
	viper.Set("users." + name + ".user.id", idRawViper)
	viper.Set("user_logged_in", true)
	viper.Set("user_current", name)
	viper.Set("users_current_number", idRaw)
	viper.WriteConfig()
	idRaw += 1
	viper.Set("users_future_number", idRaw)
	viper.WriteConfig()
}
