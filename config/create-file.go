package create

import (
	"os"
//	"io"
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/Josh012422/gocharm/misc"
)

func Execute (filetype string) bool {
	success := createFile(filetype)
	if success == true {
		return true
	} else {
	    return false
	}
}

var home, erro = homedir.Dir()
var err error;

func createFile (filetype string) bool {
	path := home + "/.config." + filetype
	if erro != nil {
		fmt.Println(misc.Red("Sorry there was an error: "), erro)
		return true
		os.Exit(1)
	}
	var fileExists, err = os.Stat(path)

	if fileExists != nil {
		fmt.Println(misc.Yellow("✗"), misc.Red("Error: Config file already exists in:"), misc.Yellow(path))
		return true
		os.Exit(0)
	}

	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if isError(err) {
			return false
		}
		defer file.Close()
		fmt.Println(misc.Cyan("Config file created succesfully at"), misc.Green(path))
		return true
		os.Exit(0)
	}
	return true

}

func isError (err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func GetErr () error {
	return err
}
