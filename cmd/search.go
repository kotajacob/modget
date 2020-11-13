/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/util"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <Search String>",
	Short: "Query and print more information about a specific mod by MODID.",
	Run: func(cmd *cobra.Command, args []string) {
		// Exit if no search terms given
		if len(args) == 0 {
			fmt.Println("modget search requires at least one search term")
			os.Exit(1)
		}
		for _, arg := range args {
			err := search(arg)
			if err != nil {
				fmt.Printf("Search failed: %v\n", arg)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(s string) error {
	var search curse.Search
	search.GameId = 432     // Set game to minecraft
	search.SectionId = 6    // Set section to mods
	search.SearchFilter = s // Search string
	addons, err := curse.AddonSearch(search)
	if err != nil {
		return err
	}
	for _, addon := range addons {
		util.DebugAddonPrint(addon)
	}
	return nil
}
