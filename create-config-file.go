package create

import (
	"os"
	"fmt"

	"github.com/Josh012422/gocharm/misc"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func Execute(filetype string) bool {
	success := createFile(filetype)
	if success == true {
		return true
	} else {
		return false
	}
}

var home, erro = homedir.Dir()
var err error

func createFile(filetype string) bool {
	path := home + "/.gocharm." + filetype
	viper.New()
	if erro != nil {
		fmt.Println(misc.Red("Sorry there was an error: "), erro)
		os.Exit(1)
		return true
	}
	var fileExists, err = os.Stat(path)

	if fileExists != nil {
		fmt.Println(misc.Yellow("✗"), misc.Red("Error: Config file already exists in:"), misc.Yellow(path))
		os.Exit(0)
		return true
	}

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return false
		}
		defer file.Close()
		fmt.Println(misc.Cyan("Config file created succesfully at"), misc.Green(path))
		viper.Set("tasks.future_number", 1)
		viper.Set("tasks.current_number", 0)
		os.Exit(0)
		return true
	}
	return true

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func GetErr() error {
	return err
}
