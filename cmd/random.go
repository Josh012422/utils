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
	"github.com/Josh012422/gocharm/commands"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random -1 {<number 1>} -2 {<number 2>}",
	Short: "To generate a pseudo-random number",
	Long: `This command is here to help you take decisions and generate a random number for whatever you want to, For example:

	$ gocharm random`,
	Run: func(cmd *cobra.Command, args []string) {
		min, _ := cmd.Flags().GetInt("min")
		max, _ := cmd.Flags().GetInt("max")

		command.RanNum(min, max)

	},
}

var randomDecisionCmd = &cobra.Command{
	Use:   "decision -1 {<decision 1>} -2 {<decision 2>}",
	Short: "To help you take decisions",
	Long: `This command is here to help you, undecise person, take decisions. For example:

	$ gocharm random decision -1 {<decision 1>} -2 {<decision>}`,
	Run: func(cmd *cobra.Command, args []string) {
		dec1, _ := cmd.Flags().GetString("one")
		dec2, _ := cmd.Flags().GetString("two")

		command.RanDecision(dec1, dec2)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.AddCommand(randomDecisionCmd)

	randomCmd.Flags().IntP("min", "1", 0, "The minimun number")
	randomCmd.Flags().IntP("max", "2", 99, "The maximun number")
	randomDecisionCmd.Flags().StringP("one", "1", "", "The decision number one")
	randomDecisionCmd.Flags().StringP("two", "2", "", "The decision number two")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
