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
	"bufio"
	"fmt"
	"io"
	"os"
	str "strings"

	sur "github.com/AlecAivazis/survey/v2"
	"github.com/Josh012422/gocharm/commands"
	"github.com/Josh012422/gocharm/config"
	"github.com/Josh012422/gocharm/misc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filetype string

var cfgFiletype = []*sur.Question{
	{
		Name: "filetype",
		Prompt: &sur.Select{
			Message: "Choose a config file type:",
			Options: []string{
				"YAML",
				"TOML",
				"JSON",
				"HCL",
				"PROPERTIES",
			},
		},
		Validate: sur.Required,
	},
	{
		Name: "set",
		Prompt: &sur.Confirm{
			Message: "Set default timezone now?",
		},
		Validate: sur.Required,
	},
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Creates a config file",
	Long:  `Create a config file if not exists`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(misc.Gray("Creating config file..."))
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.New()

		answers := struct {
			Filetype string
			Set      bool
		}{}

		timezone := ""

		prompted := viper.GetBool("config.created")

		filetype, _ = cmd.Flags().GetString("filetype")

		if filetype == "" && prompted == true {
			filetype = viper.GetString("config.filetype")
		}

		if filetype == "" && prompted != true {

			err := sur.Ask(cfgFiletype, &answers)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if answers.Filetype == "yaml" {
				filetype = "yml"
			}

			viper.Set("config.created", true)
			viper.Set("config.filetype", filetype)
			viper.WriteConfig()

			filetype = str.ToLower(answers.Filetype)

			tz := &sur.Input{
				Message: "Default Timezone:",
			}

			if answers.Set == true {
				_ = sur.AskOne(tz, &timezone)

				viper.Set("default", timezone)
				fmt.Println(misc.Green("✓ Succesfully saved default timezone"))
			}

		}


		success := create.Execute(filetype)

		fmt.Println("Press enter to continue")
		_ = waitForEnter(os.Stdin)

		if success != true {
			fmt.Println(misc.Red("Sorry there was an unexpected error"), create.GetErr())
		}

		viper.Set("config.created", true)
		viper.Set("config.filetype", filetype)
		viper.WriteConfig()

		command.SetFt(filetype)

	},
}

func getFiletype() string {
	return filetype
}

func waitForEnter(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Err()
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringP("filetype", "f", "", "This wil define the type of the config file. (Will prompt if not provided, Default is yaml)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
