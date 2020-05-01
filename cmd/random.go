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
//	"fmt"
//	"math/rand"

	"github.com/spf13/cobra"
///	"github.com/Josh012422/gocharm/misc"
	"github.com/Josh012422/gocharm/commands"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		min, _ := cmd.Flags().GetInt("min")
		max, _ := cmd.Flags().GetInt("max")

		command.RanNum(min, max)
	},
}

var randomDecisionCmd = &cobra.Command{
	Use: "decision -1 {<decision 1>} -2 {<decision 2>}",
/*	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Please provide decisions")
		}

		return nil
	},*/
	Short: "TP",
	Long: `TP`,
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
