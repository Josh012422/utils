package command

import "fmt"
import "github.com/spf13/viper"

var FT string

func SetFt(filetype string) {
	viper.New()
	viper.Set("config.filetype", filetype)
	viper.Set("config.created", true)
	viper.WriteConfig()
	filetype = FT
	fmt.Println(filetype, FT)
}

func GetFt() string {
	viper.New()
	FT = viper.GetString("config.filetype")
	return FT
}
