/*
Copyright © 2020 DJBlueSlime

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

	sur "github.com/AlecAivazis/survey/v2"
	"github.com/Josh012422/gocharm/commands"
	"github.com/Josh012422/gocharm/misc"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	str "strconv"
)

// newCmd represents the new command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "This command displays the REAL TIME of any city",
	Long:  `This command uses the IANA timezone database, so is the REAL TIME`,
	Run: func(cmd *cobra.Command, args []string) {
		check := misc.Green("✓")
		currUser := viper.GetString("user_current")
		home, erro := homedir.Dir()
		if erro != nil {
			fmt.Println(erro)
		}

		path := home + "/.gocharm." + command.GetFt()
		viper.New()
		txt := misc.Cyan("The time is: ")

		tzFlag, _ := cmd.Flags().GetString("timezone")
		hourFormatFlag, _ := cmd.Flags().GetBool("12hour")

		location, _ := time.LoadLocation(tzFlag)
		defaultTzFlag, _ := cmd.Flags().GetString("default")

		if defaultTzFlag != "" {
			viper.AddConfigPath("..")
			viper.Set("users."+currUser+".user.default_timezone", defaultTzFlag)
			viper.WriteConfig()

			fmt.Println(check, misc.Cyan("Default Timezone succesfully saved in:"), misc.Green(path))
		}

		if tzFlag != "" && defaultTzFlag == "" && hourFormatFlag == false {
			t := time.Now().In(location)
			t.String()
			fmt.Println(txt + t.Format("15:04:05 pm"))
		}

		if tzFlag != "" && defaultTzFlag == "" && hourFormatFlag == true {
			t := time.Now().In(location)
			t.String()
			fmt.Println(txt + t.Format("3:04:05 pm"))
		}

		if tzFlag == "" && defaultTzFlag == "" && hourFormatFlag == false {
			viperLocationRaw := viper.GetString("users." + currUser + ".user.default_timezone")
			viperLocation, _ := time.LoadLocation(viperLocationRaw)

			t := time.Now().In(viperLocation)
			t.String()
			fmt.Println(txt + t.Format("15:04:05 pm"))
		}

		if tzFlag == "" && defaultTzFlag == "" && hourFormatFlag == true {
			viperLocationRaw := viper.GetString("users." + currUser + ".user.default_timezone")
			viperLocation, _ := time.LoadLocation(viperLocationRaw)

			t := time.Now().In(viperLocation)
			t.String()
			fmt.Println(txt + t.Format("3:04:05 pm"))
		}

	},
}

var timeCompareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Returns two timezones instead of one",
	Long:  `Returns two timezones so you can compare one to each other`,
	Run: func(cmd *cobra.Command, args []string) {
		tz1, _ := cmd.Flags().GetString("timezone")
		tz2, _ := cmd.Flags().GetString("2timezone")
		hff, _ := cmd.Flags().GetBool("12hour")
		tz1sur := ""
		tz2sur := ""
		var (
			tzo1, tzo2 string
			err        error
		)
		var diff int

		if tz1 == "" {

			tz1survey := &sur.Input{
				Message: "Timezone 1:",
			}

			_ = sur.AskOne(tz1survey, &tz1sur)
			tz1 = tz1sur
		}

		if tz2 == "" {
			tz2survey := &sur.Input{
				Message: "Timezone 2:",
			}

			_ = sur.AskOne(tz2survey, &tz2sur)
			tz2 = tz2sur
		}

		tzo1, tzo2, err, diff = command.CompareTime(tz1, tz2, hff)

		if err != nil {
			fmt.Println(err)
		}

		diffStr := str.Itoa(diff)
		fmt.Printf(misc.Bold(misc.Cyan("The time in ")) + misc.Bold(misc.Blue(tz1)) + ": " + tzo1 + " " + "\n" + misc.Bold(misc.Blue("The time in ")) + misc.Bold(misc.Cyan(tz2)) + ": " + tzo2 + "\n")
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.AddCommand(timeCompareCmd)

	// Here you will define your flags and configuration settings.
	timeCmd.Flags().StringP("timezone", "t", "", "The timezone to obtain the time zone. MUST BE: CONTINENT/CITY. MUST BE string with double quotes.")

	timeCmd.Flags().BoolP("12hour", "1", false, "Displays the time in 12 hour format, if not provided defaults to 24 hour format")

	timeCmd.Flags().StringP("default", "d", "", "This will set the default timezone to he string so you don't have to use -t flag everytime")

	timeCompareCmd.Flags().StringP("timezone", "t", "", "The timezone to obtain the time zone. MUST BE: CONTINENT/CITY. MUST BE string with double quotes.")

	timeCompareCmd.Flags().StringP("2timezone", "s", "", "The second timezone to obtain values. MUST BE: CONTINENT/CITY. MUST BE string with double quotes, It is mandatory for this command")

	timeCompareCmd.Flags().BoolP("12hour", "1", false, "Displays the time in 12 hour format, if not provided defaults to 24 hour format")

	timeCompareCmd.Flags().StringP("default", "d", "", "This will set the default timezone to he string so you don't have to use -t flag everytime")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
