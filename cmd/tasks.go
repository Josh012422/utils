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
	"errors"

	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
	"github.com/AlecAivazis/survey/v2"
	"github.com/Josh012422/gocharm/misc"
	"github.com/Josh012422/gocharm/commands"
	"github.com/Josh012422/gocharm/prompts"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks {<add> | <list> | <view> | <complete>}",
	Short: "A small and simple task manager",
	Long: `With this simple task manager you can add, list, view and complete tasks to do, For example:

	gocharm tasks add -t "Do something"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(misc.Red("Please use one of the following commands: add, list, view, complete."))
	},
}

var tasksAddCmd = &cobra.Command{
	Use: "add {<title>}",
	Short: "For add tasks to the list",
	Long: `Use this command to add tasks to a list saved in the system`,
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

var tasksAddItemCmd = &cobra.Command{
	Use: "item -c {<content>} -t {<task number>}",
	Short: "Tp",
	Long: `Tp`,
	Run: func(cmd *cobra.Command, args []string) {
		content, _ := cmd.Flags().GetString("content")
		task, _ := cmd.Flags().GetInt("task")


		if content == "" {
			contentSurvey := &survey.Input{
				Message: "Please provide some content:",			}

			_ = survey.AskOne(contentSurvey, &content)
		}


		command.AddItem(content, task)
	},
}

var tasksListCmd = &cobra.Command{
	Use: "list",
	Short: "For list all tasks",
	Long: `Use this command to list all the pending tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		command.List()
	},
}

var tasksViewCmd = &cobra.Command{
	Use: "view {<number>}",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Please provide a task number")
		}

		return nil
	},

	Short: "For view a single task",
	Long: `Use this command to view a single task instead of all`,
	Run: func(cmd *cobra.Command, args []string) {
		command.View(args[0])
	},
}

var tasksCompleteCmd = &cobra.Command{
	Use: "complete {<number>}",
	Args: func(cmd *cobra.Command, args []string) error{
		if len(args) < 1 {
			return errors.New("Pleas provide a number")
		}

		return nil
	},
	Short: "For complete tasks",
	Long: `Use this command to mark tasks as completed. This project use viper for saving the tasks and due to technical limits it can't be deleted so you have to manually delete the task. Usage:

	gocharm tasks complete {<task number>}`,
	Run: func(cmd *cobra.Command, args []string) {
		command.Complete(args[0])
	},
}

var tasksCompleteItemCmd = &cobra.Command{
	Use: "item {<task number>} {<item number>}",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Please provide task and item numbers")
		}

		return nil
	},

	Short: "Tp",
	Long: `Tp`,
	Run: func(cmd *cobra.Command, args []string) {
		taskNumber := args[0]
		itemNumber := args[1]

		command.CompleteItem(taskNumber, itemNumber)
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.AddCommand(tasksAddCmd)
	tasksCmd.AddCommand(tasksListCmd)
	tasksCmd.AddCommand(tasksViewCmd)
	tasksCmd.AddCommand(tasksCompleteCmd)
	tasksAddCmd.AddCommand(tasksAddItemCmd)
	tasksCompleteCmd.AddCommand(tasksCompleteItemCmd)

	tasksAddItemCmd.Flags().IntP("task", "t", 99, "The number of the task to add the item to")

	tasksAddItemCmd.Flags().StringP("content", "c", "", "The content of the new item task")

	tasksAddCmd.Flags().StringP("title", "t", "", "The title of new task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
