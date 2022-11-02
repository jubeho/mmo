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

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "to write tag - value pairs to given file(s)",
	Long:  `mmo set title "to some title" "to/this/file/or/folder"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Printf("subcommand 'set' needs exact 3 arguments: tagname, tagvalue and file(s); got %d args: %v\n", len(args), args)
			return
		}
		/*
			tagName := args[0]
			tagValue := args[1]
			fileOrFolder := args[2]

			var err error

			mymo, err = mmo.NewMMO([]string{fileOrFolder})
			if err != nil {
				logrus.Fatal(err)
			}

			for _, af := range mymo.AudioFiles {
				af.Set(tagName, tagValue)
				err := mymo.WriteFile(af)
				if err != nil {
					logrus.Error(err)
				}
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
