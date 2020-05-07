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
	"os"
	"time"
	//	"io"
	//	"bufio"

	"github.com/Josh012422/gocharm/misc"
	"github.com/spf13/cobra"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "A very simple and small cronometer",
	Long: `A very simple and small cronometer, just to help you measure time while in terminal:

	$gocharm cron
	Cronometer running...
	Hit enter to stop.
	// 1 minute and 30 seconds later
	Time Elapsed: 1m30.274817462s`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		fmt.Printf("Cronometer running...\n%s\n", misc.Red("Hit enter to stop."))
		error1 := waitForEnter(os.Stdin)
		if error1 != nil {
			fmt.Println(error1)
			os.Exit(1)
		}
		t := time.Now()
		t.String()
		elapsed := t.Sub(start)

		fmt.Printf("%s %s\n", misc.Bold("Time elapsed: "), elapsed)
	},
}

/*func waitForEnter (r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Err()
}*/

func init() {
	rootCmd.AddCommand(cronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
