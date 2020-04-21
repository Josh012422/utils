/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Josh012422/utils/misc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
)

// newCmd represents the new command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "This command displays the REAL TIME of any city",
	Long: `Thid command uses the IANA timezone database, so is the REAL TIME`,
	Run: func(cmd *cobra.Command, args []string) {
		home, erro := homedir.Dir()
		if erro != nil {
			fmt.Println(erro)
		}

		path := home + "/.config.yml"
		viper.New()
		txt := misc.Green("The time is: ")
		tzFlag,_ := cmd.Flags().GetString("timezone")
		hourFormatFlag,_ := cmd.Flags().GetBool("12hour")
		location,_ := time.LoadLocation(tzFlag)
		defaultTzFlag,_ := cmd.Flags().GetString("default")
		if (defaultTzFlag != "") {
			viper.AddConfigPath("..")
			viper.Set("default", defaultTzFlag)
			viper.WriteConfig()
			fmt.Println(misc.Cyan("Default Timezone succesfully saved in:"), misc.Green(path))
		}

		if (tzFlag != "" && defaultTzFlag == "" && hourFormatFlag == false) {
			t := time.Now().In(location)
			t.String()
			fmt.Println(txt + t.Format("15:04:05 pm"))
		}

		if (tzFlag != "" && defaultTzFlag == "" && hourFormatFlag == true) {
			t := time.Now().In(location)
			t.String()
			fmt.Println(txt + t.Format("3:04:05 pm"))
		}

		if (tzFlag == "" && defaultTzFlag == "" && hourFormatFlag == false) {
			viperLocationRaw := viper.GetString("default")
			viperLocation,_ := time.LoadLocation(viperLocationRaw)

			t := time.Now().In(viperLocation)
			t.String()
			fmt.Println(txt + t.Format(/*2006-06-02*/"15:04:05 pm"))
		}

		if (tzFlag == "" && defaultTzFlag == "" && hourFormatFlag == true) {
			viperLocationRaw := viper.GetString("default")
			viperLocation,_ := time.LoadLocation(viperLocationRaw)

			t := time.Now().In(viperLocation)
			t.String()
			fmt.Println(txt + t.Format(/*2006-06-02*/"3:04:05 pm"))
		}

	},

}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.
	timeCmd.Flags().StringP("timezone", "t", "", "The timezone to obtain the time zone. MUST BE: CONTINENT/CITY. MUST BE string with double quotes.")

	timeCmd.Flags().BoolP("12hour", "1", false, "Displays the time in 12 hour format, if not provided defaults to 24 hour format")

	timeCmd.Flags().StringP("default", "d", "", "This will set the default timezone to he string so you don't have to use -t flag everytime")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
