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

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/dice/matchers"
)

var Verbose bool

var Matchers = matchers.All()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Turn on verbose logging")
}

func initConfig() {
	if Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(RootCmd.OutOrStdout())
	logrus.StandardLogger().Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	logrus.StandardLogger().Formatter.(*logrus.TextFormatter).DisableLevelTruncation = true

}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "roll NdM",
	Short: "Roll some dice!",
	Args:  cobra.MinimumNArgs(1),
	Long: `Roll some dice! For example:

	$ roll 2d4

would roll (2) 4-sided dice. Other valid configurations:

` + makeExamples(),
	RunE: func(cmd *cobra.Command, args []string) error {
		// combine all args into one string
		allArgs := strings.Join(args, " ")
		_, result, err := roll(allArgs)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	},
}

func makeExamples() string {
	// examples := matchers.AllExamples()
	inputs := make([]string, len(Matchers))
	descriptions := make([]string, len(Matchers))
	maxInput := 0
	for i, m := range Matchers {
		input, desc := m.Example()
		inputs[i] = input
		descriptions[i] = desc
		if len(input) > maxInput {
			maxInput = len(input)
		}
	}

	gutter := "  "
	ret := "\tInput"
	for i := 0; i < maxInput-5; i++ {
		ret += " "
	}
	ret += gutter + "Description\n"
	for i := 0; i < len(Matchers); i++ {
		ret += "\t" + inputs[i]
		for j := 0; j < maxInput-len(inputs[i]); j++ {
			ret += " "
		}
		ret += gutter + descriptions[i] + "\n"
	}
	return ret
}

func roll(input string) (int, string, error) {
	for _, m := range Matchers {
		if m.Matches(input) {
			return m.Run(input)
		}
	}

	return -1, "", fmt.Errorf("unrecognized input %s", input)
}
