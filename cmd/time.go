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

	"github.com/Josh012422/cli/misc"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "This command displays the REAL TIME of Colombia",
	Long: `Thid command uses the IANA timezone database, so is the REAL TIME`,
	Run: func(cmd *cobra.Command, args []string) {
		tzFlag,_ := cmd.Flags().GetString("timezone")
		location,_ := time.LoadLocation(tzFlag)

		txt :=misc.Green("The time is: ")
		t := time.Now().In(location)
		t.String()
		fmt.Println(txt +  t.Format("2006-06-02 15:04:05"))
		//fmt.Println(tzFlag)
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.
	timeCmd.Flags().StringP("timezone", "t", "", "The timezone to obtain the time zone. MUST BE: CONTINENT/CITY. MUST BE string with double quotes")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
