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
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "This command displays the REAL DATE of any city",
	Long: `Thid command uses the IANA timezone database, so is the REAL TIME`,
	Run: func(cmd *cobra.Command, args []string) {
		home, erro := homedir.Dir()
		if erro != nil {
			fmt.Println(erro)
		}

		path := home + "/.config.yml"
		viper.New()
		txt := misc.Green("The date is: ")
		tzFlag,_ := cmd.Flags().GetString("datezone")
		location,_ := time.LoadLocation(tzFlag)
		defaultTzFlag,_ := cmd.Flags().GetString("default")

		if (defaultTzFlag != "") {
			viper.AddConfigPath("..")
			viper.Set("default", defaultTzFlag)
			viper.WriteConfig()
			fmt.Println(misc.Cyan("Default Timezone succesfully saved in:"), misc.Green(path))
		}

		if (tzFlag != "" && defaultTzFlag == "") {
			t := time.Now().In(location)
			t.String()
			fmt.Println(txt + t.Format("2006-06-02 Monday January"))
		}

		if (tzFlag == "" && defaultTzFlag == "") {
			viperLocationRaw := viper.GetString("default")
			viperLocation,_ := time.LoadLocation(viperLocationRaw)

			t := time.Now().In(viperLocation)
			t.String()
			fmt.Println(txt + t.Format("2006-06-02 Monday January"))
		}


	},

}

func init() {
	rootCmd.AddCommand(dateCmd)

	// Here you will define your flags and configuration settings.
	dateCmd.Flags().StringP("datezone", "z", "", "The timezone to obtain the time zone. MUST BE: CONTINENT/CITY. MUST BE string with double quotes.")

	dateCmd.Flags().StringP("default", "d", "", "This will set the default timezone to the string so you don't have to use -t flag everytime")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
