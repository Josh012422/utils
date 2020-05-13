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
	"github.com/spf13/cobra"
	"github.com/Josh012422/gocharm/commands"
)

// welcomeCmd represents the welcome command
var welcomeCmd = &cobra.Command{
	Use:   "welcome",
	Short: "A command to welcome you.",
	Long: `Hi, I am gocharm and this command will welcome you to this awesome tool and give a guide (If you want) through this CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		command.Welcome()
	},
}

func init() {
	rootCmd.AddCommand(welcomeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// welcomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// welcomeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
