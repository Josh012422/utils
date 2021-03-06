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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	sur "github.com/AlecAivazis/survey/v2"
	"github.com/Josh012422/gocharm/commands"
	"github.com/Josh012422/gocharm/misc"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocharm",
	Short: "A multiple utilities cli",
	Long: `This cli is made to make your life easier, you can use it from terminal, so if your a dev that is so busy to install an utilities app, then this CLI is for you:

	Example: gocharm time -t "Continent/City"
	Displays the current time`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fileType := getFiletype()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gocharm")
		viper.SetConfigType(fileType)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.ConfigFileUsed()
	}

	cfgFileExists := viper.Get("config.created")
	if cfgFileExists != nil {
		userLoggedIn := viper.GetBool("user_logged_in")

		if userLoggedIn == false {
			var createUserConfirm bool
			fmt.Println(misc.Bold(misc.Red("FATAL: Error: No user found.")))

			createUserQst := &sur.Confirm{
				Message: "Create a new user now?",
			}


			_ = sur.AskOne(createUserQst, &createUserConfirm)
			if createUserConfirm == true {
				command.Create()
				fmt.Println(misc.Bold(misc.Green("✓ Error is now fixed, Continue your action")))
			}

			if createUserConfirm == false {
				fmt.Println(misc.Bold(misc.Red("Error is FATAL, Exiting...")))
				os.Exit(1)
				return
			}
		}
	}
}
