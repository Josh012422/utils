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

	"github.com/spf13/cobra"
	"github.com/Josh012422/gocharm/commands"
	"github.com/Josh012422/gocharm/misc"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(misc.Bold(misc.Red("Please use one the following commands: info, change, view, use.")))
	},
}

// TODO: Change temporal placeholders in master branch

var userInfoCmd = &cobra.Command{
	Use: "info",
	Short: "TEMPORAL DEV PLACEHOLDER",
	Long: `TEMPORAL DEV PLACEHOLDER`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("name")

		if user == "" {
			user = args[0]
		}

		command.Info(user)
	},
}

var userCreateCmd = &cobra.Command{
	Use: "create",
	Short: "TEMPORAL DEV PLACEHOLDER",
	Long: `TEMPORAL DEV PLACEHOLDER`,
	Run: func(cmd *cobra.Command, args []string) {
		command.Create()
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userInfoCmd)
	userCmd.AddCommand(userCreateCmd)

	userInfoCmd.Flags().StringP("name", "n", "", "The name of user to retrieve info from")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
