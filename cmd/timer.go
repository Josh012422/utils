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
	"time"
	"os"

	"github.com/spf13/cobra"
	"github.com/Josh012422/gocharm/misc"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "A timer for your baking needs.",
	Long: `A simple timer for your baking needs, Example:
	    
	       $ gocharm timer -s 30 -m 3 -o 3
	       Note: "-o" equals hours, because "-h" means help.`,
	Run: func(cmd *cobra.Command, args []string) {
		var finalDuration int
		seconds, _ := cmd.Flags().GetInt("seconds")
		minutes, _ := cmd.Flags().GetInt("minutes")
		hours, _ := cmd.Flags().GetInt("hours")

		for i := 0; i < 1; i++{
			finalDuration += seconds
		}

		for i := 0; i < minutes; i++{
			finalDuration += 60
		}

		for i := 0; i < hours; i++{
			finalDuration += 3600
		}

		timer := time.NewTimer(time.Duration(finalDuration) * time.Second)
		fmt.Printf("Timer fired!\n%s\n\n" + misc.Bold(misc.Cyan("Seconds: ")) + "%d\n" + misc.Bold(misc.Cyan("Minutes: ")) + "%d\n" + misc.Bold(misc.Cyan("Hours: ")) + "%d\n", misc.Red("Press enter to stop it."), seconds, minutes, hours)

		go func (){
			<-timer.C
			fmt.Printf("\nTimer expired!\n%s", misc.Bold(misc.Red("Press enter to continue!")))
			_ = waitForEnter(os.Stdin)
		}()

		_ = waitForEnter(os.Stdin)
		stop := timer.Stop()
		if stop {
			fmt.Printf("Timer stopped!\n%s\n", misc.Bold(misc.Red("Exiting...")))
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)
	timerCmd.Flags().IntP("seconds", "s", 0, "The number of seconds for the timer")
	timerCmd.Flags().IntP("minutes", "m", 0, "The number of minutes for the timer")
	timerCmd.Flags().IntP("hours", "o", 0, "The number of hours for timer")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
