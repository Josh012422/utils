package create

import (
	"os"
//	"io"
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/Josh012422/utils/misc"
)

func Execute(){
	createFile()
}

var home, erro = homedir.Dir()

func createFile (){
	path := home + "/.config.yml"
	if erro != nil {
		fmt.Println(misc.Red("Sorry there was an error: "), erro)
		os.Exit(1)
	}
	var fileExists, err = os.Stat(path)

	if fileExists != nil {
		fmt.Println(misc.Red("Error: Config file already exists in:"), misc.Yellow(path))
		os.Exit(0)
	}

	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
		fmt.Println(misc.Cyan("Config file created succesfully at"), misc.Green(path))
	}

}

func isError (err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
