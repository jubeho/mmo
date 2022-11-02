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

// libCmd represents the lib command
var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "commands to work with mmo library",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		isCreate, err := cmd.Flags().GetBool("create")
		if err != nil {
			logrus.Fatal("flagerror", err)
		}

		mymo, err = mmo.NewMMO(args)
		if err != nil {
			logrus.Fatal(err)
		}

		if isCreate {
			err = mymo.AudiofilesToLib("/home/juergen/mmoaudiolib.csv")
			if err != nil {
				logrus.Fatal(err)
			}
			return
		}

		for _, af := range mymo.AudioFiles {
			fmt.Println(af)
		}

	},
}

func init() {
	rootCmd.AddCommand(libCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// libCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	libCmd.Flags().BoolP("create", "c", false, "to create libfile (csv)")
}
