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

	"github.com/spf13/cobra"
	prom "github.com/manifoldco/promptui"
	"github.com/Josh012422/gocharm/config"
	"github.com/Josh012422/gocharm/misc"
	"github.com/Josh012422/gocharm/main.go"
	"github.com/Josh012422/gocharm/commands"
)

var filetype string;

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Creates a config file",
	Long: `Create a config file if not exists`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(misc.Yellow("Creating config file..."))
	},
	Run: func(cmd *cobra.Command, args []string) {
		var erro error;
		filetype, _ = cmd.Flags().GetString("filetype")
		if filetype == "" {
			promptFt := prom.Select{
				Label: "Config file type",
				Items: []string{"Json", "Toml", "Yaml", "Hcl", "Ini"}
			}

			_, filetype, erro = promptFt.Run()
		}

		success := create.Execute(filetype)

		if success != true {
			fmt.Println(misc.Red("Sorry there was an unexpected error"), create.GetErr())
		}

		commands.setFt(filetype)

	},
}

func getFiletype () string {
	return filetype
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringP("filetype", "f", "yml", "This wil define the type of the config file. (Will prompt if not provided)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
