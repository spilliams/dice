// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spilliams/dice/cmd/matchers"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "roll NdM",
	Short: "Roll some dice!",
	Args:  cobra.MinimumNArgs(1),
	Long: `Roll some dice! For example:

	$ roll 2d4

would roll (2) 4-sided dice. Other valid configurations:

	Input			Description
	1d20+4			(1) 20-sided die, with a modifier of 4 added
	4d6 drop lowest	(4) 6-sided dice, with the lowest value dropped`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// combine all args into one string
		allArgs := strings.Join(args, "")
		result, err := roll(allArgs)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	},
}

func roll(input string) (int, error) {
	for _, m := range matchers.All() {
		if m.Matches(input) {
			return m.Run(input)
		}
	}

	return -1, fmt.Errorf("unrecognized input %s", input)
}
