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
//	"github.com/spf13/viper"
	"github.com/Josh012422/gocharm/misc"
	"github.com/Josh012422/gocharm/commands"
	"github.com/Josh012422/gocharm/prompts"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(misc.Red("Please use one of the following commands: add, list, view, complete."))
	},
}

var tasksAddCmd = &cobra.Command{
	Use: "add",
	Short: "Temporal Placeholder",
	Long: `Temporal Placeholder`,
	Run: func(cmd *cobra.Command, args []string) {
	//	viper.New()
		title, _ := cmd.Flags().GetString("title")

		if title == "" {
			fmt.Println(misc.Red("Please provide a title: "))
			title, _ = prompts.PromptTitle("Title")
		}

		command.Add(title)

	},
}

var tasksListCmd = &cobra.Command{
	Use: "list",
	Short: "Temporal Placeholder",
	Long: `Temporal Placeholder`,
	Run: func(cmd *cobra.Command, args []string) {
		command.List()
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.AddCommand(tasksAddCmd)
	tasksCmd.AddCommand(tasksListCmd)

	tasksAddCmd.Flags().StringP("title", "t", "", "The title of new task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
