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
	"beckx.online/mmo/fui"
	"beckx.online/mmo/mmo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fuiCmd represents the fui command
var fuiCmd = &cobra.Command{
	Use:   "fui",
	Short: "fui starts the grafical user interface for mmo",
	Long:  `mmo fui path/to/files/`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = append(args, ".")
		}
		var err error
		mymo, err = mmo.NewMMO(args)
		if err != nil {
			logrus.Fatal(err)
		}

		err = fui.InitFui(mymo)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fuiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fuiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fuiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
