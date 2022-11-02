/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	"beckx.online/mmo/mmo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "lists all frames/tags from given files/folders",
	Long:  `mmo info foo/bar.flac`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			return
		}
		var err error
		showSumm, err := cmd.Flags().GetBool("summary")
		if err != nil {
			logrus.Fatal("flagerror", err)
		}

		mymo, err = mmo.NewMMO(args)
		if err != nil {
			logrus.Fatal(err)
		}

		if showSumm {
			fmt.Println(mymo.GetAudiofileSummary())
			return
		}
		for _, af := range mymo.AudioFiles {
			fmt.Println(af)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	infoCmd.Flags().BoolP("summary", "s", false, "prints summary informations about collected audio files")
}
