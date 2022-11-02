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

// genreCmd represents the genre command
var genreCmd = &cobra.Command{
	Use:   "genre",
	Short: "lists a summary of the genres of th given audio files",
	Long: `if no argument is given: lists genres of audio lib file
	mmo genre
	mmo genre path/to/audio/files
`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		mymo, err = mmo.NewMMO(args)
		if err != nil {
			logrus.Fatal(err)
		}

		flagList, err := cmd.Flags().GetString("list")
		if err != nil {
			logrus.Panicf("flag error list: %v", err)
		}

		// using "_NO_LIST_FLAG_" as default so i can search for empty genres ""
		if flagList != "_NO_LIST_FLAG_" {
			v, ok := mymo.Genres[flagList]
			if ok {
				fmt.Printf("found %d files with genre '%s'\n", len(v), flagList)
				for _, k := range v {
					fmt.Printf("\t%s\n", k)
					fmt.Printf("frame type: %s\n", k.TagType)
				}
			}
			return
		}

		for k, v := range mymo.Genres {
			fmt.Println(k, len(v))
		}
	},
}

func init() {
	rootCmd.AddCommand(genreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	genreCmd.Flags().StringP("list", "l", "_NO_LIST_FLAG_", "list files for given genre")
}
