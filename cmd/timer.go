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
	"time"
	"os"
	"log"
	str "strconv"

	"github.com/spf13/cobra"
	"github.com/Josh012422/gocharm/misc"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Please provide a number of seconds")
		}

		return nil
	},

	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		duration, err := str.Atoi(args[0])

		if err != nil {
			log.Fatal(err)
		}

		timer := time.NewTimer(time.Duration(duration) * time.Second)
		fmt.Printf("Timer fired!\n%s\n", misc.Red("Press enter to stop it."))

		go func (){
			<-timer.C
			fmt.Printf("\nTimer expired!\n%s", misc.Red("Press enter to continue!"))
			_ = waitForEnter(os.Stdin)
		}()

		_ = waitForEnter(os.Stdin)
		stop := timer.Stop()
		if stop {
			fmt.Printf("Timer stopped!\n%s\n", misc.Red("Exiting..."))
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
